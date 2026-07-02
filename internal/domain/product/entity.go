package product

import "time"

type Status string

const (
	StatusDraft    Status = "draft"
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
	StatusDeleted  Status = "deleted"
)

type Product struct {
	ID              int64
	ShopID          int64
	Name            string
	SKU             string
	Description     string
	Status          Status
	SellChannelID   string
	SellChannelType string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
