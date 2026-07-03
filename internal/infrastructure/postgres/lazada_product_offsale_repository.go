package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

// Off-sale actions + sync states.
const (
	ActionDeactivate    = "deactivate"     // /product/deactivate (whole product)
	ActionRemoveProduct = "remove_product" // /product/remove (seller skus)
	ActionRemoveSku     = "remove_sku"     // /product/sku/remove (skus of a product)

	OffsalePending = "pending"
	OffsaleDone    = "done"
	OffsaleFailed  = "failed"
)

// LazadaOffsale is a mock row describing a take-down action on Lazada. Which
// fields matter depends on Action: deactivate needs ItemID; remove_product needs
// SellerSkus; remove_sku needs ItemID + VariationName + SellerSkus.
type LazadaOffsale struct {
	ID            int64
	Action        string
	ItemID        int64
	SellerSkus    string // comma-separated seller_sku
	VariationName string // remove_sku only
}

// LazadaProductOffsaleRepository holds pending off-sale actions (test schema).
type LazadaProductOffsaleRepository struct {
	db *sql.DB
}

func NewLazadaProductOffsaleRepository(db *sql.DB) *LazadaProductOffsaleRepository {
	return &LazadaProductOffsaleRepository{db: db}
}

// EnsureSchema creates the mock table if it does not exist (test convenience;
// a real deployment would use a migration).
func (r *LazadaProductOffsaleRepository) EnsureSchema(ctx context.Context) error {
	const q = `
	CREATE TABLE IF NOT EXISTS lazada_product_offsale (
		id             BIGSERIAL   PRIMARY KEY,
		action         TEXT        NOT NULL,
		item_id        BIGINT      NOT NULL DEFAULT 0,
		seller_skus    TEXT        NOT NULL DEFAULT '',
		variation_name TEXT        NOT NULL DEFAULT '',
		sync_status    TEXT        NOT NULL DEFAULT 'pending',
		last_error     TEXT        NOT NULL DEFAULT '',
		updated_at     TIMESTAMPTZ NOT NULL DEFAULT now()
	)`
	if _, err := r.db.ExecContext(ctx, q); err != nil {
		return fmt.Errorf("ensure lazada offsale schema: %w", err)
	}
	return nil
}

// ListPending returns up to limit pending off-sale actions, oldest first.
func (r *LazadaProductOffsaleRepository) ListPending(ctx context.Context, limit int) ([]LazadaOffsale, error) {
	const q = `SELECT id, action, item_id, seller_skus, variation_name
	           FROM lazada_product_offsale
	           WHERE sync_status = $1 ORDER BY updated_at LIMIT $2`
	rows, err := r.db.QueryContext(ctx, q, OffsalePending, limit)
	if err != nil {
		return nil, fmt.Errorf("list pending offsale: %w", err)
	}
	defer rows.Close()

	var out []LazadaOffsale
	for rows.Next() {
		var o LazadaOffsale
		if err := rows.Scan(&o.ID, &o.Action, &o.ItemID, &o.SellerSkus, &o.VariationName); err != nil {
			return nil, fmt.Errorf("scan offsale: %w", err)
		}
		out = append(out, o)
	}
	return out, rows.Err()
}

// MarkDone records a successful action.
func (r *LazadaProductOffsaleRepository) MarkDone(ctx context.Context, id int64) error {
	const q = `UPDATE lazada_product_offsale
	           SET sync_status = $1, last_error = '', updated_at = now() WHERE id = $2`
	if _, err := r.db.ExecContext(ctx, q, OffsaleDone, id); err != nil {
		return fmt.Errorf("mark offsale done: %w", err)
	}
	return nil
}

// MarkFailed records a failed action with its reason.
func (r *LazadaProductOffsaleRepository) MarkFailed(ctx context.Context, id int64, errMsg string) error {
	const q = `UPDATE lazada_product_offsale
	           SET sync_status = $1, last_error = $2, updated_at = now() WHERE id = $3`
	if _, err := r.db.ExecContext(ctx, q, OffsaleFailed, errMsg, id); err != nil {
		return fmt.Errorf("mark offsale failed: %w", err)
	}
	return nil
}
