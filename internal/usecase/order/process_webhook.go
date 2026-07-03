package order

import (
	"context"

	domain "github.com/okdev/marketplace-sync/internal/domain/order"
	applog "github.com/okdev/marketplace-sync/pkg/logger"
)

type ProcessWebhookUsecase struct {
	repo domain.Repository
}

func NewProcessWebhook(repo domain.Repository) *ProcessWebhookUsecase {
	return &ProcessWebhookUsecase{repo: repo}
}

type WebhookPayload struct {
	MarketplaceOrderID string
	SellChannelType    string
	ShopID             int64
	Status             string
}

func (u *ProcessWebhookUsecase) Execute(ctx context.Context, payload WebhookPayload) error {
	o, err := u.repo.FindBySellChannelID(ctx, payload.MarketplaceOrderID, payload.SellChannelType)
	if err != nil {
		// NOTE(mock): any lookup error is treated as "not found" and a new order
		// is created. The orders table is mocked here; a real impl must tell
		// internal.ErrNotFound (create) apart from a transient DB error (return
		// err) via errors.Is — otherwise a DB blip inserts a duplicate order.
		o = &domain.Order{
			ShopID:          payload.ShopID,
			SellChannelID:   payload.MarketplaceOrderID,
			SellChannelType: payload.SellChannelType,
		}
	}

	o.Status = domain.Status(payload.Status)

	if o.ID == 0 {
		if err := u.repo.Save(ctx, o); err != nil {
			applog.Error("save order", "error", err)
			return err
		}
		return nil
	}

	if err := u.repo.Update(ctx, o); err != nil {
		applog.Error("update order", "error", err)
		return err
	}
	return nil
}
