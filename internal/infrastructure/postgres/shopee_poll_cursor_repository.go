package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// ShopeePollCursor records the last successfully polled timestamp for a shop.
// The scheduler uses this to advance the window forward rather than always
// re-scanning from now-30m, and to provide a reproducible audit trail.
type ShopeePollCursor struct {
	ShopID    int64
	CursorAt  time.Time
	UpdatedAt time.Time
}

// ShopeePollCursorRepository persists per-shop poll cursors in PostgreSQL.
type ShopeePollCursorRepository struct {
	db *sql.DB
}

func NewShopeePollCursorRepository(db *sql.DB) *ShopeePollCursorRepository {
	return &ShopeePollCursorRepository{db: db}
}

// EnsureSchema creates the shopee_poll_cursors table if it does not exist.
func (r *ShopeePollCursorRepository) EnsureSchema(ctx context.Context) error {
	const q = `
	CREATE TABLE IF NOT EXISTS shopee_poll_cursors (
		shop_id    BIGINT      PRIMARY KEY,
		cursor_at  TIMESTAMPTZ NOT NULL,
		updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	)`
	if _, err := r.db.ExecContext(ctx, q); err != nil {
		return fmt.Errorf("ensure shopee_poll_cursors schema: %w", err)
	}
	return nil
}

// Load returns the cursor for a shop, or nil if this is the first run.
func (r *ShopeePollCursorRepository) Load(ctx context.Context, shopID int64) (*ShopeePollCursor, error) {
	const q = `SELECT shop_id, cursor_at, updated_at FROM shopee_poll_cursors WHERE shop_id = $1`
	c := &ShopeePollCursor{}
	err := r.db.QueryRowContext(ctx, q, shopID).Scan(&c.ShopID, &c.CursorAt, &c.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil // first run for this shop
	}
	if err != nil {
		return nil, fmt.Errorf("load poll cursor: %w", err)
	}
	return c, nil
}

// Save upserts the cursor position for a shop.
func (r *ShopeePollCursorRepository) Save(ctx context.Context, shopID int64, cursorAt time.Time) error {
	const q = `
	INSERT INTO shopee_poll_cursors (shop_id, cursor_at, updated_at)
	VALUES ($1, $2, NOW())
	ON CONFLICT (shop_id) DO UPDATE
		SET cursor_at  = EXCLUDED.cursor_at,
		    updated_at = NOW()`
	if _, err := r.db.ExecContext(ctx, q, shopID, cursorAt); err != nil {
		return fmt.Errorf("save poll cursor: %w", err)
	}
	return nil
}
