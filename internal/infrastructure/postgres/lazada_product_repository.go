package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

// Product sync states.
const (
	ProductSyncPending = "pending"
	ProductSyncSynced  = "synced"
	ProductSyncFailed  = "failed"
)

// LazadaProduct is a mock row of a product whose price/stock we push to Lazada.
// Price is DOUBLE PRECISION in the mock table; real money should use integer
// minor units or a decimal type.
type LazadaProduct struct {
	SellerSku string
	Price     float64
	Quantity  int
}

// LazadaProductRepository holds products pending a price/stock push (test schema).
type LazadaProductRepository struct {
	db *sql.DB
}

func NewLazadaProductRepository(db *sql.DB) *LazadaProductRepository {
	return &LazadaProductRepository{db: db}
}

// EnsureSchema creates the mock table if it does not exist (test convenience;
// a real deployment would use a migration).
func (r *LazadaProductRepository) EnsureSchema(ctx context.Context) error {
	const q = `
	CREATE TABLE IF NOT EXISTS lazada_product_sync (
		seller_sku  TEXT             PRIMARY KEY,
		price       DOUBLE PRECISION NOT NULL,
		quantity    INTEGER          NOT NULL,
		sync_status TEXT             NOT NULL DEFAULT 'pending',
		last_error  TEXT             NOT NULL DEFAULT '',
		updated_at  TIMESTAMPTZ      NOT NULL DEFAULT now()
	)`
	if _, err := r.db.ExecContext(ctx, q); err != nil {
		return fmt.Errorf("ensure lazada product schema: %w", err)
	}
	return nil
}

// ListPending returns up to limit products awaiting a price/stock push, oldest
// first.
func (r *LazadaProductRepository) ListPending(ctx context.Context, limit int) ([]LazadaProduct, error) {
	const q = `SELECT seller_sku, price, quantity FROM lazada_product_sync
	           WHERE sync_status = $1 ORDER BY updated_at LIMIT $2`
	rows, err := r.db.QueryContext(ctx, q, ProductSyncPending, limit)
	if err != nil {
		return nil, fmt.Errorf("list pending products: %w", err)
	}
	defer rows.Close()

	var out []LazadaProduct
	for rows.Next() {
		var p LazadaProduct
		if err := rows.Scan(&p.SellerSku, &p.Price, &p.Quantity); err != nil {
			return nil, fmt.Errorf("scan product: %w", err)
		}
		out = append(out, p)
	}
	return out, rows.Err()
}

// MarkResult records the outcome of a push for one seller_sku.
func (r *LazadaProductRepository) MarkResult(ctx context.Context, sellerSku, status, errMsg string) error {
	const q = `UPDATE lazada_product_sync
	           SET sync_status = $1, last_error = $2, updated_at = now()
	           WHERE seller_sku = $3`
	if _, err := r.db.ExecContext(ctx, q, status, errMsg, sellerSku); err != nil {
		return fmt.Errorf("mark product result: %w", err)
	}
	return nil
}
