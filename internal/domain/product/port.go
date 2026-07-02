package product

import "context"

type Repository interface {
	FindByID(ctx context.Context, id int64) (*Product, error)
	FindByShopID(ctx context.Context, shopID int64, limit, offset int) ([]*Product, error)
	FindPendingSync(ctx context.Context, marketplaceType string, limit int) ([]*Product, error)
	Save(ctx context.Context, p *Product) error
	Update(ctx context.Context, p *Product) error
}

type SellChannel interface {
	PublishProduct(ctx context.Context, shopID int64, p *Product) (itemID int64, err error)
	UpdateProduct(ctx context.Context, shopID int64, p *Product) error
}

type EventPublisher interface {
	ProductPublished(ctx context.Context, p *Product) error
	ProductUpdated(ctx context.Context, p *Product) error
}
