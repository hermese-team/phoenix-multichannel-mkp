package client

import (
	"time"

	"github.com/okdev/marketplace-sync/pkg/httpclient"
)

type Config struct {
	PartnerID int64
	AppKey    string
	AppSecret string
	BaseURL   string
}

type Client struct {
	cfg  Config
	http *httpclient.Client
}

func New(cfg Config) *Client {
	return &Client{
		cfg:  cfg,
		http: httpclient.New(cfg.BaseURL, 30*time.Second),
	}
}
