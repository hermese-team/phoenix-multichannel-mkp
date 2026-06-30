package service

import (
	"context"

	entity "github.com/ascend/phoenix-multichannel-mkp/domain"
)

// CreateProductInput is the request payload for creating a product.
type CreateProductInput struct {
	SellerID    string
	Name        string
	Description string
	Price       float64
	Stock       int
	Channels    []entity.Channel
}

// CreateProductOutput is the result of creating a product.
type CreateProductOutput struct {
	Product *entity.Product
}

// ListProductsInput holds the optional criteria for listing products.
type ListProductsInput struct {
	SellerID string
	Channel  entity.Channel
	Status   entity.ProductStatus
	Page     int
	Limit    int
}

// ListProductsOutput is a paginated list of products.
type ListProductsOutput struct {
	Products []*entity.Product
	Total    int
	Page     int
	Limit    int
}

// CreateOrderItemInput is a single line item of a new order.
type CreateOrderItemInput struct {
	ProductID string
	Quantity  int
}

// CreateOrderInput is the request payload for placing an order.
type CreateOrderInput struct {
	BuyerID string
	Channel entity.Channel
	Items   []CreateOrderItemInput
}

// CreateOrderOutput is the result of placing an order.
type CreateOrderOutput struct {
	Order *entity.Order
}

// ProductService is the application port for product use cases.
type ProductService interface {
	Create(ctx context.Context, input CreateProductInput) (*CreateProductOutput, error)
	Get(ctx context.Context, id string) (*entity.Product, error)
	List(ctx context.Context, input ListProductsInput) (*ListProductsOutput, error)
}

// OrderService is the application port for order use cases.
type OrderService interface {
	Create(ctx context.Context, input CreateOrderInput) (*CreateOrderOutput, error)
}

// EventPublisher is the port for emitting domain events to a message broker.
type EventPublisher interface {
	Publish(ctx context.Context, topic, key string, payload []byte) error
}

// MarketplaceGateway is the outbound port for a single external sales channel
// (TikTok, Lazada, Shopee, ...). Implementations live in infrastructure.
type MarketplaceGateway interface {
	Channel() entity.Channel
	PushProduct(ctx context.Context, product *entity.Product) (externalID string, err error)
}

// MarketplaceRegistry resolves the gateway responsible for a given channel.
type MarketplaceRegistry interface {
	Get(channel entity.Channel) (MarketplaceGateway, error)
}
