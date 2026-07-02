package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// RefreshToken calls GET /auth/token/refresh on the auth gateway using the
// configured refresh token, and returns the new token pair. It signs without an
// access_token (matching LazopClient). The token fields sit at the response
// root, not under "data".
func (c *Client) RefreshToken(ctx context.Context) (*TokenResult, error) {
	body, _, err := c.do(ctx, http.MethodGet, c.cfg.AuthURL, "", "/auth/token/refresh",
		map[string]string{"refresh_token": c.cfg.RefreshToken})
	if err != nil {
		return nil, fmt.Errorf("refresh token: %w", err)
	}
	var dto tokenResponseDTO
	if err := json.Unmarshal(body, &dto); err != nil {
		return nil, fmt.Errorf("decode refresh token: %w", err)
	}
	return &TokenResult{
		AccessToken:      dto.AccessToken,
		ExpiresIn:        int64(dto.ExpiresIn),
		RefreshToken:     dto.RefreshToken,
		RefreshExpiresIn: int64(dto.RefreshExpiresIn),
		ExpiresAt:        time.Now().UTC().Add(time.Duration(dto.ExpiresIn) * time.Second),
	}, nil
}
