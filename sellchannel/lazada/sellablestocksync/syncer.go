// Package sellablestocksync pushes warehouse sellable-quantity changes to Lazada
// via /product/stock/sellable/{adjust,update}. Distinct from productsync, which
// updates price+stock via /product/price_quantity/update.
package sellablestocksync

import (
	"context"
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/go-resty/resty/v2"

	"github.com/okdev/marketplace-sync/internal/infrastructure/postgres"
	"github.com/okdev/marketplace-sync/pkg/logger"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/client"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/tokenstore"
)

type Syncer struct {
	cfg    client.Config
	tokens *tokenstore.Store
	repo   *postgres.LazadaSellableStockRepository
	http   *resty.Client // shared across per-call clients to reuse connections
}

func New(cfg client.Config, tokens *tokenstore.Store, repo *postgres.LazadaSellableStockRepository) *Syncer {
	return &Syncer{cfg: cfg, tokens: tokens, repo: repo, http: client.NewHTTPClient()}
}

// SyncPending pushes up to limit pending stock changes, one call each, and
// records each outcome.
func (s *Syncer) SyncPending(ctx context.Context, limit int) error {
	rows, err := s.repo.ListPending(ctx, limit)
	if err != nil {
		return err
	}
	if len(rows) == 0 {
		logger.Info("no pending sellable-stock changes")
		return nil
	}

	done := 0
	for _, o := range rows {
		if err := s.process(ctx, o); err != nil {
			logger.Error("sellable-stock", "id", o.ID, "action", o.Action, "error", err)
			if merr := s.repo.MarkFailed(ctx, o.ID, err.Error()); merr != nil {
				logger.Error("mark sellable-stock failed", "id", o.ID, "error", merr)
			}
			continue
		}
		if merr := s.repo.MarkDone(ctx, o.ID); merr != nil {
			logger.Error("mark sellable-stock done", "id", o.ID, "error", merr)
			continue
		}
		done++
		logger.Info("sellable-stock done", "id", o.ID, "action", o.Action, "sku_id", o.SkuID)
	}
	logger.Info("sellable-stock sync done", "processed", len(rows), "done", done)
	return nil
}

// process runs the row's stock change, refreshing the token once on an auth error.
func (s *Syncer) process(ctx context.Context, o postgres.LazadaSellableStock) error {
	payload, err := buildStockXML(o)
	if err != nil {
		return err
	}
	c, err := s.clientWithStoredToken(ctx)
	if err != nil {
		return err
	}
	err = s.doStock(ctx, c, o.Action, payload)
	if err != nil && client.IsAuthError(err) {
		logger.Warn("sellable-stock auth error, refreshing token", "id", o.ID, "error", err)
		if rerr := s.refreshTokens(ctx); rerr != nil {
			return err
		}
		if c, err = s.clientWithStoredToken(ctx); err != nil {
			return err
		}
		err = s.doStock(ctx, c, o.Action, payload)
	}
	return err
}

func (s *Syncer) doStock(ctx context.Context, c *client.Client, action, payload string) error {
	switch action {
	case postgres.StockAdjust:
		return c.AdjustSellableQuantity(ctx, payload)
	case postgres.StockUpdate:
		return c.UpdateSellableQuantity(ctx, payload)
	default:
		return fmt.Errorf("unknown stock action %q", action)
	}
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

// --- sellable-stock XML payload (multi-warehouse; mirrors the reference) ------
//
// <Request><Product><Skus><Sku><ItemId>…</ItemId><SkuId>…</SkuId>
//   <MultiWarehouseInventories><MultiWarehouseInventory>
//     <WarehouseCode>…</WarehouseCode><SellableQuantity>…</SellableQuantity>
//   </MultiWarehouseInventory></MultiWarehouseInventories>
// </Sku></Skus></Product></Request>

type stockInv struct {
	WarehouseCode    string `xml:"WarehouseCode"`
	SellableQuantity int    `xml:"SellableQuantity"`
}

type stockSku struct {
	ItemID string   `xml:"ItemId"`
	SkuID  string   `xml:"SkuId"`
	Inv    stockInv `xml:"MultiWarehouseInventories>MultiWarehouseInventory"`
}

type stockRequest struct {
	XMLName xml.Name   `xml:"Request"`
	Skus    []stockSku `xml:"Product>Skus>Sku"`
}

func buildStockXML(o postgres.LazadaSellableStock) (string, error) {
	if o.ItemID == 0 || o.SkuID == 0 {
		return "", fmt.Errorf("sellable-stock: need item_id and sku_id")
	}
	// Multi-warehouse inventory requires a real warehouse code; an empty one is
	// rejected, so fail fast with a clear message instead.
	if o.WarehouseCode == "" {
		return "", fmt.Errorf("sellable-stock: missing warehouse_code")
	}
	req := stockRequest{Skus: []stockSku{{
		ItemID: strconv.FormatInt(o.ItemID, 10),
		SkuID:  strconv.FormatInt(o.SkuID, 10),
		Inv:    stockInv{WarehouseCode: o.WarehouseCode, SellableQuantity: o.Quantity},
	}}}
	b, err := xml.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("marshal sellable-stock xml: %w", err)
	}
	return string(b), nil
}
