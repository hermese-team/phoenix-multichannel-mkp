package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	orderUC "github.com/okdev/marketplace-sync/internal/usecase/order"
)

type Server struct {
	engine       *gin.Engine
	processOrder *orderUC.ProcessWebhookUsecase
}

func New(processOrder *orderUC.ProcessWebhookUsecase) *Server {
	s := &Server{
		engine:       gin.New(),
		processOrder: processOrder,
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.engine.Use(gin.Recovery(), logger())

	s.engine.GET("/health", s.health)

	oauth := s.engine.Group("/oauth")
	{
		oauth.GET("/callback", s.oauthCallback)
	}

	webhook := s.engine.Group("/webhook")
	{
		webhook.POST("/order", s.handleOrderWebhook)
		webhook.POST("/product", s.handleProductWebhook)
	}
}

func (s *Server) Run(addr string) error {
	return s.engine.Run(addr)
}

func (s *Server) Handler() http.Handler {
	return s.engine
}
