package httpclient

import (
	"time"

	"github.com/go-resty/resty/v2"
)

// New builds a shared resty client with sane defaults. One instance is safe
// for concurrent use and reuses connections across all outbound gateways.
// Per-call concerns (base URL, auth token) are set on the request, not here.
func New(timeout time.Duration) *resty.Client {
	return resty.New().
		SetTimeout(timeout).
		SetRetryCount(2).
		SetRetryWaitTime(200 * time.Millisecond).
		SetRetryMaxWaitTime(2 * time.Second).
		SetHeader("Accept", "application/json")
}
