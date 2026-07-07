package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	kafkaAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/kafka"
	pgAdapter     "github.com/okdev/marketplace-sync/internal/infrastructure/postgres"
	redisAdapter  "github.com/okdev/marketplace-sync/internal/infrastructure/redis"
	orderUC       "github.com/okdev/marketplace-sync/internal/usecase/order"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/client"
	"github.com/okdev/marketplace-sync/sellchannel/shopee/tokenstore"
)

type Server struct {
	engine        *gin.Engine
	processOrder  *orderUC.ProcessWebhookUsecase
	shopeeClient  *client.Client
	tokens        *tokenstore.Store
	producer      *kafkaAdapter.Producer
	partnerID     int64
	appSecret     string
	webhookVerify bool
}

type Config struct {
	PartnerID     int64
	AppSecret     string
	WebhookVerify bool
}

func New(processOrder *orderUC.ProcessWebhookUsecase, shopeeClient *client.Client, rdb *redisAdapter.Client, tokenRepo *pgAdapter.ShopeeTokenRepository, producer *kafkaAdapter.Producer, cfg Config) *Server {
	s := &Server{
		engine:        gin.New(),
		processOrder:  processOrder,
		shopeeClient:  shopeeClient,
		tokens:        tokenstore.New(tokenRepo, rdb),
		producer:      producer,
		partnerID:     cfg.PartnerID,
		appSecret:     cfg.AppSecret,
		webhookVerify: cfg.WebhookVerify,
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.engine.Use(gin.Recovery(), logger())

	s.engine.GET("/health", s.health)

	oauth := s.engine.Group("/oauth")
	{
		oauth.GET("/authorize", s.oauthAuthorize)
		oauth.GET("/callback", s.oauthCallback)
	}

	webhook := s.engine.Group("/webhook")
	webhook.Use(shopeeWebhookAuth(s.partnerID, s.appSecret, s.webhookVerify))
	{
		webhook.POST("/order", s.handleOrderWebhook)
		webhook.POST("/product", s.handleProductWebhook)
	}

	// Debug endpoints — fetch from Shopee API and publish to Kafka (non-production use)
	debug := s.engine.Group("/debug")
	{
		debug.GET("/order", s.debugFetchOrder)
	}
}

func (s *Server) Run(addr string) error {
	return s.engine.Run(addr)
}

func (s *Server) Handler() http.Handler {
	return s.engine
}
