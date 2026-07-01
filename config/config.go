package config

// Config is the root configuration for channel-adapter-shopee.
// All fields are populated from config.json (seed) overridden by environment
// variables injected by ExternalSecret from Vault at runtime.
//
// Env var naming: nested keys joined by "_" and uppercased.
// e.g. shopee.partner_key → SHOPEE_PARTNER_KEY
type Config struct {
	AppEnv  string `mapstructure:"app_env"     json:"app_env"`
	AppPort int    `mapstructure:"server_port" json:"server_port"`

	Logger struct {
		ServiceOwner string `mapstructure:"service_owner" json:"service_owner"`
		LogLevel     string `mapstructure:"log_level"     json:"log_level"`
	} `mapstructure:"logger" json:"logger"`

	Shopee ShopeeConfig `mapstructure:"shopee" json:"shopee"`

	Kafka KafkaConfig `mapstructure:"kafka" json:"kafka"`

	Redis RedisConfig `mapstructure:"redis" json:"redis"`

	Telemetry struct {
		Endpoint    string `mapstructure:"endpoint"     json:"endpoint"`
		ServiceName string `mapstructure:"service_name" json:"service_name"`
	} `mapstructure:"telemetry" json:"telemetry"`

	Poller PollerConfig `mapstructure:"poller" json:"poller"`
}

type ShopeeConfig struct {
	// Credentials — injected from Vault
	PartnerID  int64  `mapstructure:"partner_id"  json:"partner_id"`
	PartnerKey string `mapstructure:"partner_key" json:"partner_key"`
	ShopID     int64  `mapstructure:"shop_id"     json:"shop_id"`

	// OAuth tokens — injected from Vault or refreshed at runtime
	AccessToken  string `mapstructure:"access_token"  json:"access_token"`
	RefreshToken string `mapstructure:"refresh_token" json:"refresh_token"`

	// API endpoints
	BaseURL    string `mapstructure:"base_url"    json:"base_url"`
	TimeoutSec int    `mapstructure:"timeout_sec" json:"timeout_sec"`

	// Redirect URL registered in Shopee app (used for OAuth authorize flow)
	RedirectURL string `mapstructure:"redirect_url" json:"redirect_url"`
}

type KafkaConfig struct {
	Brokers  []string `mapstructure:"brokers"   json:"brokers"`
	GroupID  string   `mapstructure:"group_id"  json:"group_id"`
	Username string   `mapstructure:"user"      json:"user"`
	Password string   `mapstructure:"pass"      json:"pass"`

	// Topic names — follow architecture doc naming convention
	TopicRawOrderAccepted string `mapstructure:"topic_raw_order_accepted" json:"topic_raw_order_accepted"`
}

type RedisConfig struct {
	Mode     string   `mapstructure:"mode"     json:"mode"`
	Addrs    []string `mapstructure:"addrs"    json:"addrs"`
	User     string   `mapstructure:"user"     json:"user"`
	Password string   `mapstructure:"password" json:"password"`
	DB       int      `mapstructure:"db"       json:"db"`
}

type PollerConfig struct {
	// How many orders per page (max 100 per Shopee API spec)
	PageSize int `mapstructure:"page_size" json:"page_size"`

	// Interval between polling cycles in seconds
	IntervalSec int `mapstructure:"interval_sec" json:"interval_sec"`

	// How far back to start polling if no cursor is saved (seconds from now)
	LookbackSec int64 `mapstructure:"lookback_sec" json:"lookback_sec"`

	// Max outbound API calls per minute (Shopee quota — leave 20% reserve)
	RateLimitPerMin int `mapstructure:"rate_limit_per_min" json:"rate_limit_per_min"`
}
