package handler

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ascend/phoenix-multichannel-mkp/internal/handler/middleware"
	"github.com/ascend/phoenix-multichannel-mkp/pkg/logger"
)

type Config struct {
	Port            string
	ShutdownTimeout time.Duration
}

type Server struct {
	httpServer *http.Server
	engine     *gin.Engine
}

func New(cfg Config, productHandler *ProductHandler, orderHandler *OrderHandler) *Server {
	engine := gin.New()
	engine.Use(middleware.Logger(), middleware.Recovery())

	v1 := engine.Group("/api/v1")
	productHandler.RegisterRoutes(v1)
	orderHandler.RegisterRoutes(v1)

	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return &Server{
		engine: engine,
		httpServer: &http.Server{
			Addr:         ":" + cfg.Port,
			Handler:      engine,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	logger.Info("starting server", "addr", s.httpServer.Addr)
	if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	logger.Info("shutting down server")
	return s.httpServer.Shutdown(ctx)
}
