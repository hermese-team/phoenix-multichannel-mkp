package postgres

import (
	"context"
	"database/sql"

	"github.com/okdev/marketplace-sync/internal"
	"github.com/okdev/marketplace-sync/internal/domain/order"
	applog "github.com/okdev/marketplace-sync/pkg/logger"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// EnsureSchema creates the orders table if it does not exist.
// Safe to call on every startup — idempotent.
func (r *OrderRepository) EnsureSchema(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS orders (
			id               BIGSERIAL PRIMARY KEY,
			shop_id          BIGINT       NOT NULL,
			marketplace_id   TEXT         NOT NULL,
			marketplace_type TEXT         NOT NULL,
			status           TEXT         NOT NULL,
			total_amount     NUMERIC(18,2) NOT NULL DEFAULT 0,
			created_at       TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
			updated_at       TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
			UNIQUE (marketplace_id, marketplace_type)
		)
	`)
	return err
}

// Upsert inserts a new order or updates an existing one idempotently.
// The unique key is (marketplace_id, marketplace_type).
// Returns isNew=true when the row was inserted for the first time.
// Uses the PostgreSQL xmax trick: xmax=0 after INSERT, non-zero after UPDATE.
func (r *OrderRepository) Upsert(ctx context.Context, o *order.Order) (isNew bool, err error) {
	const q = `
		INSERT INTO orders (shop_id, marketplace_id, marketplace_type, status, total_amount, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		ON CONFLICT (marketplace_id, marketplace_type) DO UPDATE
			SET status       = EXCLUDED.status,
			    total_amount = EXCLUDED.total_amount,
			    updated_at   = NOW()
		RETURNING id, (xmax = 0) AS is_new, created_at, updated_at
	`
	err = r.db.QueryRowContext(ctx, q,
		o.ShopID, o.SellChannelID, o.SellChannelType, string(o.Status), o.TotalAmount,
	).Scan(&o.ID, &isNew, &o.CreatedAt, &o.UpdatedAt)
	return isNew, err
}

func (r *OrderRepository) FindBySellChannelID(ctx context.Context, marketplaceID, marketplaceType string) (*order.Order, error) {
	const q = `SELECT id, shop_id, marketplace_id, marketplace_type, status, total_amount, created_at, updated_at
	           FROM orders WHERE marketplace_id = $1 AND marketplace_type = $2`
	o := &order.Order{}
	err := r.db.QueryRowContext(ctx, q, marketplaceID, marketplaceType).Scan(
		&o.ID, &o.ShopID, &o.SellChannelID, &o.SellChannelType,
		&o.Status, &o.TotalAmount, &o.CreatedAt, &o.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, internal.ErrNotFound
	}
	if err != nil {
		applog.Error("query order", "error", err)
		return nil, err
	}
	return o, nil
}

func (r *OrderRepository) Save(ctx context.Context, o *order.Order) error {
	const q = `INSERT INTO orders (shop_id, marketplace_id, marketplace_type, status, total_amount, created_at, updated_at)
	           VALUES ($1, $2, $3, $4, $5, NOW(), NOW()) RETURNING id`
	if err := r.db.QueryRowContext(ctx, q, o.ShopID, o.SellChannelID, o.SellChannelType, o.Status, o.TotalAmount).Scan(&o.ID); err != nil {
		applog.Error("insert order", "error", err)
		return err
	}
	return nil
}

func (r *OrderRepository) Update(ctx context.Context, o *order.Order) error {
	const q = `UPDATE orders SET status=$1, updated_at=NOW() WHERE id=$2`
	_, err := r.db.ExecContext(ctx, q, o.Status, o.ID)
	if err != nil {
		applog.Error("update order", "error", err)
		return err
	}
	return nil
}
