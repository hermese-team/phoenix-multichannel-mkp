package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Load reads config.json as a seed, then overrides every field with
// the corresponding environment variable (uppercased, dots→underscores).
// e.g. shopee.partner_key is overridden by SHOPEE_PARTNER_KEY.
func Load() (Config, error) {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("json")
	v.AddConfigPath(".")
	v.AddConfigPath("/app")

	if err := v.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("reading config file: %w", err)
	}

	// Allow any env var to override config keys.
	// Viper matches by replacing "." with "_" and uppercasing.
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return Config{}, fmt.Errorf("unmarshalling config: %w", err)
	}

	// KAFKA_BROKERS may arrive as a comma-separated string from ExternalSecret.
	// Viper handles []string from env as comma-split automatically, but guard here.
	if raw := os.Getenv("KAFKA_BROKERS"); raw != "" && len(cfg.Kafka.Brokers) == 0 {
		cfg.Kafka.Brokers = strings.Split(raw, ",")
	}

	// REDIS_ADDRS same pattern.
	if raw := os.Getenv("REDIS_ADDRS"); raw != "" && len(cfg.Redis.Addrs) == 0 {
		cfg.Redis.Addrs = strings.Split(raw, ",")
	}

	if err := validate(cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func validate(cfg Config) error {
	if cfg.Shopee.PartnerID == 0 {
		return fmt.Errorf("shopee.partner_id is required")
	}
	if cfg.Shopee.PartnerKey == "" {
		return fmt.Errorf("shopee.partner_key is required")
	}
	if cfg.Shopee.ShopID == 0 {
		return fmt.Errorf("shopee.shop_id is required")
	}
	if cfg.Shopee.BaseURL == "" {
		return fmt.Errorf("shopee.base_url is required")
	}
	if len(cfg.Kafka.Brokers) == 0 {
		return fmt.Errorf("kafka.brokers is required")
	}
	if cfg.Kafka.TopicRawOrderAccepted == "" {
		return fmt.Errorf("kafka.topic_raw_order_accepted is required")
	}
	return nil
}

// MustLoad is a convenience wrapper that panics on error (use in main only).
func MustLoad() Config {
	cfg, err := Load()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
	return cfg
}

// RedactedJSON returns config as JSON with sensitive fields masked — safe for logging.
func RedactedJSON(cfg Config) string {
	type redacted struct {
		AppEnv  string `json:"app_env"`
		AppPort int    `json:"server_port"`
		Shopee  struct {
			PartnerID int64  `json:"partner_id"`
			ShopID    int64  `json:"shop_id"`
			BaseURL   string `json:"base_url"`
			// secrets omitted intentionally
		} `json:"shopee"`
	}
	r := redacted{AppEnv: cfg.AppEnv, AppPort: cfg.AppPort}
	r.Shopee.PartnerID = cfg.Shopee.PartnerID
	r.Shopee.ShopID = cfg.Shopee.ShopID
	r.Shopee.BaseURL = cfg.Shopee.BaseURL
	b, _ := json.Marshal(r)
	return string(b)
}
