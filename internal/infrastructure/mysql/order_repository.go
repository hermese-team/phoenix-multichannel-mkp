package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/okdev/marketplace-sync/internal"
	"github.com/okdev/marketplace-sync/internal/domain/order"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) FindBySellChannelID(ctx context.Context, marketplaceID, marketplaceType string) (*order.Order, error) {
	const q = `SELECT id, shop_id, marketplace_id, marketplace_type, status, total_amount, created_at, updated_at
	           FROM orders WHERE marketplace_id = ? AND marketplace_type = ?`
	o := &order.Order{}
	err := r.db.QueryRowContext(ctx, q, marketplaceID, marketplaceType).Scan(
		&o.ID, &o.ShopID, &o.SellChannelID, &o.SellChannelType,
		&o.Status, &o.TotalAmount, &o.CreatedAt, &o.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, internal.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("query order: %w", err)
	}
	return o, nil
}

func (r *OrderRepository) Save(ctx context.Context, o *order.Order) error {
	const q = `INSERT INTO orders (shop_id, marketplace_id, marketplace_type, status, total_amount, created_at, updated_at)
	           VALUES (?, ?, ?, ?, ?, NOW(), NOW())`
	result, err := r.db.ExecContext(ctx, q, o.ShopID, o.SellChannelID, o.SellChannelType, o.Status, o.TotalAmount)
	if err != nil {
		return fmt.Errorf("insert order: %w", err)
	}
	o.ID, _ = result.LastInsertId()
	return nil
}

func (r *OrderRepository) Update(ctx context.Context, o *order.Order) error {
	const q = `UPDATE orders SET status=?, updated_at=NOW() WHERE id=?`
	_, err := r.db.ExecContext(ctx, q, o.Status, o.ID)
	if err != nil {
		return fmt.Errorf("update order: %w", err)
	}
	return nil
}
