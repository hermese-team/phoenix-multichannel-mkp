// Package productupdatesync pushes listing-content edits (name + attributes) to
// existing Lazada products via /product/update. Outbound loop like
// productcreatesync; price/stock edits go through productsync instead.
package productupdatesync

import (
	"context"
	"encoding/xml"
	"fmt"
	"strconv"
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
	repo   *postgres.LazadaProductUpdateRepository
	http   *resty.Client // shared across per-call clients to reuse connections
}

func New(cfg client.Config, tokens *tokenstore.Store, repo *postgres.LazadaProductUpdateRepository) *Syncer {
	return &Syncer{cfg: cfg, tokens: tokens, repo: repo, http: client.NewHTTPClient()}
}

// UpdatePending pushes up to limit pending edits to Lazada, one call each,
// recording each product's outcome.
func (s *Syncer) UpdatePending(ctx context.Context, limit int) error {
	products, err := s.repo.ListPending(ctx, limit)
	if err != nil {
		return err
	}
	if len(products) == 0 {
		logger.Info("no pending product updates")
		return nil
	}

	updated := 0
	for _, p := range products {
		if err := s.updateOne(ctx, p); err != nil {
			logger.Error("update product", "seller_sku", p.SellerSku, "error", err)
			if merr := s.repo.MarkFailed(ctx, p.SellerSku, err.Error()); merr != nil {
				logger.Error("mark product update failed", "seller_sku", p.SellerSku, "error", merr)
			}
			continue
		}
		if merr := s.repo.MarkUpdated(ctx, p.SellerSku); merr != nil {
			logger.Error("mark product updated", "seller_sku", p.SellerSku, "error", merr)
			continue
		}
		updated++
		logger.Info("product updated", "seller_sku", p.SellerSku, "item_id", p.ItemID)
	}
	logger.Info("product update done", "processed", len(products), "updated", updated)
	return nil
}

// updateOne builds the payload and updates one product, refreshing the token
// once on an auth error and retrying.
func (s *Syncer) updateOne(ctx context.Context, p postgres.LazadaUpdateProduct) error {
	payload, err := buildUpdateProductXML(p)
	if err != nil {
		return err
	}
	c, err := s.clientWithStoredToken(ctx)
	if err != nil {
		return err
	}
	_, err = c.UpdateProduct(ctx, payload)
	if err != nil {
		if !client.IsAuthError(err) {
			return err
		}
		logger.Warn("update product auth error, refreshing token", "seller_sku", p.SellerSku, "error", err)
		if rerr := s.refreshTokens(ctx); rerr != nil {
			return err
		}
		if c, err = s.clientWithStoredToken(ctx); err != nil {
			return err
		}
		if _, err = c.UpdateProduct(ctx, payload); err != nil {
			return err
		}
	}
	return nil
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

// --- UpdateProduct XML payload ------------------------------------------------
//
// Edits product-level attributes only (identified by ItemId). <Attributes> keys
// are dynamic, so the document is built by hand with escaped values.
//
// NOTE(mock): sku-level edits (price/quantity) are handled by productsync, so no
// <Skus> block is sent here. If Lazada requires a Sku identifier (SkuId) for
// certain attribute edits, extend this payload.

func buildUpdateProductXML(p postgres.LazadaUpdateProduct) (string, error) {
	if p.ItemID == 0 {
		return "", fmt.Errorf("update product %s: missing item_id", p.SellerSku)
	}
	if p.Name == "" && len(p.Attributes) == 0 {
		return "", fmt.Errorf("update product %s: nothing to update", p.SellerSku)
	}
	var b strings.Builder
	b.WriteString("<Request><Product>")
	b.WriteString("<ItemId>")
	b.WriteString(strconv.FormatInt(p.ItemID, 10))
	b.WriteString("</ItemId>")

	b.WriteString("<Attributes>")
	if p.Name != "" {
		writeEl(&b, "name", p.Name)
	}
	for k, v := range p.Attributes {
		writeEl(&b, k, v)
	}
	b.WriteString("</Attributes>")

	b.WriteString("</Product></Request>")
	return b.String(), nil
}

// writeEl writes <name>escaped-value</name>.
func writeEl(b *strings.Builder, name, value string) {
	b.WriteByte('<')
	b.WriteString(name)
	b.WriteByte('>')
	_ = xml.EscapeText(b, []byte(value))
	b.WriteString("</")
	b.WriteString(name)
	b.WriteByte('>')
}
