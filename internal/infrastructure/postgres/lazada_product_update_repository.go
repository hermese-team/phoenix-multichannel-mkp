package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
)

// Product-update sync states.
const (
	ProductUpdatePending = "pending"
	ProductUpdateUpdated = "updated"
	ProductUpdateFailed  = "failed"
)

// LazadaUpdateProduct is a mock row describing a listing-content edit to push to
// Lazada. ItemID identifies the product; Name + Attributes are the fields to
// change. Price/stock edits go through the price/stock loop, not here.
type LazadaUpdateProduct struct {
	SellerSku  string
	ItemID     int64
	Name       string
	Attributes map[string]string
}

// LazadaProductUpdateRepository holds product edits pending push to Lazada.
type LazadaProductUpdateRepository struct {
	db *sql.DB
}

func NewLazadaProductUpdateRepository(db *sql.DB) *LazadaProductUpdateRepository {
	return &LazadaProductUpdateRepository{db: db}
}

// EnsureSchema creates the mock table if it does not exist (test convenience;
// a real deployment would use a migration).
func (r *LazadaProductUpdateRepository) EnsureSchema(ctx context.Context) error {
	const q = `
	CREATE TABLE IF NOT EXISTS lazada_product_update (
		seller_sku  TEXT        PRIMARY KEY,
		item_id     BIGINT      NOT NULL,
		name        TEXT        NOT NULL DEFAULT '',
		attributes  JSONB       NOT NULL DEFAULT '{}',
		sync_status TEXT        NOT NULL DEFAULT 'pending',
		last_error  TEXT        NOT NULL DEFAULT '',
		updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
	)`
	if _, err := r.db.ExecContext(ctx, q); err != nil {
		return fmt.Errorf("ensure lazada product-update schema: %w", err)
	}
	return nil
}

// ListPending returns up to limit edits awaiting push, oldest first.
func (r *LazadaProductUpdateRepository) ListPending(ctx context.Context, limit int) ([]LazadaUpdateProduct, error) {
	const q = `SELECT seller_sku, item_id, name, attributes
	           FROM lazada_product_update
	           WHERE sync_status = $1 ORDER BY updated_at LIMIT $2`
	rows, err := r.db.QueryContext(ctx, q, ProductUpdatePending, limit)
	if err != nil {
		return nil, fmt.Errorf("list pending product-update: %w", err)
	}
	defer rows.Close()

	var out []LazadaUpdateProduct
	for rows.Next() {
		var (
			p        LazadaUpdateProduct
			attrsRaw []byte
		)
		if err := rows.Scan(&p.SellerSku, &p.ItemID, &p.Name, &attrsRaw); err != nil {
			return nil, fmt.Errorf("scan product-update: %w", err)
		}
		if len(attrsRaw) > 0 {
			if err := json.Unmarshal(attrsRaw, &p.Attributes); err != nil {
				return nil, fmt.Errorf("decode attributes for %s: %w", p.SellerSku, err)
			}
		}
		out = append(out, p)
	}
	return out, rows.Err()
}

// MarkUpdated records a successful edit.
func (r *LazadaProductUpdateRepository) MarkUpdated(ctx context.Context, sellerSku string) error {
	const q = `UPDATE lazada_product_update
	           SET sync_status = $1, last_error = '', updated_at = now()
	           WHERE seller_sku = $2`
	if _, err := r.db.ExecContext(ctx, q, ProductUpdateUpdated, sellerSku); err != nil {
		return fmt.Errorf("mark product updated: %w", err)
	}
	return nil
}

// MarkFailed records a failed edit with its reason.
func (r *LazadaProductUpdateRepository) MarkFailed(ctx context.Context, sellerSku, errMsg string) error {
	const q = `UPDATE lazada_product_update
	           SET sync_status = $1, last_error = $2, updated_at = now()
	           WHERE seller_sku = $3`
	if _, err := r.db.ExecContext(ctx, q, ProductUpdateFailed, errMsg, sellerSku); err != nil {
		return fmt.Errorf("mark product update failed: %w", err)
	}
	return nil
}
