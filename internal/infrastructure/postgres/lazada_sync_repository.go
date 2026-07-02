package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

// LazadaSyncedOrder is a mock row of a synced Lazada order (test schema).
type LazadaSyncedOrder struct {
	OrderNumber   string
	OrderID       int64
	Statuses      string
	PaymentMethod string
	ItemCount     int
}

// LazadaSyncRepository persists synced Lazada orders into a mock table.
type LazadaSyncRepository struct {
	db *sql.DB
}

func NewLazadaSyncRepository(db *sql.DB) *LazadaSyncRepository {
	return &LazadaSyncRepository{db: db}
}

// EnsureSchema creates the mock table if it does not exist (test convenience;
// a real deployment would use a migration).
func (r *LazadaSyncRepository) EnsureSchema(ctx context.Context) error {
	const q = `
	CREATE TABLE IF NOT EXISTS lazada_synced_orders (
		order_number   TEXT PRIMARY KEY,
		order_id       BIGINT      NOT NULL,
		statuses       TEXT        NOT NULL DEFAULT '',
		payment_method TEXT        NOT NULL DEFAULT '',
		item_count     INTEGER     NOT NULL DEFAULT 0,
		synced_at      TIMESTAMPTZ NOT NULL DEFAULT now()
	)`
	if _, err := r.db.ExecContext(ctx, q); err != nil {
		return fmt.Errorf("ensure lazada schema: %w", err)
	}
	return nil
}

// SaveOrder upserts a synced order keyed by order_number.
func (r *LazadaSyncRepository) SaveOrder(ctx context.Context, o LazadaSyncedOrder) error {
	const q = `
	INSERT INTO lazada_synced_orders (order_number, order_id, statuses, payment_method, item_count, synced_at)
	VALUES ($1, $2, $3, $4, $5, now())
	ON CONFLICT (order_number) DO UPDATE SET
		order_id       = EXCLUDED.order_id,
		statuses       = EXCLUDED.statuses,
		payment_method = EXCLUDED.payment_method,
		item_count     = EXCLUDED.item_count,
		synced_at      = now()`
	if _, err := r.db.ExecContext(ctx, q, o.OrderNumber, o.OrderID, o.Statuses, o.PaymentMethod, o.ItemCount); err != nil {
		return fmt.Errorf("save synced order: %w", err)
	}
	return nil
}
