package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
)

// Product-create sync states.
const (
	ProductCreatePending = "pending"
	ProductCreateCreated = "created"
	ProductCreateFailed  = "failed"
)

// LazadaCreateProduct is a mock row describing a product we want to create on
// Lazada. Attributes are category-specific (name→value) and stored as JSONB;
// Price is DOUBLE PRECISION for the mock (real money should use minor units).
type LazadaCreateProduct struct {
	SellerSku     string
	Name          string
	CategoryID    int64
	Price         float64
	Quantity      int
	PackageWeight float64
	PackageLength float64
	PackageWidth  float64
	PackageHeight float64
	ImageURL      string
	Attributes    map[string]string
}

// LazadaProductCreateRepository holds products pending creation on Lazada.
type LazadaProductCreateRepository struct {
	db *sql.DB
}

func NewLazadaProductCreateRepository(db *sql.DB) *LazadaProductCreateRepository {
	return &LazadaProductCreateRepository{db: db}
}

// EnsureSchema creates the mock table if it does not exist (test convenience;
// a real deployment would use a migration).
func (r *LazadaProductCreateRepository) EnsureSchema(ctx context.Context) error {
	const q = `
	CREATE TABLE IF NOT EXISTS lazada_product_create (
		seller_sku     TEXT             PRIMARY KEY,
		name           TEXT             NOT NULL,
		category_id    BIGINT           NOT NULL,
		price          DOUBLE PRECISION NOT NULL,
		quantity       INTEGER          NOT NULL,
		package_weight DOUBLE PRECISION NOT NULL DEFAULT 0,
		package_length DOUBLE PRECISION NOT NULL DEFAULT 0,
		package_width  DOUBLE PRECISION NOT NULL DEFAULT 0,
		package_height DOUBLE PRECISION NOT NULL DEFAULT 0,
		image_url      TEXT             NOT NULL DEFAULT '',
		attributes     JSONB            NOT NULL DEFAULT '{}',
		sync_status    TEXT             NOT NULL DEFAULT 'pending',
		item_id        BIGINT,
		last_error     TEXT             NOT NULL DEFAULT '',
		updated_at     TIMESTAMPTZ      NOT NULL DEFAULT now()
	)`
	if _, err := r.db.ExecContext(ctx, q); err != nil {
		return fmt.Errorf("ensure lazada product-create schema: %w", err)
	}
	return nil
}

// ListPending returns up to limit products awaiting creation, oldest first.
func (r *LazadaProductCreateRepository) ListPending(ctx context.Context, limit int) ([]LazadaCreateProduct, error) {
	const q = `SELECT seller_sku, name, category_id, price, quantity,
	                  package_weight, package_length, package_width, package_height,
	                  image_url, attributes
	           FROM lazada_product_create
	           WHERE sync_status = $1 ORDER BY updated_at LIMIT $2`
	rows, err := r.db.QueryContext(ctx, q, ProductCreatePending, limit)
	if err != nil {
		return nil, fmt.Errorf("list pending product-create: %w", err)
	}
	defer rows.Close()

	var out []LazadaCreateProduct
	for rows.Next() {
		var (
			p        LazadaCreateProduct
			attrsRaw []byte
		)
		if err := rows.Scan(
			&p.SellerSku, &p.Name, &p.CategoryID, &p.Price, &p.Quantity,
			&p.PackageWeight, &p.PackageLength, &p.PackageWidth, &p.PackageHeight,
			&p.ImageURL, &attrsRaw,
		); err != nil {
			return nil, fmt.Errorf("scan product-create: %w", err)
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

// MarkCreated records a successful creation and the new Lazada item id.
func (r *LazadaProductCreateRepository) MarkCreated(ctx context.Context, sellerSku string, itemID int64) error {
	const q = `UPDATE lazada_product_create
	           SET sync_status = $1, item_id = $2, last_error = '', updated_at = now()
	           WHERE seller_sku = $3`
	if _, err := r.db.ExecContext(ctx, q, ProductCreateCreated, itemID, sellerSku); err != nil {
		return fmt.Errorf("mark product created: %w", err)
	}
	return nil
}

// MarkFailed records a failed creation attempt with its reason.
func (r *LazadaProductCreateRepository) MarkFailed(ctx context.Context, sellerSku, errMsg string) error {
	const q = `UPDATE lazada_product_create
	           SET sync_status = $1, last_error = $2, updated_at = now()
	           WHERE seller_sku = $3`
	if _, err := r.db.ExecContext(ctx, q, ProductCreateFailed, errMsg, sellerSku); err != nil {
		return fmt.Errorf("mark product failed: %w", err)
	}
	return nil
}
