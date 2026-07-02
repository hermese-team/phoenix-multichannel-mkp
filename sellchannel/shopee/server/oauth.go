package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) oauthCallback(c *gin.Context) {
	var req OAuthCallbackRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO: exchange code for access token via client.GetAccessToken and store it
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
