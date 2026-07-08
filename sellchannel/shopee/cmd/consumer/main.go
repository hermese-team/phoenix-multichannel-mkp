package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"

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

	productConsumer, err := shopee.NewConsumer(cfg.Shopee, cfg.Postgres, cfg.Kafka)
	if err != nil {
		log.Fatalf("init shopee product consumer: %v", err)
	}

	classifierConsumer, err := shopee.NewClassifierConsumer(cfg.Kafka)
	if err != nil {
		log.Fatalf("init shopee classifier consumer: %v", err)
	}

	orderConsumer, err := shopee.NewOrderConsumer(cfg.Shopee, cfg.Postgres, cfg.Redis, cfg.Kafka)
	if err != nil {
		log.Fatalf("init shopee order consumer: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() { <-quit; cancel() }()

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error { return productConsumer.Start(ctx) })
	g.Go(func() error { return classifierConsumer.Start(ctx) })
	g.Go(func() error { return orderConsumer.Start(ctx) })

	log.Println("starting shopee consumers")
	if err := g.Wait(); err != nil {
		log.Fatalf("consumer error: %v", err)
	}
}
