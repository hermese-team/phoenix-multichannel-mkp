package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

// LazadaCatalogSku is a snapshot row of a Lazada product sku pulled down for
// reconciliation (drift detection against our own catalog).
type LazadaCatalogSku struct {
	ItemID    int64
	SkuID     int64
	SellerSku string
	ShopSku   string
	Name      string
	Status    string
	Price     float64
	Quantity  int
}

// LazadaProductCatalogRepository stores the Lazada catalog snapshot.
type LazadaProductCatalogRepository struct {
	db *sql.DB
}

func NewLazadaProductCatalogRepository(db *sql.DB) *LazadaProductCatalogRepository {
	return &LazadaProductCatalogRepository{db: db}
}

// EnsureSchema creates the snapshot table if it does not exist (test convenience;
// a real deployment would use a migration).
func (r *LazadaProductCatalogRepository) EnsureSchema(ctx context.Context) error {
	const q = `
	CREATE TABLE IF NOT EXISTS lazada_product_catalog (
		item_id    BIGINT           NOT NULL,
		sku_id     BIGINT           NOT NULL,
		seller_sku TEXT             NOT NULL DEFAULT '',
		shop_sku   TEXT             NOT NULL DEFAULT '',
		name       TEXT             NOT NULL DEFAULT '',
		status     TEXT             NOT NULL DEFAULT '',
		price      DOUBLE PRECISION NOT NULL DEFAULT 0,
		quantity   INTEGER          NOT NULL DEFAULT 0,
		synced_at  TIMESTAMPTZ      NOT NULL DEFAULT now(),
		PRIMARY KEY (item_id, sku_id)
	)`
	if _, err := r.db.ExecContext(ctx, q); err != nil {
		return fmt.Errorf("ensure lazada catalog schema: %w", err)
	}
	return nil
}

// UpsertSku inserts or refreshes a snapshot row keyed by (item_id, sku_id).
func (r *LazadaProductCatalogRepository) UpsertSku(ctx context.Context, s LazadaCatalogSku) error {
	const q = `
	INSERT INTO lazada_product_catalog
		(item_id, sku_id, seller_sku, shop_sku, name, status, price, quantity, synced_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, now())
	ON CONFLICT (item_id, sku_id) DO UPDATE SET
		seller_sku = EXCLUDED.seller_sku,
		shop_sku   = EXCLUDED.shop_sku,
		name       = EXCLUDED.name,
		status     = EXCLUDED.status,
		price      = EXCLUDED.price,
		quantity   = EXCLUDED.quantity,
		synced_at  = now()`
	if _, err := r.db.ExecContext(ctx, q,
		s.ItemID, s.SkuID, s.SellerSku, s.ShopSku, s.Name, s.Status, s.Price, s.Quantity,
	); err != nil {
		return fmt.Errorf("upsert catalog sku: %w", err)
	}
	return nil
}
