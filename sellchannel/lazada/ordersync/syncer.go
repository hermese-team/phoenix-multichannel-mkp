// Package ordersync fetches Lazada orders (detail + items) and persists them.
// It is shared by the scheduler (poll) and the webhook consumer (push).
package ordersync

import (
	"context"
	"strings"
	"time"

	"github.com/okdev/marketplace-sync/internal/infrastructure/postgres"
	"github.com/okdev/marketplace-sync/pkg/logger"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/client"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/tokenstore"
)

type Syncer struct {
	cfg    client.Config
	tokens *tokenstore.Store
	repo   *postgres.LazadaSyncRepository
}

func New(cfg client.Config, tokens *tokenstore.Store, repo *postgres.LazadaSyncRepository) *Syncer {
	return &Syncer{cfg: cfg, tokens: tokens, repo: repo}
}

// SyncOrder fetches an order's detail + items and upserts it. On an API error it
// refreshes the token once and retries.
func (s *Syncer) SyncOrder(ctx context.Context, orderID string) error {
	detail, items, err := s.fetch(ctx, orderID)
	if err != nil {
		logger.Warn("fetch order failed, refreshing token", "order", orderID, "error", err)
		if rerr := s.refreshTokens(ctx); rerr != nil {
			return err
		}
		if detail, items, err = s.fetch(ctx, orderID); err != nil {
			return err
		}
	}
	return s.repo.SaveOrder(ctx, postgres.LazadaSyncedOrder{
		OrderNumber:   orderID,
		OrderID:       detail.OrderID,
		Statuses:      strings.Join(detail.Statuses, ","),
		PaymentMethod: detail.PaymentMethod,
		ItemCount:     len(items),
	})
}

// ListOrders returns order ids updated within [after, before]. On an API error
// it refreshes the token once and retries.
func (s *Syncer) ListOrders(ctx context.Context, after, before time.Time, offset int) (int, []string, error) {
	c, err := s.clientWithStoredToken(ctx)
	if err != nil {
		return 0, nil, err
	}
	total, ids, err := c.GetOrderList(ctx, after, before, offset)
	if err != nil {
		logger.Warn("order list failed, refreshing token", "error", err)
		if rerr := s.refreshTokens(ctx); rerr != nil {
			return 0, nil, err
		}
		if c, err = s.clientWithStoredToken(ctx); err != nil {
			return 0, nil, err
		}
		if total, ids, err = c.GetOrderList(ctx, after, before, offset); err != nil {
			return 0, nil, err
		}
	}
	return total, ids, nil
}

func (s *Syncer) fetch(ctx context.Context, orderID string) (*client.OrderDetail, []client.OrderItem, error) {
	c, err := s.clientWithStoredToken(ctx)
	if err != nil {
		return nil, nil, err
	}
	detail, err := c.GetOrderDetail(ctx, orderID)
	if err != nil {
		return nil, nil, err
	}
	items, err := c.GetOrderItems(ctx, orderID)
	if err != nil {
		return nil, nil, err
	}
	return detail, items, nil
}

// clientWithStoredToken builds a client using the tokens currently in Redis.
func (s *Syncer) clientWithStoredToken(ctx context.Context) (*client.Client, error) {
	access, refresh, err := s.tokens.Get(ctx)
	if err != nil {
		return nil, err
	}
	cfg := s.cfg
	cfg.AccessToken = access
	cfg.RefreshToken = refresh
	return client.New(cfg), nil
}

// refreshTokens uses the stored refresh token to obtain a new pair and writes it
// back to Redis.
func (s *Syncer) refreshTokens(ctx context.Context) error {
	c, err := s.clientWithStoredToken(ctx)
	if err != nil {
		return err
	}
	res, err := c.RefreshToken(ctx)
	if err != nil {
		return err
	}
	logger.Info("lazada token refreshed", "expires_in", res.ExpiresIn)
	return s.tokens.Set(ctx, res.AccessToken, res.RefreshToken, 0)
}
