package wdt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dromara/carbon/v2"
)

const DateTimeLayout = "2006-01-02 15:04:05"

type Time struct {
	carbon.Carbon
}

func (t *Time) UnmarshalJSON(data []byte) error {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 || bytes.Equal(trimmed, []byte("null")) || bytes.Equal(trimmed, []byte(`""`)) {
		t.Carbon = carbon.Carbon{}
		return nil
	}

	text := string(trimmed)
	if trimmed[0] == '"' {
		if err := json.Unmarshal(trimmed, &text); err != nil {
			return fmt.Errorf("wdt: decode time string: %w", err)
		}
	}

	parsed, err := ParseTime(strings.TrimSpace(text))
	if err != nil {
		return err
	}
	t.Carbon = *carbon.CreateFromStdTime(parsed)
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(t.ToDateTimeString())
}

func ParseTime(text string) (time.Time, error) {
	if text == "" {
		return time.Time{}, nil
	}
	if unixValue, err := strconv.ParseInt(text, 10, 64); err == nil {
		switch {
		case len(text) >= 13:
			return time.UnixMilli(unixValue), nil
		case len(text) >= 10:
			return time.Unix(unixValue, 0), nil
		}
	}
	layouts := []string{DateTimeLayout, time.RFC3339, time.RFC3339Nano, "2006-01-02"}
	for _, layout := range layouts {
		if parsed, err := time.ParseInLocation(layout, text, time.Local); err == nil {
			return parsed, nil
		}
	}
	return time.Time{}, fmt.Errorf("wdt: unsupported time format %q", text)
}
