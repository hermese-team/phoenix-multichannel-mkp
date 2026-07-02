package product

import (
	"context"
	"fmt"

	domain "github.com/okdev/marketplace-sync/internal/domain/product"
)

type PublishUsecase struct {
	repo        domain.Repository
	sellchannel domain.SellChannel
}

func NewPublish(repo domain.Repository, sellchannel domain.SellChannel) *PublishUsecase {
	return &PublishUsecase{repo: repo, sellchannel: sellchannel}
}

func (u *PublishUsecase) Execute(ctx context.Context, shopID, productID int64) error {
	p, err := u.repo.FindByID(ctx, productID)
	if err != nil {
		return fmt.Errorf("find product: %w", err)
	}
	if p.Status != domain.StatusActive {
		return domain.ErrNotActive
	}

	itemID, err := u.sellchannel.PublishProduct(ctx, shopID, p)
	if err != nil {
		return fmt.Errorf("publish product: %w", err)
	}

	p.SellChannelID = fmt.Sprintf("%d", itemID)
	if err := u.repo.Update(ctx, p); err != nil {
		return fmt.Errorf("update product marketplace id: %w", err)
	}
	return nil
}
