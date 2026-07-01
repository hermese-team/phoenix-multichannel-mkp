package handler

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.uber.org/zap"
)

type Config struct {
	Port            string
	ServiceName     string
	ShutdownTimeout time.Duration
}

type Server struct {
	httpServer *http.Server
	engine     *gin.Engine
	log        *zap.Logger
}

func New(cfg Config, log *zap.Logger, productHandler *ProductHandler, orderHandler *OrderHandler) *Server {
	if cfg.ServiceName == "" {
		cfg.ServiceName = "phoenix-multichannel-mkp"
	}

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(otelgin.Middleware(cfg.ServiceName)) // OTel trace per HTTP request
	engine.Use(zapLogger(log))

	v1 := engine.Group("/api/v1")
	productHandler.RegisterRoutes(v1)
	orderHandler.RegisterRoutes(v1)

	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return &Server{
		engine: engine,
		log:    log,
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
	s.log.Info("starting server", zap.String("addr", s.httpServer.Addr))
	if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.log.Info("shutting down server")
	return s.httpServer.Shutdown(ctx)
}

// zapLogger logs each request with structured zap fields.
func zapLogger(log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Info("request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", time.Since(start)),
			zap.String("client_ip", c.ClientIP()),
		)
	}
}
