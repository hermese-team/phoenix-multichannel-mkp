// Package productoffsalesync takes products/SKUs off sale on Lazada:
// deactivate a product, remove a product's seller skus, or remove specific SKUs.
// Outbound loop; the action per row is chosen by its Action field.
package productoffsalesync

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"

	"github.com/okdev/marketplace-sync/internal/infrastructure/postgres"
	"github.com/okdev/marketplace-sync/pkg/logger"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/client"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/tokenstore"
)

type Syncer struct {
	cfg    client.Config
	tokens *tokenstore.Store
	repo   *postgres.LazadaProductOffsaleRepository
	http   *resty.Client // shared across per-call clients to reuse connections
}

func New(cfg client.Config, tokens *tokenstore.Store, repo *postgres.LazadaProductOffsaleRepository) *Syncer {
	return &Syncer{cfg: cfg, tokens: tokens, repo: repo, http: client.NewHTTPClient()}
}

// ProcessPending runs up to limit pending off-sale actions and records outcomes.
func (s *Syncer) ProcessPending(ctx context.Context, limit int) error {
	rows, err := s.repo.ListPending(ctx, limit)
	if err != nil {
		return err
	}
	if len(rows) == 0 {
		logger.Info("no pending offsale actions")
		return nil
	}

	done := 0
	for _, o := range rows {
		if err := s.process(ctx, o); err != nil {
			logger.Error("offsale", "id", o.ID, "action", o.Action, "error", err)
			if merr := s.repo.MarkFailed(ctx, o.ID, err.Error()); merr != nil {
				logger.Error("mark offsale failed", "id", o.ID, "error", merr)
			}
			continue
		}
		if merr := s.repo.MarkDone(ctx, o.ID); merr != nil {
			logger.Error("mark offsale done", "id", o.ID, "error", merr)
			continue
		}
		done++
		logger.Info("offsale done", "id", o.ID, "action", o.Action)
	}
	logger.Info("offsale sync done", "processed", len(rows), "done", done)
	return nil
}

// process runs the row's action, refreshing the token once on an auth error.
func (s *Syncer) process(ctx context.Context, o postgres.LazadaOffsale) error {
	c, err := s.clientWithStoredToken(ctx)
	if err != nil {
		return err
	}
	err = s.doAction(ctx, c, o)
	if err != nil && client.IsAuthError(err) {
		logger.Warn("offsale auth error, refreshing token", "id", o.ID, "error", err)
		if rerr := s.refreshTokens(ctx); rerr != nil {
			return err
		}
		if c, err = s.clientWithStoredToken(ctx); err != nil {
			return err
		}
		err = s.doAction(ctx, c, o)
	}
	return err
}

// doAction dispatches to the right client call for the row's Action.
func (s *Syncer) doAction(ctx context.Context, c *client.Client, o postgres.LazadaOffsale) error {
	switch o.Action {
	case postgres.ActionDeactivate:
		if o.ItemID == 0 {
			return fmt.Errorf("deactivate: missing item_id")
		}
		return c.DeactivateProduct(ctx, o.ItemID)

	case postgres.ActionRemoveProduct:
		skus := splitCSV(o.SellerSkus)
		if len(skus) == 0 {
			return fmt.Errorf("remove_product: no seller_skus")
		}
		return c.RemoveProduct(ctx, skus)

	case postgres.ActionRemoveSku:
		skus := splitCSV(o.SellerSkus)
		if o.ItemID == 0 || len(skus) == 0 {
			return fmt.Errorf("remove_sku: need item_id and seller_skus")
		}
		return c.RemoveSku(ctx, o.ItemID, o.VariationName, skus)

	default:
		return fmt.Errorf("unknown offsale action %q", o.Action)
	}
}

func splitCSV(csv string) []string {
	parts := strings.Split(csv, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if p = strings.TrimSpace(p); p != "" {
			out = append(out, p)
		}
	}
	return out
}

// clientWithStoredToken / refreshTokens mirror ordersync: tokens live in Redis.
func (s *Syncer) clientWithStoredToken(ctx context.Context) (*client.Client, error) {
	access, refresh, err := s.tokens.Get(ctx)
	if err != nil {
		return nil, err
	}
	cfg := s.cfg
	cfg.AccessToken = access
	cfg.RefreshToken = refresh
	return client.NewWithHTTP(cfg, s.http), nil
}

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
