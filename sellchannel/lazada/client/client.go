package client

import (
	"time"

	"github.com/go-resty/resty/v2"
)

// DefaultAuthURL is the Lazada OAuth gateway used for token refresh. It is the
// single source of truth for this default — config falls back to it too.
const DefaultAuthURL = "https://auth.lazada.com/rest"

// Config holds the Lazada Open Platform credentials for a single shop.
// BaseURL is the regional API gateway, e.g. https://api.lazada.co.th/rest.
// AuthURL is the OAuth gateway used for token refresh (defaults if empty).
type Config struct {
	AppKey       string
	AppSecret    string
	AccessToken  string
	RefreshToken string
	BaseURL      string
	AuthURL      string
}

const httpTimeout = 15 * time.Second

type Client struct {
	cfg  Config
	http *resty.Client
}

// New builds a Client with its own HTTP client. For many calls, prefer sharing
// one resty client via NewHTTPClient + NewWithHTTP to reuse the connection pool.
func New(cfg Config) *Client {
	return NewWithHTTP(cfg, NewHTTPClient())
}

// NewHTTPClient returns a resty client configured for the Lazada API. Share one
// instance across Client values so the underlying connection pool is reused.
func NewHTTPClient() *resty.Client {
	return resty.New().SetTimeout(httpTimeout)
}

// NewWithHTTP builds a Client that reuses the given resty client (shared
// transport / connection pool) instead of creating its own.
func NewWithHTTP(cfg Config, httpClient *resty.Client) *Client {
	if cfg.AuthURL == "" {
		cfg.AuthURL = DefaultAuthURL
	}
	return &Client{cfg: cfg, http: httpClient}
}
