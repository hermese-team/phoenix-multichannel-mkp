package marketplace

import (
	"context"

	"github.com/go-resty/resty/v2"

	"github.com/ascend/phoenix-multichannel-mkp/config"
	entity "github.com/ascend/phoenix-multichannel-mkp/domain"
	"github.com/ascend/phoenix-multichannel-mkp/internal/service"
)

type TikTokGateway struct{ baseGateway }

func NewTikTokGateway(rc *resty.Client, cfg config.ChannelEndpoint) *TikTokGateway {
	return &TikTokGateway{baseGateway{
		channel: entity.ChannelTikTok,
		rc:      rc,
		baseURL: cfg.BaseURL,
		apiKey:  cfg.APIKey,
	}}
}

// tiktokProduct is TikTok Shop's product payload shape.
type tiktokProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	StockQty    int     `json:"stock_qty"`
}

func (g *TikTokGateway) PushProduct(ctx context.Context, p *entity.Product) (string, error) {
	return g.post(ctx, "/api/products", tiktokProduct{
		Title:       p.Name,
		Description: p.Description,
		Price:       p.Price,
		StockQty:    p.Stock,
	})
}

var _ service.MarketplaceGateway = (*TikTokGateway)(nil)
