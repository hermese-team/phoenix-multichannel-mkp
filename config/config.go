package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/okdev/marketplace-sync/internal/infrastructure/kafka"
	"github.com/okdev/marketplace-sync/internal/infrastructure/postgres"
	"github.com/okdev/marketplace-sync/internal/infrastructure/redis"
	lazadaclient "github.com/okdev/marketplace-sync/sellchannel/lazada/client"
	"github.com/okdev/marketplace-sync/sellchannel/shopee"
)

type App struct {
	Env      string
	Port     string
	LogLevel string
}

type Config struct {
	App      App
	Shopee   shopee.Config
	Lazada   lazadaclient.Config
	Postgres postgres.Config
	Redis    redis.Config
	Kafka    kafka.Config
}

func Load() (*Config, error) {
	pgLifetime, err := getEnvDuration("POSTGRES_CONN_MAX_LIFETIME", 5*time.Minute)
	if err != nil {
		return nil, err
	}

	return &Config{
		App: App{
			Env:      getEnv("APP_ENV", "development"),
			Port:     getEnv("APP_PORT", "8080"),
			LogLevel: getEnv("LOG_LEVEL", "info"),
		},
		Shopee: shopee.Config{
			PartnerID: int64(getEnvInt("SHOPEE_PARTNER_ID", 0)),
			AppKey:    os.Getenv("SHOPEE_APP_KEY"),
			AppSecret: os.Getenv("SHOPEE_APP_SECRET"),
			BaseURL:   getEnv("SHOPEE_BASE_URL", "https://partner.shopeemobile.com"),
		},
		Lazada: lazadaclient.Config{
			AppKey:    os.Getenv("LAZADA_APP_KEY"),
			AppSecret: os.Getenv("LAZADA_APP_SECRET"),
			BaseURL:   getEnv("LAZADA_URL_API", "https://api.lazada.co.th/rest"),
			AuthURL:   getEnv("LAZADA_AUTH_URL", "https://auth.lazada.com/rest"),
		},
		Postgres: postgres.Config{
			DSN:             getEnv("POSTGRES_DSN", "postgres://postgres:postgres@localhost:5432/marketplace_sync?sslmode=disable"),
			MaxOpenConns:    getEnvInt("POSTGRES_MAX_OPEN_CONNS", 25),
			MaxIdleConns:    getEnvInt("POSTGRES_MAX_IDLE_CONNS", 5),
			ConnMaxLifetime: pgLifetime,
		},
		Redis: redis.Config{
			Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       getEnvInt("REDIS_DB", 0),
			PoolSize: getEnvInt("REDIS_POOL_SIZE", 10),
		},
		Kafka: kafka.Config{
			Brokers: []string{getEnv("KAFKA_BROKER", "localhost:9092")},
			GroupID: getEnv("KAFKA_GROUP_ID", "marketplace-sync"),
		},
	}, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return fallback
}

func getEnvDuration(key string, fallback time.Duration) (time.Duration, error) {
	v := os.Getenv(key)
	if v == "" {
		return fallback, nil
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		return 0, fmt.Errorf("env %s: expected duration (e.g. 5m), got %q", key, v)
	}
	return d, nil
}
