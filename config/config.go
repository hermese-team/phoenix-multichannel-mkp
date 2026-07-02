package config

import (
	"os"
	"strconv"

	"github.com/okdev/marketplace-sync/internal/infrastructure/kafka"
	"github.com/okdev/marketplace-sync/internal/infrastructure/mysql"
	"github.com/okdev/marketplace-sync/internal/infrastructure/redis"
	"github.com/okdev/marketplace-sync/sellchannel/shopee"
)

type App struct {
	Env  string
	Port string
}

type Config struct {
	App    App
	Shopee shopee.Config
	MySQL  mysql.Config
	Redis  redis.Config
	Kafka  kafka.Config
}

func Load() *Config {
	return &Config{
		App: App{
			Env:  getEnv("APP_ENV", "development"),
			Port: getEnv("APP_PORT", "8080"),
		},
		Shopee: shopee.Config{
			PartnerID: int64(getEnvInt("SHOPEE_PARTNER_ID", 0)),
			AppKey:    os.Getenv("SHOPEE_APP_KEY"),
			AppSecret: os.Getenv("SHOPEE_APP_SECRET"),
			BaseURL:   getEnv("SHOPEE_BASE_URL", "https://partner.shopeemobile.com"),
		},
		MySQL: mysql.Config{
			DSN:          os.Getenv("MYSQL_DSN"),
			MaxOpenConns: getEnvInt("MYSQL_MAX_OPEN_CONNS", 25),
			MaxIdleConns: getEnvInt("MYSQL_MAX_IDLE_CONNS", 5),
		},
		Redis: redis.Config{
			Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       getEnvInt("REDIS_DB", 0),
		},
		Kafka: kafka.Config{
			Brokers: []string{getEnv("KAFKA_BROKER", "localhost:9092")},
		},
	}
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
