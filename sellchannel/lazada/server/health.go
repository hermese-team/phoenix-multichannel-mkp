package server

import (
	"github.com/gin-gonic/gin"

	"github.com/okdev/marketplace-sync/pkg/httpserver"
)

func (s *Server) health(c *gin.Context) {
	httpserver.NewSuccessResponse(c, gin.H{"status": "ok"})
}
