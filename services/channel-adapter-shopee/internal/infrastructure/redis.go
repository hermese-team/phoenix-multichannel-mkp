package infrastructure

import (
	"context"
	"fmt"
	"time"

	"git.amaze-x.com/phoenix/channel-adapter-shopee/config"
	"github.com/redis/go-redis/v9"
)

const redisPingTimeout = 5 * time.Second

// NewRedis returns a UniversalClient supporting standalone and cluster mode.
// Mode is set via config.redis.mode: "standalone" | "cluster"
func NewRedis(cfg config.Config) (redis.UniversalClient, error) {
	var client redis.UniversalClient

	switch cfg.Redis.Mode {
	case "cluster":
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    cfg.Redis.Addrs,
			Username: cfg.Redis.User,
			Password: cfg.Redis.Password,
		})
	default: // "standalone"
		addr := "localhost:6379"
		if len(cfg.Redis.Addrs) > 0 {
			addr = cfg.Redis.Addrs[0]
		}
		client = redis.NewClient(&redis.Options{
			Addr:     addr,
			Username: cfg.Redis.User,
			Password: cfg.Redis.Password,
			DB:       cfg.Redis.DB,
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), redisPingTimeout)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis ping failed: %w", err)
	}
	return client, nil
}
