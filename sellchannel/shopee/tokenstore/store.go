// Package tokenstore persists Shopee OAuth tokens.
// PostgreSQL is the source of truth; Redis is a fast-path cache.
package tokenstore

import (
	"context"
	"fmt"
	"log"
	"time"

	pgAdapter    "github.com/okdev/marketplace-sync/internal/infrastructure/postgres"
	redisAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/redis"
)

const (
	keyAccessFmt     = "shopee:%d:access_token"
	keyRefreshFmt    = "shopee:%d:refresh_token"
	warnThreshold    = 24 * time.Hour // warn when token expires within 24h
)

type Store struct {
	repo *pgAdapter.ShopeeTokenRepository
	rdb  *redisAdapter.Client
}

func New(repo *pgAdapter.ShopeeTokenRepository, rdb *redisAdapter.Client) *Store {
	return &Store{repo: repo, rdb: rdb}
}

// Set saves tokens to PostgreSQL (source of truth) and warms Redis cache.
// expiredAt is the absolute expiry time returned by Shopee (expire_in seconds from now).
func (s *Store) Set(ctx context.Context, shopID int64, access, refresh string, expiredAt time.Time) error {
	// 1. persist to Postgres
	if err := s.repo.Upsert(ctx, shopID, access, refresh, expiredAt); err != nil {
		return fmt.Errorf("persist token: %w", err)
	}
	// 2. warm Redis cache
	ttl := time.Until(expiredAt)
	if ttl > 30*time.Second {
		ttl -= 30 * time.Second // evict from Redis 30s before actual expiry
	}
	if ttl > 0 {
		_ = s.rdb.SetMany(ctx, ttl, map[string]string{
			fmt.Sprintf(keyAccessFmt, shopID):  access,
			fmt.Sprintf(keyRefreshFmt, shopID): refresh,
		})
	}
	return nil
}

// Get returns the access and refresh tokens for a shop.
// It checks Redis first; on miss it falls back to Postgres and re-warms Redis.
// Logs a warning when the token is expiring soon — the caller should renew it via Vault/UI.
func (s *Store) Get(ctx context.Context, shopID int64) (access, refresh string, err error) {
	// fast path: Redis
	access, _, err = s.rdb.GetString(ctx, fmt.Sprintf(keyAccessFmt, shopID))
	if err == nil && access != "" {
		refresh, _, _ = s.rdb.GetString(ctx, fmt.Sprintf(keyRefreshFmt, shopID))
		return access, refresh, nil
	}

	// slow path: Postgres
	rec, err := s.repo.FindByShopID(ctx, shopID)
	if err != nil {
		return "", "", fmt.Errorf("load token from db: %w", err)
	}
	if rec == nil || rec.AccessToken == "" {
		return "", "", nil // not found
	}

	ttl := time.Until(rec.ExpiredAt)

	// Warn when token is expired or expiring soon — renew via UI and update Vault.
	if ttl <= 0 {
		log.Printf("[tokenstore] WARNING: token for shop %d EXPIRED at %s — please renew via Seller Center",
			shopID, rec.ExpiredAt.Format(time.RFC3339))
	} else if ttl < warnThreshold {
		log.Printf("[tokenstore] WARNING: token for shop %d expires in %.0f minutes (%s) — please renew via Seller Center",
			shopID, ttl.Minutes(), rec.ExpiredAt.Format(time.RFC3339))
	}

	// re-warm Redis if token not yet expired
	if ttl > 30*time.Second {
		_ = s.rdb.SetMany(ctx, ttl-30*time.Second, map[string]string{
			fmt.Sprintf(keyAccessFmt, shopID):  rec.AccessToken,
			fmt.Sprintf(keyRefreshFmt, shopID): rec.RefreshToken,
		})
	}

	return rec.AccessToken, rec.RefreshToken, nil
}
