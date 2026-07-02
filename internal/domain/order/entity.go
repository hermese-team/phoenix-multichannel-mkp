package order

import "time"

type Status string

const (
	StatusPending    Status = "pending"
	StatusProcessing Status = "processing"
	StatusCompleted  Status = "completed"
	StatusCancelled  Status = "cancelled"
)

type Order struct {
	ID              int64
	ShopID          int64
	SellChannelID   string
	SellChannelType string
	Status          Status
	TotalAmount     float64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
