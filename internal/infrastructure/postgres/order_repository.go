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
