package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/ascend/phoenix-multichannel-mkp/config"
	"github.com/ascend/phoenix-multichannel-mkp/internal/app"
	"github.com/ascend/phoenix-multichannel-mkp/pkg/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "load config: %v\n", err)
		os.Exit(1)
	}

	log, lp, err := logger.New(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "init logger: %v\n", err)
		os.Exit(1)
	}
	defer log.Sync()                        //nolint:errcheck
	defer lp.Shutdown(context.Background()) //nolint:errcheck

	ctx := context.Background()
	application, err := app.New(ctx, cfg, log)
	if err != nil {
		log.Fatal("init app", zap.Error(err))
	}

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := application.Start(); err != nil {
			log.Error("server error", zap.Error(err))
			quit <- syscall.SIGTERM
		}
	}()

	<-quit
	log.Info("shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := application.Shutdown(shutdownCtx); err != nil {
		log.Error("shutdown error", zap.Error(err))
	}
	log.Info("server stopped")
}
