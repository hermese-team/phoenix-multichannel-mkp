package main

import (
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

	srv, err := shopee.NewServer(cfg.Postgres)
	if err != nil {
		log.Fatalf("init shopee server: %v", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("shopee server listening on :%s", cfg.App.Port)
		if err := srv.Run(":" + cfg.App.Port); err != nil {
			log.Fatalf("server error: %v", err)
		}
	}()

	<-quit
	log.Println("shutting down shopee server")
}
