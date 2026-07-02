package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/okdev/marketplace-sync/config"
	"github.com/okdev/marketplace-sync/sellchannel/shopee"
)

func main() {
	cfg := config.Load()

	app, err := shopee.New(cfg.Shopee, cfg.MySQL, cfg.Redis, cfg.Kafka)
	if err != nil {
		log.Fatalf("init shopee: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		cancel()
	}()

	log.Println("starting shopee scheduler")
	app.Scheduler().Start(ctx)
}
