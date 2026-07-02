package server

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	kafkaAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/kafka"
	redisAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/redis"
	orderUC "github.com/okdev/marketplace-sync/internal/usecase/order"
)

type Server struct {
	engine       *gin.Engine
	processOrder *orderUC.ProcessWebhookUsecase
	redis        *redisAdapter.Client
	producer     *kafkaAdapter.Producer
	idemPrefix   string
	topic        string
}

func New(processOrder *orderUC.ProcessWebhookUsecase, redis *redisAdapter.Client, producer *kafkaAdapter.Producer) *Server {
	idemPrefix := os.Getenv("LAZADA_WEBHOOK_IDEM_PREFIX")
	if idemPrefix == "" {
		idemPrefix = "lazada:webhook:idem"
	}
	topic := os.Getenv("LAZADA_WEBHOOK_TOPIC")
	if topic == "" {
		topic = "lazada.order.webhook"
	}
	s := &Server{
		engine:       gin.New(),
		processOrder: processOrder,
		redis:        redis,
		producer:     producer,
		idemPrefix:   idemPrefix,
		topic:        topic,
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.engine.Use(gin.Recovery(), logger())

	s.engine.GET("/health", s.health)

	webhook := s.engine.Group("/webhook")
	{
		webhook.POST("/order", s.handleOrderWebhook)
	}
}

func (s *Server) Run(addr string) error {
	return s.engine.Run(addr)
}

func (s *Server) Handler() http.Handler {
	return s.engine
}
