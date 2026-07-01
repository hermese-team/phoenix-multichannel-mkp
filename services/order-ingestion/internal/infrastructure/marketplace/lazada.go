package marketplace

import (
	"context"

	"github.com/go-resty/resty/v2"

	"github.com/ascend/phoenix-multichannel-mkp/config"
	entity "github.com/ascend/phoenix-multichannel-mkp/domain"
	"github.com/ascend/phoenix-multichannel-mkp/internal/service"
)

type LazadaGateway struct{ baseGateway }

func NewLazadaGateway(rc *resty.Client, cfg config.ChannelEndpoint) *LazadaGateway {
	return &LazadaGateway{baseGateway{
		channel: entity.ChannelLazada,
		rc:      rc,
		baseURL: cfg.BaseURL,
		apiKey:  cfg.APIKey,
	}}
}

// lazadaProduct is Lazada's product payload shape.
type lazadaProduct struct {
	Name      string  `json:"name"`
	ShortDesc string  `json:"short_description"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	SellerSKU string  `json:"seller_sku"`
}

func (g *LazadaGateway) PushProduct(ctx context.Context, p *entity.Product) (string, error) {
	return g.post(ctx, "/product/create", lazadaProduct{
		Name:      p.Name,
		ShortDesc: p.Description,
		Price:     p.Price,
		Quantity:  p.Stock,
		SellerSKU: p.ID,
	})
}

var _ service.MarketplaceGateway = (*LazadaGateway)(nil)
