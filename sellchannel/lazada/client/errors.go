package client

import (
	"errors"
	"fmt"
)

// APIError is a Lazada response with a non-"0" envelope code.
type APIError struct {
	APIPath string
	Code    string
	Message string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("lazada %s failed: code=%s message=%s", e.APIPath, e.Code, e.Message)
}

// IsAuthError reports whether err is a Lazada API error caused by an invalid or
// expired access token (code "IllegalAccessToken") — the recoverable case where
// a token refresh + retry is worthwhile. Transport errors (network, timeout)
// are not APIErrors and return false, so callers won't refresh needlessly.
func IsAuthError(err error) bool {
	var apiErr *APIError
	return errors.As(err, &apiErr) && apiErr.Code == "IllegalAccessToken"
}
