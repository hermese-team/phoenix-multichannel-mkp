package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

// Image sync states. Upload is persisted before Set so a failure resumes at Set
// (never re-uploads). There is no terminal "failed": a failure keeps the current
// status and is retried next run.
const (
	ImagePending  = "pending"  // not uploaded yet
	ImageUploaded = "uploaded" // uploaded to Lazada, awaiting SetImages
	ImageDone     = "done"     // images set on the product
)

// LazadaProductImage is a mock row: upload an image (from a local path or http
// URL) and attach the resulting Lazada URL to a product's sku.
type LazadaProductImage struct {
	ID          int64
	SellerSku   string
	Source      string // local file path or http(s) URL
	UploadedURL string // Lazada URL after upload
	SyncStatus  string
}

// LazadaProductImageRepository holds pending image work (test schema).
type LazadaProductImageRepository struct {
	db *sql.DB
}

func NewLazadaProductImageRepository(db *sql.DB) *LazadaProductImageRepository {
	return &LazadaProductImageRepository{db: db}
}

// EnsureSchema creates the mock table if it does not exist (test convenience;
// a real deployment would use a migration).
func (r *LazadaProductImageRepository) EnsureSchema(ctx context.Context) error {
	const q = `
	CREATE TABLE IF NOT EXISTS lazada_product_image (
		id           BIGSERIAL   PRIMARY KEY,
		seller_sku   TEXT        NOT NULL,
		source       TEXT        NOT NULL,
		uploaded_url TEXT        NOT NULL DEFAULT '',
		sync_status  TEXT        NOT NULL DEFAULT 'pending',
		last_error   TEXT        NOT NULL DEFAULT '',
		updated_at   TIMESTAMPTZ NOT NULL DEFAULT now()
	)`
	if _, err := r.db.ExecContext(ctx, q); err != nil {
		return fmt.Errorf("ensure lazada image schema: %w", err)
	}
	return nil
}

// ListPending returns up to limit rows still in progress (pending or uploaded),
// oldest first. Including "uploaded" lets a row resume at SetImages.
func (r *LazadaProductImageRepository) ListPending(ctx context.Context, limit int) ([]LazadaProductImage, error) {
	const q = `SELECT id, seller_sku, source, uploaded_url, sync_status
	           FROM lazada_product_image
	           WHERE sync_status IN ($1, $2) ORDER BY updated_at LIMIT $3`
	rows, err := r.db.QueryContext(ctx, q, ImagePending, ImageUploaded, limit)
	if err != nil {
		return nil, fmt.Errorf("list pending image: %w", err)
	}
	defer rows.Close()

	var out []LazadaProductImage
	for rows.Next() {
		var im LazadaProductImage
		if err := rows.Scan(&im.ID, &im.SellerSku, &im.Source, &im.UploadedURL, &im.SyncStatus); err != nil {
			return nil, fmt.Errorf("scan image: %w", err)
		}
		out = append(out, im)
	}
	return out, rows.Err()
}

// MarkUploaded stores the Lazada URL and advances to "uploaded" (clears error).
func (r *LazadaProductImageRepository) MarkUploaded(ctx context.Context, id int64, url string) error {
	const q = `UPDATE lazada_product_image
	           SET uploaded_url = $1, sync_status = $2, last_error = '', updated_at = now()
	           WHERE id = $3`
	if _, err := r.db.ExecContext(ctx, q, url, ImageUploaded, id); err != nil {
		return fmt.Errorf("mark image uploaded: %w", err)
	}
	return nil
}

// MarkDone advances to "done" (clears error).
func (r *LazadaProductImageRepository) MarkDone(ctx context.Context, id int64) error {
	const q = `UPDATE lazada_product_image
	           SET sync_status = $1, last_error = '', updated_at = now() WHERE id = $2`
	if _, err := r.db.ExecContext(ctx, q, ImageDone, id); err != nil {
		return fmt.Errorf("mark image done: %w", err)
	}
	return nil
}

// MarkError records a failure without changing status, so the row resumes next run.
func (r *LazadaProductImageRepository) MarkError(ctx context.Context, id int64, errMsg string) error {
	const q = `UPDATE lazada_product_image SET last_error = $1, updated_at = now() WHERE id = $2`
	if _, err := r.db.ExecContext(ctx, q, errMsg, id); err != nil {
		return fmt.Errorf("mark image error: %w", err)
	}
	return nil
}
