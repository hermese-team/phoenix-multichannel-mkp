package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// ShopeeToken holds OAuth token data for a Shopee shop.
type ShopeeToken struct {
	ShopID       int64
	AccessToken  string
	RefreshToken string
	CreatedAt    time.Time
	ExpiredAt    time.Time
}

// ShopeeTokenRepository persists Shopee OAuth tokens in PostgreSQL.
type ShopeeTokenRepository struct {
	db *sql.DB
}

func NewShopeeTokenRepository(db *sql.DB) *ShopeeTokenRepository {
	return &ShopeeTokenRepository{db: db}
}

// EnsureSchema creates the shopee_tokens table if it does not exist.
func (r *ShopeeTokenRepository) EnsureSchema(ctx context.Context) error {
	const q = `
	CREATE TABLE IF NOT EXISTS shopee_tokens (
		shop_id       BIGINT      PRIMARY KEY,
		access_token  TEXT        NOT NULL,
		refresh_token TEXT        NOT NULL,
		created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
		expired_at    TIMESTAMPTZ NOT NULL
	)`
	if _, err := r.db.ExecContext(ctx, q); err != nil {
		return fmt.Errorf("ensure shopee_tokens schema: %w", err)
	}
	return nil
}

// Upsert inserts or updates the token for a shop.
func (r *ShopeeTokenRepository) Upsert(ctx context.Context, shopID int64, accessToken, refreshToken string, expiredAt time.Time) error {
	const q = `
	INSERT INTO shopee_tokens (shop_id, access_token, refresh_token, created_at, expired_at)
	VALUES ($1, $2, $3, now(), $4)
	ON CONFLICT (shop_id) DO UPDATE SET
		access_token  = EXCLUDED.access_token,
		refresh_token = EXCLUDED.refresh_token,
		created_at    = now(),
		expired_at    = EXCLUDED.expired_at`
	if _, err := r.db.ExecContext(ctx, q, shopID, accessToken, refreshToken, expiredAt); err != nil {
		return fmt.Errorf("upsert shopee token: %w", err)
	}
	return nil
}

// FindByShopID returns the token record for a shop, or nil if not found.
func (r *ShopeeTokenRepository) FindByShopID(ctx context.Context, shopID int64) (*ShopeeToken, error) {
	const q = `
	SELECT shop_id, access_token, refresh_token, created_at, expired_at
	FROM shopee_tokens
	WHERE shop_id = $1`
	row := r.db.QueryRowContext(ctx, q, shopID)
	var t ShopeeToken
	if err := row.Scan(&t.ShopID, &t.AccessToken, &t.RefreshToken, &t.CreatedAt, &t.ExpiredAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("find shopee token: %w", err)
	}
	return &t, nil
}
