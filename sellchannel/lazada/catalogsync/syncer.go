// Package catalogsync pulls the Lazada product catalog down into a local
// snapshot for reconciliation (drift detection). Inbound poll-and-store, like
// ordersync — not an outbound push loop.
package catalogsync

import (
	"context"

	"github.com/go-resty/resty/v2"

	"github.com/okdev/marketplace-sync/internal/infrastructure/postgres"
	"github.com/okdev/marketplace-sync/pkg/logger"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/client"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/tokenstore"
)

// pageLimit is GetProducts' max products per page.
const pageLimit = 50

// maxOffset is Lazada's cap on the (deprecated) offset parameter. Catalogs
// larger than this need date-scrolling (update_after) for a full reconcile —
// a follow-up; for now we stop cleanly at the cap instead of hitting an API error.
const maxOffset = 10000

type Syncer struct {
	cfg    client.Config
	tokens *tokenstore.Store
	repo   *postgres.LazadaProductCatalogRepository
	http   *resty.Client
}

func New(cfg client.Config, tokens *tokenstore.Store, repo *postgres.LazadaProductCatalogRepository) *Syncer {
	return &Syncer{cfg: cfg, tokens: tokens, repo: repo, http: client.NewHTTPClient()}
}

// Reconcile pages through GetProducts (filtered by status) and upserts every sku
// into the snapshot table.
func (s *Syncer) Reconcile(ctx context.Context, filter string) error {
	total, synced := 0, 0
	for offset := 0; ; {
		t, products, err := s.getPage(ctx, filter, offset)
		if err != nil {
			return err
		}
		if offset == 0 {
			total = t
			logger.Info("catalog reconcile started", "filter", filter, "total_products", total)
		}
		if len(products) == 0 {
			break
		}
		for _, p := range products {
			for _, sku := range p.Skus {
				if err := s.repo.UpsertSku(ctx, postgres.LazadaCatalogSku{
					ItemID:    p.ItemID,
					SkuID:     sku.SkuID,
					SellerSku: sku.SellerSku,
					ShopSku:   sku.ShopSku,
					Name:      p.Attributes.Name,
					Status:    sku.Status,
					Price:     sku.Price,
					Quantity:  sku.Quantity,
				}); err != nil {
					logger.Error("upsert catalog sku", "item_id", p.ItemID, "sku_id", sku.SkuID, "error", err)
					continue
				}
				synced++
			}
		}
		// Advance by products actually returned; stop once caught up to total.
		offset += len(products)
		if total > 0 && offset >= total {
			break
		}
		if offset >= maxOffset {
			logger.Warn("catalog reconcile stopped at offset cap; use date-scrolling for larger catalogs", "offset", offset)
			break
		}
	}
	logger.Info("catalog reconcile done", "total_products", total, "skus_synced", synced)
	return nil
}

// getPage fetches one page, refreshing the token once on an auth error.
func (s *Syncer) getPage(ctx context.Context, filter string, offset int) (int, []client.CatalogProduct, error) {
	c, err := s.clientWithStoredToken(ctx)
	if err != nil {
		return 0, nil, err
	}
	total, products, err := c.GetProducts(ctx, filter, offset, pageLimit)
	if err != nil && client.IsAuthError(err) {
		logger.Warn("catalog page auth error, refreshing token", "error", err)
		if rerr := s.refreshTokens(ctx); rerr != nil {
			return 0, nil, err
		}
		if c, err = s.clientWithStoredToken(ctx); err != nil {
			return 0, nil, err
		}
		total, products, err = c.GetProducts(ctx, filter, offset, pageLimit)
	}
	return total, products, err
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
