package entity

import (
	"time"

	domainerrors "github.com/ascend/phoenix-multichannel-mkp/pkg/errors"
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusConfirmed OrderStatus = "confirmed"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCancelled OrderStatus = "cancelled"
)

type OrderItem struct {
	ProductID string
	Name      string
	Price     float64
	Quantity  int
}

func (i *OrderItem) Subtotal() float64 {
	return i.Price * float64(i.Quantity)
}

type Order struct {
	ID         string
	BuyerID    string
	SellerID   string
	Channel    Channel
	Items      []OrderItem
	Status     OrderStatus
	TotalPrice float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewOrder(buyerID, sellerID string, channel Channel, items []OrderItem) (*Order, error) {
	if len(items) == 0 {
		return nil, domainerrors.New("EMPTY_ORDER", "order must have at least one item", domainerrors.ErrInvalidInput)
	}
	var total float64
	for _, item := range items {
		if item.Quantity <= 0 {
			return nil, domainerrors.New("INVALID_QUANTITY", "item quantity must be positive", domainerrors.ErrInvalidInput)
		}
		total += item.Subtotal()
	}
	now := time.Now()
	return &Order{
		BuyerID:    buyerID,
		SellerID:   sellerID,
		Channel:    channel,
		Items:      items,
		Status:     OrderStatusPending,
		TotalPrice: total,
		CreatedAt:  now,
		UpdatedAt:  now,
	}, nil
}

func (o *Order) Confirm() error {
	if o.Status != OrderStatusPending {
		return domainerrors.New("INVALID_TRANSITION", "only pending orders can be confirmed", domainerrors.ErrInvalidInput)
	}
	o.Status = OrderStatusConfirmed
	o.UpdatedAt = time.Now()
	return nil
}

func (o *Order) Cancel() error {
	if o.Status == OrderStatusDelivered || o.Status == OrderStatusShipped {
		return domainerrors.New("INVALID_TRANSITION", "cannot cancel a shipped or delivered order", domainerrors.ErrInvalidInput)
	}
	o.Status = OrderStatusCancelled
	o.UpdatedAt = time.Now()
	return nil
}
