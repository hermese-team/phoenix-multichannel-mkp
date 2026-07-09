package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/okdev/marketplace-sync/pkg/logger"
)

// handleShopAuthCanceled handles Shopee push code 2 (shop_authorization_canceled_push).
// Shopee sends this when a seller revokes the app's access from Seller Center.
// We clear the token from Redis and Postgres immediately so no further API calls
// are attempted for this shop.
func (s *Server) handleShopAuthCanceled(c *gin.Context, shopID int64) {
	ctx := c.Request.Context()

	logger.InfoContext(ctx, "shop authorization canceled",
		"event", "webhook.shop_auth.canceled",
		"shop_id", shopID,
	)

	if err := s.tokens.Revoke(ctx, shopID); err != nil {
		logger.ErrorContext(ctx, "revoke token failed",
			"event", "webhook.shop_auth.revoke_error",
			"shop_id", shopID,
			"error", err,
		)
		// Still return 200 — Shopee must get 2xx or it retries indefinitely.
		// The token will expire naturally even if we fail to delete it now.
	} else {
		logger.InfoContext(ctx, "shop token revoked",
			"event", "webhook.shop_auth.revoked",
			"shop_id", shopID,
		)
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
