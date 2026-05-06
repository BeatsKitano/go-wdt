package wdt

import "fmt"

type APIError struct {
	Route   Route
	Status  int64
	Code    string
	Message string
	Body    []byte
}

func (e *APIError) Error() string {
	if e == nil {
		return ""
	}
	if e.Code != "" {
		return fmt.Sprintf("wdt: %s api error code=%s status=%d message=%s", e.Route, e.Code, e.Status, e.Message)
	}
	return fmt.Sprintf("wdt: %s api error status=%d message=%s", e.Route, e.Status, e.Message)
}
