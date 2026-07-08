package shopee

type Config struct {
	PartnerID     int64
	AppKey        string
	AppSecret     string
	BaseURL       string
	WebhookVerify    bool   // set SHOPEE_WEBHOOK_VERIFY=true in production
	PollSpec         string // cron spec for fallback polling, e.g. "@every 5m"
	FulfillmentSpec  string // cron spec for fulfillment scheduler, e.g. "@every 5m"
	FulfillmentLimit int    // max fulfillments per run
}
