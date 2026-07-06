package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/okdev/marketplace-sync/config"
	"github.com/okdev/marketplace-sync/pkg/logger"
	"github.com/okdev/marketplace-sync/sellchannel/shopee"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}
	if err := logger.Init(cfg.App.LogLevel); err != nil {
		log.Fatalf("init logger: %v", err)
	}
	defer logger.Sync()

	c, err := shopee.NewConsumer(cfg.Shopee, cfg.Postgres, cfg.Kafka)
	if err != nil {
		log.Fatalf("init shopee consumer: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() { <-quit; cancel() }()

	log.Println("starting shopee consumer")
	if err := c.Start(ctx); err != nil {
		log.Fatalf("consumer error: %v", err)
	}
}
