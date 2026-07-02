package client

import "time"

// tokenResponseDTO captures the root-level fields of /auth/token/refresh
// (these are NOT nested under "data").
type tokenResponseDTO struct {
	AccessToken      string  `json:"access_token"`
	ExpiresIn        float64 `json:"expires_in"`
	RefreshToken     string  `json:"refresh_token"`
	RefreshExpiresIn float64 `json:"refresh_expires_in"`
}

// TokenResult is the public result of a token refresh.
type TokenResult struct {
	AccessToken      string
	ExpiresIn        int64 // seconds
	RefreshToken     string
	RefreshExpiresIn int64 // seconds
	ExpiresAt        time.Time
}
