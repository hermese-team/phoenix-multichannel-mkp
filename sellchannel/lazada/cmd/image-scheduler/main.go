package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/okdev/marketplace-sync/config"
	"github.com/okdev/marketplace-sync/pkg/logger"
	"github.com/okdev/marketplace-sync/sellchannel/lazada"
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

	sched, err := lazada.NewImageScheduler(cfg.Lazada, cfg.Postgres, cfg.Redis)
	if err != nil {
		log.Fatalf("init lazada image scheduler: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		cancel()
	}()

	if err := sched.Start(ctx); err != nil {
		log.Fatalf("image scheduler error: %v", err)
	}
}
