package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"git.amaze-x.com/phoenix/channel-adapter-shopee/config"
	"git.amaze-x.com/phoenix/channel-adapter-shopee/internal/infrastructure"
	"git.amaze-x.com/phoenix/channel-adapter-shopee/internal/shopee"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.uber.org/zap"
)

func main() {
	cfg := config.MustLoad()

	// Logger (with OTel log bridge when telemetry endpoint is set)
	log, lp, err := infrastructure.NewLogger(cfg)
	if err != nil {
		panic("logger init failed: " + err.Error())
	}
	defer log.Sync()
	defer lp.Shutdown(context.Background()) //nolint:errcheck

	if cfg.AppEnv == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	log.Info("channel-adapter-shopee server starting", zap.String("config", config.RedactedJSON(cfg)))

	// Tracer
	tp, err := infrastructure.NewTracer(cfg)
	if err != nil {
		log.Fatal("tracer init failed", zap.Error(err))
	}

	// Redis
	rdb, err := infrastructure.NewRedis(cfg)
	if err != nil {
		log.Fatal("redis init failed", zap.Error(err))
	}

	// Kafka producer — acks=all per architecture requirement
	producer, err := infrastructure.NewKafkaProducer(cfg)
	if err != nil {
		log.Fatal("kafka producer init failed", zap.Error(err))
	}
	defer producer.Close()

	// Token store — seed initial tokens from config (Vault-injected at startup).
	signer := shopee.NewSigner(cfg.Shopee.PartnerID, cfg.Shopee.PartnerKey)
	tokenStore := shopee.NewRedisTokenStore(rdb, cfg.Shopee.BaseURL, signer,
		cfg.Shopee.PartnerID, cfg.Shopee.ShopID)

	if cfg.Shopee.AccessToken != "" {
		if err := tokenStore.Seed(context.Background(),
			cfg.Shopee.AccessToken, cfg.Shopee.RefreshToken,
			4*time.Hour,
		); err != nil {
			log.Warn("seeding initial tokens failed — will attempt refresh on first call", zap.Error(err))
		}
	}

	// Router
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(otelgin.Middleware("shopee-adapter"))

	r.GET("/system/health", healthHandler)
	r.GET("/auth/shopee", authURLHandler(cfg, signer))
	r.GET("/auth/shopee/callback", callbackHandler(tokenStore, log))

	// Shopee push notification webhook.
	// Shopee POSTs order/logistics events here; we verify the signature,
	// archive the raw payload to Kafka, and reply 200 immediately.
	r.POST("/webhook/shopee", webhookHandler(cfg, producer, log))

	addr := fmt.Sprintf(":%d", cfg.AppPort)
	srv := &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Info("server listening", zap.String("addr", addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server error", zap.Error(err))
		}
	}()

	<-quit
	log.Info("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("server shutdown error", zap.Error(err))
	}
	if err := tp.Shutdown(ctx); err != nil {
		log.Error("tracer shutdown error", zap.Error(err))
	}
	log.Info("server stopped")
}

// healthHandler returns a simple liveness probe response.
func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// authURLHandler redirects the operator's browser to the Shopee OAuth consent page.
// Used during initial setup to obtain the first access/refresh token pair.
func authURLHandler(cfg config.Config, signer *shopee.Signer) gin.HandlerFunc {
	return func(c *gin.Context) {
		authURL := shopee.GetAuthURL(cfg.Shopee.BaseURL, cfg.Shopee.PartnerID, signer, cfg.Shopee.RedirectURL)
		c.Redirect(http.StatusFound, authURL)
	}
}

// callbackHandler receives the Shopee OAuth callback, exchanges the code for
// access/refresh tokens, and stores them in Redis so the worker can poll immediately.
func callbackHandler(tokenStore *shopee.RedisTokenStore, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Query("code")
		if code == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing code"})
			return
		}

		if err := tokenStore.ExchangeCode(c.Request.Context(), code); err != nil {
			log.Error("token exchange failed", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		log.Info("shopee tokens exchanged and stored in redis")
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "tokens stored — worker is ready",
		})
	}
}
