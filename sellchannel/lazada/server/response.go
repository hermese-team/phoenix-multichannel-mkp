package server

import (
	"github.com/gin-gonic/gin"

	"github.com/okdev/marketplace-sync/internal/apperr"
	"github.com/okdev/marketplace-sync/pkg/httpserver"
	applog "github.com/okdev/marketplace-sync/pkg/logger"
)

// respondError maps err to an AppError (sentinel-aware) and writes it via the
// shared httpserver envelope. 5xx are logged as server faults; 4xx are the
// caller's problem and left for them to see in the response.
func respondError(c *gin.Context, err error) {
	ae := apperr.From(err)
	if ae.HTTPStatusCode >= 500 {
		applog.Error("request failed",
			"path", c.Request.URL.Path,
			"code", ae.Code,
			"error", err,
		)
	}
	httpserver.NewErrorResponse(c, ae)
}
