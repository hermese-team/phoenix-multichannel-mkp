// Package productsync pushes product price/stock from our catalog up to Lazada.
// It is the outbound counterpart to ordersync: our DB is the source of truth and
// Lazada is updated via /product/price_quantity/update.
package productsync

import (
	"context"
	"encoding/xml"
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/okdev/marketplace-sync/internal/infrastructure/postgres"
	"github.com/okdev/marketplace-sync/pkg/logger"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/client"
	"github.com/okdev/marketplace-sync/sellchannel/lazada/tokenstore"
)

type Syncer struct {
	cfg    client.Config
	tokens *tokenstore.Store
	repo   *postgres.LazadaProductRepository
	http   *resty.Client // shared across per-call clients to reuse connections
}

func New(cfg client.Config, tokens *tokenstore.Store, repo *postgres.LazadaProductRepository) *Syncer {
	return &Syncer{cfg: cfg, tokens: tokens, repo: repo, http: client.NewHTTPClient()}
}

// SyncPending pushes up to limit pending products' price/stock to Lazada in one
// request and records each SKU's outcome. The push is all-or-nothing here: on
// failure every SKU in the batch is marked failed with the same reason (per-SKU
// result mapping is a follow-up).
func (s *Syncer) SyncPending(ctx context.Context, limit int) error {
	products, err := s.repo.ListPending(ctx, limit)
	if err != nil {
		return err
	}
	if len(products) == 0 {
		logger.Info("no pending products to sync")
		return nil
	}

	payload, err := buildPriceStockXML(products)
	if err != nil {
		return err
	}

	if err := s.push(ctx, payload); err != nil {
		logger.Error("push price/stock failed", "count", len(products), "error", err)
		s.markAll(ctx, products, postgres.ProductSyncFailed, err.Error())
		return err
	}
	s.markAll(ctx, products, postgres.ProductSyncSynced, "")
	logger.Info("price/stock pushed", "count", len(products))
	return nil
}

// push sends the payload, refreshing the token once on an auth error and retrying.
func (s *Syncer) push(ctx context.Context, payload string) error {
	c, err := s.clientWithStoredToken(ctx)
	if err != nil {
		return err
	}
	if _, err := c.UpdateStockPrice(ctx, payload); err != nil {
		if !client.IsAuthError(err) {
			return err
		}
		logger.Warn("price/stock push auth error, refreshing token", "error", err)
		if rerr := s.refreshTokens(ctx); rerr != nil {
			return err
		}
		if c, err = s.clientWithStoredToken(ctx); err != nil {
			return err
		}
		if _, err := c.UpdateStockPrice(ctx, payload); err != nil {
			return err
		}
	}
	return nil
}

func (s *Syncer) markAll(ctx context.Context, products []postgres.LazadaProduct, status, errMsg string) {
	for _, p := range products {
		if err := s.repo.MarkResult(ctx, p.SellerSku, status, errMsg); err != nil {
			logger.Error("mark product result", "seller_sku", p.SellerSku, "error", err)
		}
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

// --- price/stock XML payload (<Request><Product><Skus><Sku>…) ---

type skuXML struct {
	SellerSku string  `xml:"SellerSku"`
	Price     float64 `xml:"Price"`
	Quantity  int     `xml:"Quantity"`
}

type requestXML struct {
	XMLName xml.Name `xml:"Request"`
	Skus    []skuXML `xml:"Product>Skus>Sku"`
}

func buildPriceStockXML(products []postgres.LazadaProduct) (string, error) {
	req := requestXML{Skus: make([]skuXML, 0, len(products))}
	for _, p := range products {
		req.Skus = append(req.Skus, skuXML{
			SellerSku: p.SellerSku,
			Price:     p.Price,
			Quantity:  p.Quantity,
		})
	}
	b, err := xml.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("marshal price/stock xml: %w", err)
	}
	return string(b), nil
}
