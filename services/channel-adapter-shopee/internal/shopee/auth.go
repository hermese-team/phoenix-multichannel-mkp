package shopee

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	redisKeyAccessToken  = "shopee:token:access:%d"  // :shop_id
	redisKeyRefreshToken = "shopee:token:refresh:%d" // :shop_id
)

// TokenStore abstracts read/write of Shopee OAuth tokens.
// Redis is used so all pods share the same token and refresh is done once.
type TokenStore interface {
	// AccessToken returns a valid access token, refreshing if needed.
	AccessToken(ctx context.Context) (string, error)
}

// RedisTokenStore stores tokens in Redis and auto-refreshes using the refresh token.
type RedisTokenStore struct {
	rdb        redis.UniversalClient
	httpClient *http.Client
	baseURL    string
	signer     *Signer
	partnerID  int64
	shopID     int64
}

func NewRedisTokenStore(
	rdb redis.UniversalClient,
	baseURL string,
	signer *Signer,
	partnerID, shopID int64,
) *RedisTokenStore {
	return &RedisTokenStore{
		rdb:        rdb,
		httpClient: &http.Client{Timeout: 10 * time.Second},
		baseURL:    baseURL,
		signer:     signer,
		partnerID:  partnerID,
		shopID:     shopID,
	}
}

// Seed bootstraps initial tokens (called once at startup from config values).
// expiry is how long until the access token expires.
func (s *RedisTokenStore) Seed(ctx context.Context, accessToken, refreshToken string, expiry time.Duration) error {
	if err := s.rdb.Set(ctx, s.accessKey(), accessToken, expiry-30*time.Second).Err(); err != nil {
		return fmt.Errorf("seeding access token: %w", err)
	}
	// Refresh token validity is typically 30 days; store without strict TTL.
	if err := s.rdb.Set(ctx, s.refreshKey(), refreshToken, 30*24*time.Hour).Err(); err != nil {
		return fmt.Errorf("seeding refresh token: %w", err)
	}
	return nil
}

// AccessToken returns a valid access token from Redis, refreshing if absent/expired.
func (s *RedisTokenStore) AccessToken(ctx context.Context) (string, error) {
	token, err := s.rdb.Get(ctx, s.accessKey()).Result()
	if err == nil && token != "" {
		return token, nil
	}
	if err != redis.Nil {
		return "", fmt.Errorf("reading access token from redis: %w", err)
	}
	// Token absent or expired — refresh.
	return s.refresh(ctx)
}

// refresh calls Shopee's token/refresh endpoint and stores the new tokens.
func (s *RedisTokenStore) refresh(ctx context.Context) (string, error) {
	refreshToken, err := s.rdb.Get(ctx, s.refreshKey()).Result()
	if err != nil {
		return "", fmt.Errorf("reading refresh token: %w (re-authorization required)", err)
	}

	const path = "/api/v2/auth/access_token/get"
	q := s.signer.SignPartner(path)
	rawURL := s.baseURL + path + "?" + q.Encode()

	body := RefreshTokenRequest{
		RefreshToken: refreshToken,
		ShopID:       s.shopID,
		PartnerID:    s.partnerID,
	}
	b, _ := json.Marshal(body)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, rawURL, bytes.NewReader(b))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("refreshing token: %w", err)
	}
	defer resp.Body.Close()

	var tr TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tr); err != nil {
		return "", fmt.Errorf("decoding refresh response: %w", err)
	}
	if tr.IsError() {
		return "", fmt.Errorf("shopee token refresh error: %s – %s", tr.Error, tr.Message)
	}

	expiry := time.Duration(tr.ExpireIn) * time.Second
	if expiry <= 0 {
		expiry = 4 * time.Hour // Shopee default
	}

	// Store new tokens — subtract 30s buffer so we refresh before actual expiry.
	pipe := s.rdb.Pipeline()
	pipe.Set(ctx, s.accessKey(), tr.AccessToken, expiry-30*time.Second)
	pipe.Set(ctx, s.refreshKey(), tr.RefreshToken, 30*24*time.Hour)
	if _, err := pipe.Exec(ctx); err != nil {
		return "", fmt.Errorf("storing refreshed tokens: %w", err)
	}

	return tr.AccessToken, nil
}

// ExchangeCode exchanges an OAuth authorization code for access+refresh tokens
// and seeds them into Redis. Call this from the OAuth callback handler.
func (s *RedisTokenStore) ExchangeCode(ctx context.Context, code string) error {
	const path = "/api/v2/auth/token/get"
	q := s.signer.SignPartner(path)
	rawURL := s.baseURL + path + "?" + q.Encode()

	body := GetAccessTokenRequest{
		Code:      code,
		ShopID:    s.shopID,
		PartnerID: s.partnerID,
	}
	b, _ := json.Marshal(body)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, rawURL, bytes.NewReader(b))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("exchanging code: %w", err)
	}
	defer resp.Body.Close()

	var tr TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tr); err != nil {
		return fmt.Errorf("decoding token response: %w", err)
	}
	if tr.IsError() {
		return fmt.Errorf("shopee token exchange error: %s – %s", tr.Error, tr.Message)
	}

	expiry := time.Duration(tr.ExpireIn) * time.Second
	if expiry <= 0 {
		expiry = 4 * time.Hour
	}
	return s.Seed(ctx, tr.AccessToken, tr.RefreshToken, expiry)
}

// GetAuthURL returns the Shopee OAuth authorization URL for a shop to grant access.
func GetAuthURL(baseURL string, partnerID int64, signer *Signer, redirectURL string) string {
	const path = "/api/v2/shop/auth_partner"
	q := signer.SignPartner(path)
	q.Set("redirect", redirectURL)
	return baseURL + path + "?" + q.Encode()
}

func (s *RedisTokenStore) accessKey() string {
	return fmt.Sprintf(redisKeyAccessToken, s.shopID)
}
func (s *RedisTokenStore) refreshKey() string {
	return fmt.Sprintf(redisKeyRefreshToken, s.shopID)
}
