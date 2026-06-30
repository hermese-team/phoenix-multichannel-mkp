package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ascend/phoenix-multichannel-mkp/config"
	"github.com/ascend/phoenix-multichannel-mkp/internal/app"
	"github.com/ascend/phoenix-multichannel-mkp/pkg/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		logger.Error("load config", "error", err)
		os.Exit(1)
	}

	if err := logger.Init(cfg.App.LogLevel); err != nil {
		logger.Error("init logger", "error", err)
		os.Exit(1)
	}
	defer logger.Sync()

	ctx := context.Background()
	application, err := app.New(ctx, cfg)
	if err != nil {
		logger.Error("init app", "error", err)
		os.Exit(1)
	}

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := application.Start(); err != nil {
			logger.Error("server error", "error", err)
			quit <- syscall.SIGTERM
		}
	}()

	<-quit
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := application.Shutdown(shutdownCtx); err != nil {
		logger.Error("shutdown error", "error", err)
	}
}
