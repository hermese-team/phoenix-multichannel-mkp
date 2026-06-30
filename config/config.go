package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	App         AppConfig
	Postgres    PostgresConfig
	Redis       RedisConfig
	Kafka       KafkaConfig
	Marketplace MarketplaceConfig
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

func Load() (*Config, error) {
	pgPort, err := parseInt("POSTGRES_PORT", 5432)
	if err != nil {
		return nil, err
	}
	redisPort, err := parseInt("REDIS_PORT", 6379)
	if err != nil {
		return nil, err
	}
	redisDB, err := parseInt("REDIS_DB", 0)
	if err != nil {
		return nil, err
	}
	pgMaxOpen, err := parseInt("POSTGRES_MAX_OPEN_CONNS", 25)
	if err != nil {
		return nil, err
	}
	pgMaxIdle, err := parseInt("POSTGRES_MAX_IDLE_CONNS", 5)
	if err != nil {
		return nil, err
	}
	pgLifetime, err := parseDuration("POSTGRES_CONN_MAX_LIFETIME", 5*time.Minute)
	if err != nil {
		return nil, err
	}
	redisPool, err := parseInt("REDIS_POOL_SIZE", 10)
	if err != nil {
		return nil, err
	}
	tiktok, err := loadEndpoint("MARKETPLACE_TIKTOK", "https://mock.tiktok.local")
	if err != nil {
		return nil, err
	}
	lazada, err := loadEndpoint("MARKETPLACE_LAZADA", "https://mock.lazada.local")
	if err != nil {
		return nil, err
	}
	shopee, err := loadEndpoint("MARKETPLACE_SHOPEE", "https://mock.shopee.local")
	if err != nil {
		return nil, err
	}

	return &Config{
		App: AppConfig{
			Port:     getEnv("PORT", "8080"),
			Env:      getEnv("APP_ENV", "development"),
			LogLevel: getEnv("LOG_LEVEL", "info"),
		},
		Postgres: PostgresConfig{
			Host:            getEnv("POSTGRES_HOST", "localhost"),
			Port:            pgPort,
			User:            getEnv("POSTGRES_USER", "postgres"),
			Password:        getEnv("POSTGRES_PASSWORD", ""),
			Database:        getEnv("POSTGRES_DB", "phoenix_mkp"),
			SSLMode:         getEnv("POSTGRES_SSLMODE", "disable"),
			MaxOpenConns:    pgMaxOpen,
			MaxIdleConns:    pgMaxIdle,
			ConnMaxLifetime: pgLifetime,
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     redisPort,
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       redisDB,
			PoolSize: redisPool,
		},
		Kafka: KafkaConfig{
			Brokers:         splitCSV(getEnv("KAFKA_BROKERS", "localhost:9092")),
			GroupID:         getEnv("KAFKA_GROUP_ID", "phoenix-mkp"),
			AutoOffsetReset: getEnv("KAFKA_AUTO_OFFSET_RESET", "earliest"),
			Topics: KafkaTopics{
				OrderCreated:   getEnv("KAFKA_TOPIC_ORDER_CREATED", "order.created"),
				ProductUpdated: getEnv("KAFKA_TOPIC_PRODUCT_UPDATED", "product.updated"),
			},
		},
		Marketplace: MarketplaceConfig{
			TikTok: tiktok,
			Lazada: lazada,
			Shopee: shopee,
		},
	}, nil
}

// loadEndpoint reads BASE_URL / API_KEY / TIMEOUT for a marketplace prefix.
func loadEndpoint(prefix, defaultBaseURL string) (ChannelEndpoint, error) {
	timeout, err := parseDuration(prefix+"_TIMEOUT", 10*time.Second)
	if err != nil {
		return ChannelEndpoint{}, err
	}
	return ChannelEndpoint{
		BaseURL: getEnv(prefix+"_BASE_URL", defaultBaseURL),
		APIKey:  getEnv(prefix+"_API_KEY", ""),
		Timeout: timeout,
	}, nil
}

func getEnv(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}

func parseInt(key string, defaultVal int) (int, error) {
	v := os.Getenv(key)
	if v == "" {
		return defaultVal, nil
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("env %s: expected integer, got %q", key, v)
	}
	return n, nil
}

func parseDuration(key string, defaultVal time.Duration) (time.Duration, error) {
	v := os.Getenv(key)
	if v == "" {
		return defaultVal, nil
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		return 0, fmt.Errorf("env %s: expected duration (e.g. 5m), got %q", key, v)
	}
	return d, nil
}

func splitCSV(s string) []string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if t := strings.TrimSpace(p); t != "" {
			out = append(out, t)
		}
	}
	return out
}
