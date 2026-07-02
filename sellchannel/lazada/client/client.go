package client

import (
	"time"

	"github.com/go-resty/resty/v2"
)

const defaultAuthURL = "https://auth.lazada.com/rest"

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

type Client struct {
	cfg  Config
	http *resty.Client
}

func New(cfg Config) *Client {
	if cfg.AuthURL == "" {
		cfg.AuthURL = defaultAuthURL
	}
	return &Client{
		cfg:  cfg,
		http: resty.New().SetTimeout(15 * time.Second),
	}
}
