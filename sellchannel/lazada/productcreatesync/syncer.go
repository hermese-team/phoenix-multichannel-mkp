// Package productcreatesync creates products on Lazada from our catalog. It is
// an outbound loop like productsync, but calls /product/create (one product per
// call) and records the returned Lazada item id.
package productcreatesync

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
	repo   *postgres.LazadaProductCreateRepository
	http   *resty.Client // shared across per-call clients to reuse connections
}

func New(cfg client.Config, tokens *tokenstore.Store, repo *postgres.LazadaProductCreateRepository) *Syncer {
	return &Syncer{cfg: cfg, tokens: tokens, repo: repo, http: client.NewHTTPClient()}
}

// CreatePending creates up to limit pending products on Lazada, one call each,
// recording each product's outcome (item id or error).
func (s *Syncer) CreatePending(ctx context.Context, limit int) error {
	products, err := s.repo.ListPending(ctx, limit)
	if err != nil {
		return err
	}
	if len(products) == 0 {
		logger.Info("no pending products to create")
		return nil
	}

	created := 0
	for _, p := range products {
		resp, err := s.createOne(ctx, p)
		if err != nil {
			logger.Error("create product", "seller_sku", p.SellerSku, "error", err)
			if merr := s.repo.MarkFailed(ctx, p.SellerSku, err.Error()); merr != nil {
				logger.Error("mark product failed", "seller_sku", p.SellerSku, "error", merr)
			}
			continue
		}
		if merr := s.repo.MarkCreated(ctx, p.SellerSku, resp.ItemID); merr != nil {
			logger.Error("mark product created", "seller_sku", p.SellerSku, "error", merr)
			continue
		}
		created++
		logger.Info("product created", "seller_sku", p.SellerSku, "item_id", resp.ItemID)
	}
	logger.Info("product create done", "processed", len(products), "created", created)
	return nil
}

// createOne builds the payload and creates one product, refreshing the token
// once on an auth error and retrying.
func (s *Syncer) createOne(ctx context.Context, p postgres.LazadaCreateProduct) (*client.CreateProductResponse, error) {
	payload, err := buildCreateProductXML(p)
	if err != nil {
		return nil, err
	}
	c, err := s.clientWithStoredToken(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := c.CreateProduct(ctx, payload)
	if err != nil {
		if !client.IsAuthError(err) {
			return nil, err
		}
		logger.Warn("create product auth error, refreshing token", "seller_sku", p.SellerSku, "error", err)
		if rerr := s.refreshTokens(ctx); rerr != nil {
			return nil, err
		}
		if c, err = s.clientWithStoredToken(ctx); err != nil {
			return nil, err
		}
		if resp, err = c.CreateProduct(ctx, payload); err != nil {
			return nil, err
		}
	}
	return resp, nil
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

// --- CreateProduct XML payload ------------------------------------------------
//
// <Attributes> is dynamic (keys vary by category), which encoding/xml can't
// marshal from a map, so the document is built by hand. Text values are escaped;
// attribute keys are assumed to be valid element names (they come from
// GetCategoryAttributes).

func buildCreateProductXML(p postgres.LazadaCreateProduct) (string, error) {
	var b strings.Builder
	b.WriteString("<Request><Product>")
	fmt.Fprintf(&b, "<PrimaryCategory>%d</PrimaryCategory>", p.CategoryID)

	b.WriteString("<Attributes>")
	writeEl(&b, "name", p.Name)
	for k, v := range p.Attributes {
		writeEl(&b, k, v)
	}
	b.WriteString("</Attributes>")

	b.WriteString("<Skus><Sku>")
	writeEl(&b, "SellerSku", p.SellerSku)
	fmt.Fprintf(&b, "<quantity>%d</quantity>", p.Quantity)
	writeEl(&b, "price", strconv.FormatFloat(p.Price, 'f', -1, 64))
	writeEl(&b, "package_weight", strconv.FormatFloat(p.PackageWeight, 'f', -1, 64))
	writeEl(&b, "package_length", strconv.FormatFloat(p.PackageLength, 'f', -1, 64))
	writeEl(&b, "package_width", strconv.FormatFloat(p.PackageWidth, 'f', -1, 64))
	writeEl(&b, "package_height", strconv.FormatFloat(p.PackageHeight, 'f', -1, 64))
	if p.ImageURL != "" {
		b.WriteString("<Images>")
		writeEl(&b, "Image", p.ImageURL)
		b.WriteString("</Images>")
	}
	b.WriteString("</Sku></Skus>")

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
