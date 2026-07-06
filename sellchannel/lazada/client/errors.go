package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// APIErrorDetail is one entry of a Lazada error `detail` array, e.g.
// {"field":"Product","message":"CATEGORY_ID_INVALID"}.
type APIErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// APIError is a Lazada response with a non-"0" envelope code.
type APIError struct {
	APIPath   string
	Code      string
	Type      string
	Message   string
	RequestID string
	TraceID   string
	Details   []APIErrorDetail
}

// newAPIError builds an APIError from a failed envelope. The `detail` field is
// best-effort parsed as [{field,message}]; when Lazada instead returns a string
// or object there, Details stays empty and the rest of the error is still usable.
func newAPIError(apiPath string, env envelope) *APIError {
	e := &APIError{
		APIPath:   apiPath,
		Code:      env.Code,
		Type:      env.Type,
		Message:   env.Message,
		RequestID: env.RequestID,
		TraceID:   env.TraceID,
	}
	_ = json.Unmarshal(env.Detail, &e.Details)
	return e
}

func (e *APIError) Error() string {
	var b strings.Builder
	fmt.Fprintf(&b, "lazada %s failed: code=%s message=%s", e.APIPath, e.Code, e.Message)
	for _, d := range e.Details {
		fmt.Fprintf(&b, " [%s: %s]", d.Field, d.Message)
	}
	if e.TraceID != "" {
		fmt.Fprintf(&b, " trace_id=%s", e.TraceID)
	}
	return b.String()
}

// IsAuthError reports whether err is a Lazada API error caused by an invalid or
// expired access token (code "IllegalAccessToken") — the recoverable case where
// a token refresh + retry is worthwhile. Transport errors (network, timeout)
// are not APIErrors and return false, so callers won't refresh needlessly.
func IsAuthError(err error) bool {
	var apiErr *APIError
	return errors.As(err, &apiErr) && apiErr.Code == "IllegalAccessToken"
}
