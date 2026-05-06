package wdt

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type qimenTransport struct {
	cfg        QimenConfig
	httpClient *http.Client
	now        func() time.Time
}

type qimenEnvelope struct {
	Status  int64           `json:"status"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

func newQimenTransport(cfg QimenConfig, httpClient *http.Client, now func() time.Time) (*qimenTransport, error) {
	if cfg.GatewayURL == "" {
		return nil, fmt.Errorf("wdt: qimen gateway url is required")
	}
	if cfg.AppKey == "" || cfg.AppSecret == "" || cfg.TargetAppKey == "" {
		return nil, fmt.Errorf("wdt: qimen top app key, app secret and target app key are required")
	}
	if cfg.CustomerID == "" || cfg.WDTAppKey == "" || cfg.WDTAppSecret == "" || cfg.WDTSalt == "" {
		return nil, fmt.Errorf("wdt: qimen wdt credentials are required")
	}
	if cfg.Format == "" {
		cfg.Format = "json"
	}
	if cfg.SignMethod == "" {
		cfg.SignMethod = "md5"
	}
	if cfg.Version == "" {
		cfg.Version = "2.0"
	}
	if cfg.PartnerID == "" {
		cfg.PartnerID = "go-wdt"
	}
	return &qimenTransport{cfg: cfg, httpClient: httpClient, now: now}, nil
}

func (t *qimenTransport) call(ctx context.Context, method string, pager *Pager, params any, out any) error {
	normalizedMethod := NormalizeQimenMethod(method)
	now := t.now().Format(DateTimeLayout)

	form := map[string]string{
		"datetime":         now,
		"method":           normalizedMethod,
		"wdt3_customer_id": t.cfg.CustomerID,
		"wdt_appkey":       t.cfg.WDTAppKey,
		"wdt_salt":         t.cfg.WDTSalt,
	}
	signValues := map[string]any{
		"datetime":         now,
		"method":           normalizedMethod,
		"wdt3_customer_id": t.cfg.CustomerID,
		"wdt_appkey":       t.cfg.WDTAppKey,
		"wdt_salt":         t.cfg.WDTSalt,
	}
	if pagerJSON, pagerValue, err := jsonStringAndValue(pager); err != nil {
		return fmt.Errorf("wdt: marshal qimen pager: %w", err)
	} else if pagerValue != nil {
		form["pager"] = pagerJSON
		signValues["pager"] = pagerValue
	}
	if paramsJSON, paramsValue, err := jsonStringAndValue(params); err != nil {
		return fmt.Errorf("wdt: marshal qimen params: %w", err)
	} else if paramsValue != nil {
		form["params"] = paramsJSON
		signValues["params"] = paramsValue
	}
	form["wdt_sign"] = SignQimenCustom(t.cfg.WDTAppSecret, normalizedMethod, signValues)

	public := map[string]string{
		"app_key":        t.cfg.AppKey,
		"format":         t.cfg.Format,
		"method":         normalizedMethod,
		"partner_id":     t.cfg.PartnerID,
		"sign_method":    t.cfg.SignMethod,
		"target_app_key": t.cfg.TargetAppKey,
		"timestamp":      now,
		"v":              t.cfg.Version,
	}
	merged := make(map[string]string, len(public)+len(form))
	for key, value := range public {
		merged[key] = value
	}
	for key, value := range form {
		merged[key] = value
	}
	public["sign"] = signTopMD5(t.cfg.AppSecret, merged)

	endpoint, err := url.Parse(t.cfg.GatewayURL)
	if err != nil {
		return fmt.Errorf("wdt: parse qimen gateway url: %w", err)
	}
	query := endpoint.Query()
	for key, value := range public {
		query.Set(key, value)
	}
	endpoint.RawQuery = query.Encode()

	bodyValues := url.Values{}
	for key, value := range form {
		bodyValues.Set(key, value)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), strings.NewReader(bodyValues.Encode()))
	if err != nil {
		return fmt.Errorf("wdt: build qimen request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("wdt: qimen request failed: %w", err)
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("wdt: read qimen response: %w", err)
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return &APIError{Route: RouteQimen, Status: int64(resp.StatusCode), Message: http.StatusText(resp.StatusCode), Body: raw}
	}
	return decodeQimenEnvelope(raw, out)
}

func decodeQimenEnvelope(raw []byte, out any) error {
	unwrapped, err := unwrapTopResponse(raw)
	if err != nil {
		return err
	}
	var envelope qimenEnvelope
	if err := json.Unmarshal(unwrapped, &envelope); err != nil {
		return fmt.Errorf("wdt: decode qimen envelope: %w", err)
	}
	if envelope.Status > 0 {
		return &APIError{Route: RouteQimen, Status: envelope.Status, Message: envelope.Message, Body: raw}
	}
	if out == nil || len(envelope.Data) == 0 || bytes.Equal(envelope.Data, []byte("null")) {
		return nil
	}
	if err := json.Unmarshal(envelope.Data, out); err != nil {
		return fmt.Errorf("wdt: decode qimen data: %w", err)
	}
	return nil
}

func unwrapTopResponse(raw []byte) ([]byte, error) {
	var object map[string]json.RawMessage
	if err := json.Unmarshal(raw, &object); err != nil {
		return nil, fmt.Errorf("wdt: decode top response: %w", err)
	}
	if len(object) != 1 {
		return raw, nil
	}
	for key, value := range object {
		if key == "error_response" {
			var topErr struct {
				Code string `json:"code"`
				Msg  string `json:"msg"`
			}
			_ = json.Unmarshal(value, &topErr)
			return nil, &APIError{Route: RouteQimen, Code: topErr.Code, Message: topErr.Msg, Body: raw}
		}
		return value, nil
	}
	return raw, nil
}

func jsonStringAndValue(value any) (string, any, error) {
	if value == nil {
		return "", nil, nil
	}
	raw, err := json.Marshal(value)
	if err != nil {
		return "", nil, err
	}
	var normalized any
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.UseNumber()
	if err := decoder.Decode(&normalized); err != nil {
		return "", nil, err
	}
	if isEmptyJSONValue(normalized) {
		return "", nil, nil
	}
	return string(raw), normalized, nil
}

func isEmptyJSONValue(value any) bool {
	switch typed := value.(type) {
	case nil:
		return true
	case map[string]any:
		return len(typed) == 0
	case []any:
		return len(typed) == 0
	default:
		return false
	}
}
