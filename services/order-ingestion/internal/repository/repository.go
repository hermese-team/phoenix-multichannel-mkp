package repository

import (
	"context"

	entity "github.com/ascend/phoenix-multichannel-mkp/domain"
)

// OrderFilter holds the optional criteria for listing orders.
// Zero-value fields are ignored by implementations.
type OrderFilter struct {
	BuyerID  string
	SellerID string
	Status   entity.OrderStatus
	Channel  entity.Channel
	Page     int
	Limit    int
}

// ProductFilter holds the optional criteria for listing products.
// Zero-value fields are ignored by implementations.
type ProductFilter struct {
	SellerID string
	Channel  entity.Channel
	Status   entity.ProductStatus
	Page     int
	Limit    int
}

// OrderRepository is the persistence port for orders.
type OrderRepository interface {
	Create(ctx context.Context, order *entity.Order) error
	GetByID(ctx context.Context, id string) (*entity.Order, error)
	Update(ctx context.Context, order *entity.Order) error
	List(ctx context.Context, filter OrderFilter) ([]*entity.Order, int, error)
}

// ProductRepository is the persistence port for products.
type ProductRepository interface {
	Create(ctx context.Context, product *entity.Product) error
	GetByID(ctx context.Context, id string) (*entity.Product, error)
	Update(ctx context.Context, product *entity.Product) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filter ProductFilter) ([]*entity.Product, int, error)
}