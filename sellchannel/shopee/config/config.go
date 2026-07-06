package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type App struct {
	Env      string
	Port     string
	LogLevel string
}

type Shopee struct {
	PartnerID int64
	AppKey    string
	AppSecret string
	BaseURL   string
}

type Postgres struct {
	DSN             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

type Redis struct {
	Addr     string
	Password string
	DB       int
	PoolSize int
}

type Kafka struct {
	Brokers []string
	GroupID string
}

type Config struct {
	App      App
	Shopee   Shopee
	Postgres Postgres
	Redis    Redis
	Kafka    Kafka
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
		Shopee: Shopee{
			PartnerID: int64(getEnvInt("SHOPEE_PARTNER_ID", 0)),
			AppKey:    os.Getenv("SHOPEE_APP_KEY"),
			AppSecret: os.Getenv("SHOPEE_APP_SECRET"),
			BaseURL:   getEnv("SHOPEE_BASE_URL", "https://partner.shopeemobile.com"),
		},
		Postgres: Postgres{
			DSN:             getEnv("POSTGRES_DSN", "postgres://postgres:postgres@localhost:5432/shopee?sslmode=disable"),
			MaxOpenConns:    getEnvInt("POSTGRES_MAX_OPEN_CONNS", 25),
			MaxIdleConns:    getEnvInt("POSTGRES_MAX_IDLE_CONNS", 5),
			ConnMaxLifetime: pgLifetime,
		},
		Redis: Redis{
			Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       getEnvInt("REDIS_DB", 0),
			PoolSize: getEnvInt("REDIS_POOL_SIZE", 10),
		},
		Kafka: Kafka{
			Brokers: []string{getEnv("KAFKA_BROKER", "localhost:9092")},
			GroupID: getEnv("KAFKA_GROUP_ID", "shopee-channel-adapter"),
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
