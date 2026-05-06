package wdt

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func SignDirect(secret string, values map[string]string) string {
	keys := make([]string, 0, len(values))
	for key := range values {
		if key != "sign" {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)

	builder := strings.Builder{}
	builder.WriteString(secret)
	for _, key := range keys {
		builder.WriteString(key)
		builder.WriteString(values[key])
	}
	builder.WriteString(secret)
	return md5Lower(builder.String())
}

func NormalizeQimenMethod(method string) string {
	method = strings.TrimSpace(method)
	if strings.HasPrefix(method, "wdt.") {
		return method
	}
	if idx := strings.Index(method, "wdt."); idx >= 0 {
		return method[idx:]
	}
	if idx := strings.Index(method, "."); idx >= 0 {
		return method[idx+1:]
	}
	return method
}

func SignQimenCustom(secret, method string, values map[string]any) string {
	filtered := make(map[string]any, len(values)+1)
	for key, value := range values {
		if key == "wdt3_customer_id" || key == "target_app_key" || key == "wdt_sign" {
			continue
		}
		filtered[key] = value
	}
	filtered["method"] = NormalizeQimenMethod(method)

	builder := strings.Builder{}
	builder.WriteString(secret)
	appendSignedValues(&builder, filtered)
	builder.WriteString(secret)
	return md5Lower(builder.String())
}

type CRMRequest struct {
	PageNo        int            `json:"pageNo,omitempty"`
	PageSize      int            `json:"pageSize,omitempty"`
	Fields        string         `json:"fields,omitempty"`
	ExtendProps   map[string]any `json:"extendProps,omitempty"`
	CustomerID    string         `json:"customerid,omitempty"`
	Method        string         `json:"method,omitempty"`
	SDCode        string         `json:"sd_code,omitempty"`
	StartModified string         `json:"startModified,omitempty"`
	EndModified   string         `json:"endModified,omitempty"`
}

func SignQimenCRM(secret string, req CRMRequest) string {
	values := map[string]any{
		"pageNo":        zeroAsNil(req.PageNo),
		"pageSize":      zeroAsNil(req.PageSize),
		"fields":        emptyAsNil(req.Fields),
		"customerid":    emptyAsNil(req.CustomerID),
		"method":        emptyAsNil(req.Method),
		"sd_code":       emptyAsNil(req.SDCode),
		"startModified": emptyAsNil(req.StartModified),
		"endModified":   emptyAsNil(req.EndModified),
	}
	if len(req.ExtendProps) > 0 {
		props := make(map[string]any, len(req.ExtendProps))
		for key, value := range req.ExtendProps {
			if key != "wdt_sign" {
				props[key] = value
			}
		}
		values["extendProps"] = props
	}

	builder := strings.Builder{}
	builder.WriteString(secret)
	appendSignedValues(&builder, values)
	builder.WriteString(secret)
	return md5Lower(builder.String())
}

func (r CRMRequest) SignedExtendProps(secret string) (string, error) {
	props := make(map[string]any, len(r.ExtendProps)+1)
	for key, value := range r.ExtendProps {
		props[key] = value
	}
	props["wdt_sign"] = SignQimenCRM(secret, r)
	raw, err := json.Marshal(props)
	if err != nil {
		return "", fmt.Errorf("wdt: marshal extendProps: %w", err)
	}
	return string(raw), nil
}

func signTopMD5(secret string, values map[string]string) string {
	return strings.ToUpper(SignDirect(secret, values))
}

func appendSignedValues(builder *strings.Builder, values map[string]any) {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		value := values[key]
		if value == nil {
			continue
		}
		builder.WriteString(key)
		appendSignedValue(builder, value)
	}
}

func appendSignedValue(builder *strings.Builder, value any) {
	switch typed := value.(type) {
	case nil:
		return
	case map[string]any:
		appendSignedValues(builder, typed)
	case []any:
		for _, item := range typed {
			appendSignedValue(builder, item)
		}
	case string:
		if looksLikeJSON(typed) {
			var decoded any
			if err := json.Unmarshal([]byte(typed), &decoded); err == nil {
				appendSignedValue(builder, decoded)
				return
			}
		}
		builder.WriteString(typed)
	case bool:
		if typed {
			builder.WriteString("true")
		} else {
			builder.WriteString("false")
		}
	case json.Number:
		builder.WriteString(typed.String())
	default:
		builder.WriteString(fmt.Sprint(typed))
	}
}

func looksLikeJSON(text string) bool {
	text = strings.TrimSpace(text)
	return strings.HasPrefix(text, "{") || strings.HasPrefix(text, "[")
}

func zeroAsNil(value int) any {
	if value == 0 {
		return nil
	}
	return value
}

func emptyAsNil(value string) any {
	if value == "" {
		return nil
	}
	return value
}

func md5Lower(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
