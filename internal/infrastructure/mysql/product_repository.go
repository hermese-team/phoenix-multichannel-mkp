package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/okdev/marketplace-sync/internal"
	"github.com/okdev/marketplace-sync/internal/domain/product"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) FindByID(ctx context.Context, id int64) (*product.Product, error) {
	const q = `SELECT id, shop_id, name, sku, description, status, marketplace_id, marketplace_type, created_at, updated_at
	           FROM products WHERE id = ? AND deleted_at IS NULL`
	p := &product.Product{}
	err := r.db.QueryRowContext(ctx, q, id).Scan(
		&p.ID, &p.ShopID, &p.Name, &p.SKU, &p.Description,
		&p.Status, &p.SellChannelID, &p.SellChannelType, &p.CreatedAt, &p.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, internal.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("query product: %w", err)
	}
	return p, nil
}

func (r *ProductRepository) FindByShopID(ctx context.Context, shopID int64, limit, offset int) ([]*product.Product, error) {
	const q = `SELECT id, shop_id, name, sku, description, status, marketplace_id, marketplace_type, created_at, updated_at
	           FROM products WHERE shop_id = ? AND deleted_at IS NULL LIMIT ? OFFSET ?`
	rows, err := r.db.QueryContext(ctx, q, shopID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("query products by shop: %w", err)
	}
	defer rows.Close()
	return scanProducts(rows)
}

func (r *ProductRepository) FindPendingSync(ctx context.Context, marketplaceType string, limit int) ([]*product.Product, error) {
	const q = `SELECT id, shop_id, name, sku, description, status, marketplace_id, marketplace_type, created_at, updated_at
	           FROM products WHERE marketplace_type = ? AND marketplace_id = '' AND status = 'active' AND deleted_at IS NULL LIMIT ?`
	rows, err := r.db.QueryContext(ctx, q, marketplaceType, limit)
	if err != nil {
		return nil, fmt.Errorf("query pending products: %w", err)
	}
	defer rows.Close()
	return scanProducts(rows)
}

func (r *ProductRepository) Save(ctx context.Context, p *product.Product) error {
	const q = `INSERT INTO products (shop_id, name, sku, description, status, marketplace_id, marketplace_type, created_at, updated_at)
	           VALUES (?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`
	result, err := r.db.ExecContext(ctx, q, p.ShopID, p.Name, p.SKU, p.Description, p.Status, p.SellChannelID, p.SellChannelType)
	if err != nil {
		return fmt.Errorf("insert product: %w", err)
	}
	p.ID, _ = result.LastInsertId()
	return nil
}

func (r *ProductRepository) Update(ctx context.Context, p *product.Product) error {
	const q = `UPDATE products SET name=?, sku=?, description=?, status=?, marketplace_id=?, updated_at=NOW() WHERE id=?`
	_, err := r.db.ExecContext(ctx, q, p.Name, p.SKU, p.Description, p.Status, p.SellChannelID, p.ID)
	if err != nil {
		return fmt.Errorf("update product: %w", err)
	}
	return nil
}

func scanProducts(rows *sql.Rows) ([]*product.Product, error) {
	var products []*product.Product
	for rows.Next() {
		p := &product.Product{}
		if err := rows.Scan(
			&p.ID, &p.ShopID, &p.Name, &p.SKU, &p.Description,
			&p.Status, &p.SellChannelID, &p.SellChannelType, &p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan product: %w", err)
		}
		products = append(products, p)
	}
	return products, rows.Err()
}
