package wdt

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"
)

const directEpochOffset = 1325347200

type directTransport struct {
	cfg        DirectConfig
	httpClient *http.Client
	now        func() time.Time
}

type directEnvelope struct {
	Status  int64           `json:"status"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

func newDirectTransport(cfg DirectConfig, httpClient *http.Client, now func() time.Time) (*directTransport, error) {
	if cfg.GatewayURL == "" {
		return nil, fmt.Errorf("wdt: direct gateway url is required")
	}
	if cfg.SID == "" || cfg.AppKey == "" || cfg.AppSecret == "" || cfg.Salt == "" {
		return nil, fmt.Errorf("wdt: direct sid, app key, app secret and salt are required")
	}
	if cfg.Version == "" {
		cfg.Version = "1.0"
	}
	cfg.GatewayURL = normalizeDirectGateway(cfg.GatewayURL)
	return &directTransport{cfg: cfg, httpClient: httpClient, now: now}, nil
}

func (t *directTransport) call(ctx context.Context, method string, style DirectBodyStyle, pager *Pager, params any, out any) error {
	body, bodyForSign, err := encodeDirectBody(style, params)
	if err != nil {
		return err
	}

	query := map[string]string{
		"sid":       t.cfg.SID,
		"key":       t.cfg.AppKey,
		"salt":      t.cfg.Salt,
		"method":    method,
		"timestamp": strconv.FormatInt(t.now().Unix()-directEpochOffset, 10),
		"v":         t.cfg.Version,
		"body":      bodyForSign,
	}
	if pager != nil {
		query["page_no"] = strconv.Itoa(pager.PageNo)
		query["page_size"] = strconv.Itoa(pager.PageSize)
		if pager.CalcTotal {
			query["calc_total"] = "1"
		} else {
			query["calc_total"] = "0"
		}
	}
	query["sign"] = SignDirect(t.cfg.AppSecret, query)
	delete(query, "body")

	endpoint, err := url.Parse(t.cfg.GatewayURL)
	if err != nil {
		return fmt.Errorf("wdt: parse direct gateway url: %w", err)
	}
	values := endpoint.Query()
	for key, value := range query {
		values.Set(key, value)
	}
	endpoint.RawQuery = values.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("wdt: build direct request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("wdt: direct request failed: %w", err)
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("wdt: read direct response: %w", err)
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return &APIError{Route: RouteDirect, Status: int64(resp.StatusCode), Message: http.StatusText(resp.StatusCode), Body: raw}
	}
	return decodeDirectEnvelope(raw, out)
}

func encodeDirectBody(style DirectBodyStyle, params any) ([]byte, string, error) {
	var payload any
	if style == DirectBodyArray {
		payload = params
	} else {
		payload = []any{params}
	}
	if payload == nil {
		payload = []any{map[string]any{}}
	}
	raw, err := json.Marshal(payload)
	if err != nil {
		return nil, "", fmt.Errorf("wdt: marshal direct body: %w", err)
	}
	return raw, string(raw), nil
}

func decodeDirectEnvelope(raw []byte, out any) error {
	var envelope directEnvelope
	if err := json.Unmarshal(raw, &envelope); err != nil {
		return fmt.Errorf("wdt: decode direct envelope: %w", err)
	}
	if envelope.Status > 0 {
		return &APIError{Route: RouteDirect, Status: envelope.Status, Message: envelope.Message, Body: raw}
	}
	if out == nil || len(envelope.Data) == 0 || bytes.Equal(envelope.Data, []byte("null")) {
		return nil
	}
	if err := json.Unmarshal(envelope.Data, out); err != nil {
		return fmt.Errorf("wdt: decode direct data: %w", err)
	}
	return nil
}

func normalizeDirectGateway(rawURL string) string {
	if strings.HasSuffix(rawURL, "/openapi") || strings.HasSuffix(rawURL, "openapi") {
		return rawURL
	}
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return rawURL
	}
	parsed.Path = path.Join(parsed.Path, "openapi")
	return parsed.String()
}
