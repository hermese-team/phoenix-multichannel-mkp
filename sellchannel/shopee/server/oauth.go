package server

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// GET /oauth/authorize
// Redirects the merchant to Shopee's partner authorization page.
func (s *Server) oauthAuthorize(c *gin.Context) {
	redirect := os.Getenv("SHOPEE_OAUTH_REDIRECT_URL")
	if redirect == "" {
		redirect = "http://localhost:8080/oauth/callback"
	}
	authURL := s.shopeeClient.AuthURL(redirect)
	c.Redirect(http.StatusFound, authURL)
}

// GET /oauth/callback?code=<code>&shop_id=<shop_id>
// Exchanges the authorization code for tokens and stores them in Redis.
func (s *Server) oauthCallback(c *gin.Context) {
	var req OAuthCallbackRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := s.shopeeClient.GetAccessToken(c.Request.Context(), req.Code, req.ShopID)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	if resp.Error != "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": resp.Error, "message": resp.Message})
		return
	}

	// Use shop_id from the callback query param — Shopee does not always return it in the response body.
	shopID := req.ShopID
	if shopID == 0 {
		shopID = resp.ShopID
	}

	expireIn := resp.ExpireIn
	if expireIn <= 0 {
		expireIn = int64(4 * time.Hour / time.Second)
	}
	expiredAt := time.Now().Add(time.Duration(expireIn) * time.Second)
	if err := s.tokens.Set(c.Request.Context(), shopID, resp.AccessToken, resp.RefreshToken, expiredAt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "store token: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"shop_id":    shopID,
		"expires_in": resp.ExpireIn,
		"message":    "token stored successfully",
	})
}
