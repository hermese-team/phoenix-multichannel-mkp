package main

import (
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

	// marketplace = shoppee.NewMarket()
	// productUsecase.New(config, marketplace)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("shopee server listening on :%s", cfg.App.Port)
		if err := app.Server().Run(":" + cfg.App.Port); err != nil {
			log.Fatalf("server error: %v", err)
		}
	}()

	<-quit
	log.Println("shutting down shopee server")
}
