package server

import (
	"time"

	"github.com/gin-gonic/gin"
	applogger "github.com/okdev/marketplace-sync/pkg/logger"
)

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		applogger.Info("http request",
			"component", "server",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"latency_ms", time.Since(start).Milliseconds(),
		)
	}
}
