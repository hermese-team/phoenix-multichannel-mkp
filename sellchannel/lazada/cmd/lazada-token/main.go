// lazada-token seeds the Lazada OAuth tokens into Redis (one-time bootstrap).
// After this, the scheduler reads/refreshes tokens from Redis and .env no longer
// needs to hold them.
//
// Usage:
//
//	go run ./sellchannel/lazada/cmd/lazada-token -refresh=<refresh> [-access=<access>]
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/okdev/marketplace-sync/config"
	redisAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/redis"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/tokenstore"
)

func main() {
	access := flag.String("access", "", "Lazada access token (optional; obtained via refresh if empty)")
	refresh := flag.String("refresh", "", "Lazada refresh token (required)")
	flag.Parse()

	if *refresh == "" {
		log.Fatal("missing -refresh")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	rdb, err := redisAdapter.New(cfg.Redis)
	if err != nil {
		log.Fatalf("connect redis: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := tokenstore.New(rdb).Set(ctx, *access, *refresh, 0); err != nil {
		log.Fatalf("store tokens: %v", err)
	}
	log.Println("lazada tokens written to redis")
}
