package server

import (
	"bytes"
	"crypto/subtle"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/okdev/marketplace-sync/pkg/signer"
)

const maxWebhookBodyBytes = 1 * 1024 * 1024 // 1 MB — Shopee payloads are well under 10 KB

// shopeeWebhookAuth verifies Shopee push notification authenticity.
//
// Shopee signs each push with:
//
//	HMAC-SHA256(key=appSecret, msg=partnerID + requestPath + rawBody)
//
// and puts the hex result in the Authorization header.
// Set SHOPEE_WEBHOOK_VERIFY=false to skip verification (sandbox/dev only).
func shopeeWebhookAuth(partnerID int64, appSecret string, verify bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Enforce body size limit before reading — rejects oversized payloads early.
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxWebhookBodyBytes)

		// Buffer raw body before ShouldBindJSON consumes it.
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusRequestEntityTooLarge, gin.H{"error": "request body too large"})
			return
		}
		// Restore for downstream handlers.
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		if !verify {
			c.Next()
			return
		}

		got := c.GetHeader("Authorization")
		if got == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
			return
		}

		msg := fmt.Sprintf("%d%s%s", partnerID, c.Request.URL.Path, string(body))
		expected := signer.Sign(appSecret, msg)

		if subtle.ConstantTimeCompare([]byte(got), []byte(expected)) != 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid webhook signature"})
			return
		}

		c.Next()
	}
}
