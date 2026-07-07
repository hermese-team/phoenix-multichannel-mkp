package shopee

import "time"

type Config struct {
	PartnerID     int64
	AppKey        string
	AppSecret     string
	BaseURL       string
	WebhookVerify bool          // set SHOPEE_WEBHOOK_VERIFY=true in production
	PollSpec      string        // cron spec for fallback polling, e.g. "@every 5m"
	PollWindow    time.Duration // how far back to look per poll, e.g. 10m
}
