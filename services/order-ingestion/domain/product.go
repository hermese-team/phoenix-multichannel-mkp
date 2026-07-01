package entity

import (
	"time"

	domainerrors "github.com/ascend/phoenix-multichannel-mkp/pkg/errors"
)

type ProductStatus string

const (
	ProductStatusActive   ProductStatus = "active"
	ProductStatusInactive ProductStatus = "inactive"
	ProductStatusDraft    ProductStatus = "draft"
)

type Channel string

const (
	ChannelWeb    Channel = "web"
	ChannelMobile Channel = "mobile"
	ChannelAPI    Channel = "api"
	ChannelTikTok Channel = "tiktok"
	ChannelLazada Channel = "lazada"
	ChannelShopee Channel = "shopee"
)

type Product struct {
	ID          string
	SellerID    string
	Name        string
	Description string
	Price       float64
	Stock       int
	Status      ProductStatus
	Channels    []Channel
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewProduct(sellerID, name, description string, price float64, stock int, channels []Channel) (*Product, error) {
	if name == "" {
		return nil, domainerrors.New("INVALID_NAME", "product name is required", domainerrors.ErrInvalidInput)
	}
	if price < 0 {
		return nil, domainerrors.New("INVALID_PRICE", "price must be non-negative", domainerrors.ErrInvalidInput)
	}
	if stock < 0 {
		return nil, domainerrors.New("INVALID_STOCK", "stock must be non-negative", domainerrors.ErrInvalidInput)
	}
	if len(channels) == 0 {
		return nil, domainerrors.New("INVALID_CHANNELS", "at least one channel is required", domainerrors.ErrInvalidInput)
	}
	now := time.Now()
	return &Product{
		SellerID:    sellerID,
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
		Status:      ProductStatusDraft,
		Channels:    channels,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (p *Product) Activate() {
	p.Status = ProductStatusActive
	p.UpdatedAt = time.Now()
}

func (p *Product) Deactivate() {
	p.Status = ProductStatusInactive
	p.UpdatedAt = time.Now()
}

func (p *Product) DeductStock(qty int) error {
	if p.Stock < qty {
		return domainerrors.New("INSUFFICIENT_STOCK", "insufficient stock", domainerrors.ErrInvalidInput)
	}
	p.Stock -= qty
	p.UpdatedAt = time.Now()
	return nil
}
