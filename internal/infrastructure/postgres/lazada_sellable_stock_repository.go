package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

// Sellable-stock actions + sync states.
//
// NOTE: StockAdjust is a non-idempotent delta. Auto-flow processes a row once, so
// this is safe, but manually resetting a done row back to pending would re-apply
// the delta. Prefer StockUpdate (absolute) when a row may be reprocessed.
const (
	StockAdjust = "adjust" // delta (+/-) via /product/stock/sellable/adjust
	StockUpdate = "update" // absolute via /product/stock/sellable/update

	SellableStockPending = "pending"
	SellableStockDone    = "done"
	SellableStockFailed  = "failed"
)

// LazadaSellableStock is a mock row describing a warehouse sellable-quantity
// change. Quantity is a delta for StockAdjust and an absolute value for StockUpdate.
type LazadaSellableStock struct {
	ID            int64
	Action        string
	ItemID        int64
	SkuID         int64
	WarehouseCode string
	Quantity      int
}

// LazadaSellableStockRepository holds pending sellable-stock changes (test schema).
type LazadaSellableStockRepository struct {
	db *sql.DB
}

func NewLazadaSellableStockRepository(db *sql.DB) *LazadaSellableStockRepository {
	return &LazadaSellableStockRepository{db: db}
}

// EnsureSchema creates the mock table if it does not exist (test convenience;
// a real deployment would use a migration).
func (r *LazadaSellableStockRepository) EnsureSchema(ctx context.Context) error {
	const q = `
	CREATE TABLE IF NOT EXISTS lazada_sellable_stock (
		id             BIGSERIAL   PRIMARY KEY,
		action         TEXT        NOT NULL,
		item_id        BIGINT      NOT NULL,
		sku_id         BIGINT      NOT NULL,
		warehouse_code TEXT        NOT NULL DEFAULT '',
		quantity       INTEGER     NOT NULL,
		sync_status    TEXT        NOT NULL DEFAULT 'pending',
		last_error     TEXT        NOT NULL DEFAULT '',
		updated_at     TIMESTAMPTZ NOT NULL DEFAULT now()
	)`
	if _, err := r.db.ExecContext(ctx, q); err != nil {
		return fmt.Errorf("ensure lazada sellable-stock schema: %w", err)
	}
	return nil
}

// ListPending returns up to limit pending stock changes, oldest first.
func (r *LazadaSellableStockRepository) ListPending(ctx context.Context, limit int) ([]LazadaSellableStock, error) {
	const q = `SELECT id, action, item_id, sku_id, warehouse_code, quantity
	           FROM lazada_sellable_stock
	           WHERE sync_status = $1 ORDER BY updated_at LIMIT $2`
	rows, err := r.db.QueryContext(ctx, q, SellableStockPending, limit)
	if err != nil {
		return nil, fmt.Errorf("list pending sellable-stock: %w", err)
	}
	defer rows.Close()

	var out []LazadaSellableStock
	for rows.Next() {
		var o LazadaSellableStock
		if err := rows.Scan(&o.ID, &o.Action, &o.ItemID, &o.SkuID, &o.WarehouseCode, &o.Quantity); err != nil {
			return nil, fmt.Errorf("scan sellable-stock: %w", err)
		}
		out = append(out, o)
	}
	return out, rows.Err()
}

// MarkDone records a successful change.
func (r *LazadaSellableStockRepository) MarkDone(ctx context.Context, id int64) error {
	const q = `UPDATE lazada_sellable_stock
	           SET sync_status = $1, last_error = '', updated_at = now() WHERE id = $2`
	if _, err := r.db.ExecContext(ctx, q, SellableStockDone, id); err != nil {
		return fmt.Errorf("mark sellable-stock done: %w", err)
	}
	return nil
}

// MarkFailed records a failed change with its reason.
func (r *LazadaSellableStockRepository) MarkFailed(ctx context.Context, id int64, errMsg string) error {
	const q = `UPDATE lazada_sellable_stock
	           SET sync_status = $1, last_error = $2, updated_at = now() WHERE id = $3`
	if _, err := r.db.ExecContext(ctx, q, SellableStockFailed, errMsg, id); err != nil {
		return fmt.Errorf("mark sellable-stock failed: %w", err)
	}
	return nil
}
