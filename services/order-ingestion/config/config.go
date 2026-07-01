package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App         AppConfig
	Postgres    PostgresConfig
	Redis       RedisConfig
	Kafka       KafkaConfig
	Marketplace MarketplaceConfig
	Telemetry   TelemetryConfig
}

type AppConfig struct {
	Port     string
	Env      string
	LogLevel string
}

type PostgresConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	Database        string
	SSLMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func (c PostgresConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.Database, c.SSLMode,
	)
}

// MigrateDSN returns a pgx5-scheme DSN for golang-migrate.
func (c PostgresConfig) MigrateDSN() string {
	return fmt.Sprintf("pgx5://%s:%s@%s:%d/%s?sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.Database, c.SSLMode,
	)
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
	PoolSize int
}

func (c RedisConfig) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

type KafkaConfig struct {
	Brokers         []string
	GroupID         string
	AutoOffsetReset string
	Topics          KafkaTopics
}

type KafkaTopics struct {
	OrderCreated   string
	ProductUpdated string
}

// MarketplaceConfig holds the outbound endpoints for each sales channel.
type MarketplaceConfig struct {
	TikTok ChannelEndpoint
	Lazada ChannelEndpoint
	Shopee ChannelEndpoint
}

// ChannelEndpoint is the connection info for a single external marketplace API.
type ChannelEndpoint struct {
	BaseURL string
	APIKey  string
	Timeout time.Duration
}

// TelemetryConfig controls OpenTelemetry export.
// Leave Endpoint empty to disable (uses no-op provider).
type TelemetryConfig struct {
	Endpoint    string
	ServiceName string
}

// Load reads configuration from environment variables (and optionally config.json).
// Env-var mapping: dots → underscores, e.g. "postgres.host" → POSTGRES_HOST.
func Load() (*Config, error) {
	v := viper.New()

	// Optional config file — env vars always win.
	v.SetConfigName("config")
	v.SetConfigType("json")
	v.AddConfigPath(".")
	_ = v.ReadInConfig()

	// Map POSTGRES_HOST → postgres.host, etc.
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// ── Defaults ─────────────────────────────────────────────────────────────
	v.SetDefault("app.port", "8080")
	v.SetDefault("app.env", "development")
	v.SetDefault("app.log_level", "info")

	v.SetDefault("postgres.host", "localhost")
	v.SetDefault("postgres.port", 5432)
	v.SetDefault("postgres.user", "postgres")
	v.SetDefault("postgres.database", "phoenix_mkp")
	v.SetDefault("postgres.sslmode", "disable")
	v.SetDefault("postgres.max_open_conns", 25)
	v.SetDefault("postgres.max_idle_conns", 5)
	v.SetDefault("postgres.conn_max_lifetime", "5m")

	v.SetDefault("redis.host", "localhost")
	v.SetDefault("redis.port", 6379)
	v.SetDefault("redis.db", 0)
	v.SetDefault("redis.pool_size", 10)

	v.SetDefault("kafka.brokers", []string{"localhost:9092"})
	v.SetDefault("kafka.group_id", "phoenix-mkp")
	v.SetDefault("kafka.auto_offset_reset", "earliest")
	v.SetDefault("kafka.topics.order_created", "order.created")
	v.SetDefault("kafka.topics.product_updated", "product.updated")

	v.SetDefault("marketplace.tiktok.base_url", "https://mock.tiktok.local")
	v.SetDefault("marketplace.tiktok.timeout", "10s")
	v.SetDefault("marketplace.lazada.base_url", "https://mock.lazada.local")
	v.SetDefault("marketplace.lazada.timeout", "10s")
	v.SetDefault("marketplace.shopee.base_url", "https://mock.shopee.local")
	v.SetDefault("marketplace.shopee.timeout", "10s")

	v.SetDefault("telemetry.service_name", "phoenix-multichannel-mkp")

	return &Config{
		App: AppConfig{
			Port:     v.GetString("app.port"),
			Env:      v.GetString("app.env"),
			LogLevel: v.GetString("app.log_level"),
		},
		Postgres: PostgresConfig{
			Host:            v.GetString("postgres.host"),
			Port:            v.GetInt("postgres.port"),
			User:            v.GetString("postgres.user"),
			Password:        v.GetString("postgres.password"),
			Database:        v.GetString("postgres.database"),
			SSLMode:         v.GetString("postgres.sslmode"),
			MaxOpenConns:    v.GetInt("postgres.max_open_conns"),
			MaxIdleConns:    v.GetInt("postgres.max_idle_conns"),
			ConnMaxLifetime: v.GetDuration("postgres.conn_max_lifetime"),
		},
		Redis: RedisConfig{
			Host:     v.GetString("redis.host"),
			Port:     v.GetInt("redis.port"),
			Password: v.GetString("redis.password"),
			DB:       v.GetInt("redis.db"),
			PoolSize: v.GetInt("redis.pool_size"),
		},
		Kafka: KafkaConfig{
			Brokers:         v.GetStringSlice("kafka.brokers"),
			GroupID:         v.GetString("kafka.group_id"),
			AutoOffsetReset: v.GetString("kafka.auto_offset_reset"),
			Topics: KafkaTopics{
				OrderCreated:   v.GetString("kafka.topics.order_created"),
				ProductUpdated: v.GetString("kafka.topics.product_updated"),
			},
		},
		Marketplace: MarketplaceConfig{
			TikTok: endpoint(v, "marketplace.tiktok"),
			Lazada: endpoint(v, "marketplace.lazada"),
			Shopee: endpoint(v, "marketplace.shopee"),
		},
		Telemetry: TelemetryConfig{
			Endpoint:    v.GetString("telemetry.endpoint"),
			ServiceName: v.GetString("telemetry.service_name"),
		},
	}, nil
}

func endpoint(v *viper.Viper, prefix string) ChannelEndpoint {
	timeout := v.GetDuration(prefix + ".timeout")
	if timeout == 0 {
		timeout = 10 * time.Second
	}
	return ChannelEndpoint{
		BaseURL: v.GetString(prefix + ".base_url"),
		APIKey:  v.GetString(prefix + ".api_key"),
		Timeout: timeout,
	}
}
