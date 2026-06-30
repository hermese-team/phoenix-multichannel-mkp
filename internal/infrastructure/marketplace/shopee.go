package marketplace

import (
	"context"

	"github.com/go-resty/resty/v2"

	"github.com/ascend/phoenix-multichannel-mkp/config"
	entity "github.com/ascend/phoenix-multichannel-mkp/domain"
	"github.com/ascend/phoenix-multichannel-mkp/internal/service"
)

type ShopeeGateway struct{ baseGateway }

func NewShopeeGateway(rc *resty.Client, cfg config.ChannelEndpoint) *ShopeeGateway {
	return &ShopeeGateway{baseGateway{
		channel: entity.ChannelShopee,
		rc:      rc,
		baseURL: cfg.BaseURL,
		apiKey:  cfg.APIKey,
	}}
}

// shopeeProduct is Shopee's item payload shape.
type shopeeProduct struct {
	ItemName    string  `json:"item_name"`
	Description string  `json:"description"`
	OriginPrice float64 `json:"original_price"`
	StockCount  int     `json:"stock"`
}

func (g *ShopeeGateway) PushProduct(ctx context.Context, p *entity.Product) (string, error) {
	return g.post(ctx, "/api/v2/product/add_item", shopeeProduct{
		ItemName:    p.Name,
		Description: p.Description,
		OriginPrice: p.Price,
		StockCount:  p.Stock,
	})
}

var _ service.MarketplaceGateway = (*ShopeeGateway)(nil)
