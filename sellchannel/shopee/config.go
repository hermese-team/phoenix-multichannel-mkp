package shopee

type Config struct {
	PartnerID     int64
	AppKey        string
	AppSecret     string
	BaseURL       string
	WebhookVerify bool // set SHOPEE_WEBHOOK_VERIFY=true in production
}
