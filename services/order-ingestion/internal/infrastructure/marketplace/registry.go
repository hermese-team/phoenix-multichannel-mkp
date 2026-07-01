package marketplace

import (
	entity "github.com/ascend/phoenix-multichannel-mkp/domain"
	"github.com/ascend/phoenix-multichannel-mkp/internal/service"
	domainerrors "github.com/ascend/phoenix-multichannel-mkp/pkg/errors"
)

// Registry resolves a MarketplaceGateway by channel. It implements
// service.MarketplaceRegistry.
type Registry struct {
	gateways map[entity.Channel]service.MarketplaceGateway
}

// NewRegistry indexes the given gateways by their channel.
func NewRegistry(gateways ...service.MarketplaceGateway) *Registry {
	m := make(map[entity.Channel]service.MarketplaceGateway, len(gateways))
	for _, g := range gateways {
		m[g.Channel()] = g
	}
	return &Registry{gateways: m}
}

func (r *Registry) Get(channel entity.Channel) (service.MarketplaceGateway, error) {
	g, ok := r.gateways[channel]
	if !ok {
		return nil, domainerrors.New(
			"UNSUPPORTED_CHANNEL",
			"no marketplace gateway for channel "+string(channel),
			domainerrors.ErrInvalidInput,
		)
	}
	return g, nil
}

var _ service.MarketplaceRegistry = (*Registry)(nil)
