package order

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	entity "github.com/ascend/phoenix-multichannel-mkp/domain"
	"github.com/ascend/phoenix-multichannel-mkp/internal/repository"
	domainerrors "github.com/ascend/phoenix-multichannel-mkp/pkg/errors"
)

// OrderRepository is a Postgres-backed order store.
type OrderRepository struct {
	pool *pgxpool.Pool
}

func NewOrderRepository(pool *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{pool: pool}
}

func (r *OrderRepository) Create(ctx context.Context, o *entity.Order) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	const insertOrder = `
		INSERT INTO orders (id, buyer_id, seller_id, channel, status, total_price, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	if _, err := tx.Exec(ctx, insertOrder,
		o.ID, o.BuyerID, o.SellerID, o.Channel, o.Status, o.TotalPrice, o.CreatedAt, o.UpdatedAt,
	); err != nil {
		return err
	}

	const insertItem = `
		INSERT INTO order_items (order_id, product_id, name, price, quantity)
		VALUES ($1, $2, $3, $4, $5)`
	for _, item := range o.Items {
		if _, err := tx.Exec(ctx, insertItem, o.ID, item.ProductID, item.Name, item.Price, item.Quantity); err != nil {
			return err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetByID(ctx context.Context, id string) (*entity.Order, error) {
	const q = `
		SELECT id, buyer_id, seller_id, channel, status, total_price, created_at, updated_at
		FROM orders WHERE id = $1`
	var o entity.Order
	err := r.pool.QueryRow(ctx, q, id).Scan(
		&o.ID, &o.BuyerID, &o.SellerID, &o.Channel, &o.Status, &o.TotalPrice, &o.CreatedAt, &o.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domainerrors.New("NOT_FOUND", "order not found", domainerrors.ErrNotFound)
		}
		return nil, err
	}

	items, err := r.loadItems(ctx, id)
	if err != nil {
		return nil, err
	}
	o.Items = items
	return &o, nil
}

func (r *OrderRepository) Update(ctx context.Context, o *entity.Order) error {
	const q = `
		UPDATE orders SET status=$2, total_price=$3, updated_at=$4 WHERE id=$1`
	tag, err := r.pool.Exec(ctx, q, o.ID, o.Status, o.TotalPrice, o.UpdatedAt)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domainerrors.New("NOT_FOUND", "order not found", domainerrors.ErrNotFound)
	}
	return nil
}

func (r *OrderRepository) List(ctx context.Context, filter repository.OrderFilter) ([]*entity.Order, int, error) {
	where := "WHERE 1=1"
	args := []any{}
	idx := 1
	if filter.BuyerID != "" {
		where += fmt.Sprintf(" AND buyer_id=$%d", idx)
		args = append(args, filter.BuyerID)
		idx++
	}
	if filter.SellerID != "" {
		where += fmt.Sprintf(" AND seller_id=$%d", idx)
		args = append(args, filter.SellerID)
		idx++
	}
	if filter.Status != "" {
		where += fmt.Sprintf(" AND status=$%d", idx)
		args = append(args, filter.Status)
		idx++
	}
	if filter.Channel != "" {
		where += fmt.Sprintf(" AND channel=$%d", idx)
		args = append(args, filter.Channel)
		idx++
	}

	var total int
	if err := r.pool.QueryRow(ctx, "SELECT count(*) FROM orders "+where, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	page, limit := filter.Page, filter.Limit
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 20
	}
	offset := (page - 1) * limit

	q := fmt.Sprintf(`
		SELECT id, buyer_id, seller_id, channel, status, total_price, created_at, updated_at
		FROM orders %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d`, where, idx, idx+1)
	args = append(args, limit, offset)

	rows, err := r.pool.Query(ctx, q, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var orders []*entity.Order
	for rows.Next() {
		var o entity.Order
		if err := rows.Scan(
			&o.ID, &o.BuyerID, &o.SellerID, &o.Channel, &o.Status, &o.TotalPrice, &o.CreatedAt, &o.UpdatedAt,
		); err != nil {
			return nil, 0, err
		}
		orders = append(orders, &o)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}
	return orders, total, nil
}

func (r *OrderRepository) loadItems(ctx context.Context, orderID string) ([]entity.OrderItem, error) {
	const q = `SELECT product_id, name, price, quantity FROM order_items WHERE order_id = $1`
	rows, err := r.pool.Query(ctx, q, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []entity.OrderItem
	for rows.Next() {
		var it entity.OrderItem
		if err := rows.Scan(&it.ProductID, &it.Name, &it.Price, &it.Quantity); err != nil {
			return nil, err
		}
		items = append(items, it)
	}
	return items, rows.Err()
}
