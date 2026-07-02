package order

import "context"

type Repository interface {
	FindBySellChannelID(ctx context.Context, marketplaceID, marketplaceType string) (*Order, error)
	Save(ctx context.Context, o *Order) error
	Update(ctx context.Context, o *Order) error
}

type EventPublisher interface {
	OrderProcessed(ctx context.Context, o *Order) error
}
