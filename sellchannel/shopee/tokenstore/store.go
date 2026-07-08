// Package tokenstore persists Shopee OAuth tokens.
// PostgreSQL is the source of truth; Redis is a fast-path cache.
package tokenstore

import (
	"context"
	"fmt"
	"time"

	pgAdapter    "github.com/okdev/marketplace-sync/internal/infrastructure/postgres"
	redisAdapter "github.com/okdev/marketplace-sync/internal/infrastructure/redis"
	"github.com/okdev/marketplace-sync/pkg/logger"
)

const (
	keyAccessFmt     = "shopee:%d:access_token"
	keyRefreshFmt    = "shopee:%d:refresh_token"
	warnThreshold    = 24 * time.Hour
	refreshThreshold = 30 * time.Minute // auto-refresh when < 30 min until expiry
)

// Refresher exchanges a Shopee refresh token for a new access/refresh token pair.
// Implemented by client.Client via the shopeeRefresher adapter in shopee.go.
type Refresher interface {
	RefreshToken(ctx context.Context, shopID int64, refreshToken string) (access, refresh string, expireIn int64, err error)
}

type Store struct {
	repo      *pgAdapter.ShopeeTokenRepository
	rdb       *redisAdapter.Client
	refresher Refresher // nil = no auto-refresh
}

func New(repo *pgAdapter.ShopeeTokenRepository, rdb *redisAdapter.Client, refresher Refresher) *Store {
	return &Store{repo: repo, rdb: rdb, refresher: refresher}
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
// If the token is expired or expiring within 30 min, it auto-refreshes using the refresh token.
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

	// Auto-refresh when expired or expiring soon.
	if (ttl <= 0 || ttl < refreshThreshold) && s.refresher != nil && rec.RefreshToken != "" {
		newAccess, newRefresh, expireIn, rfErr := s.refresher.RefreshToken(ctx, shopID, rec.RefreshToken)
		if rfErr != nil {
			logger.ErrorContext(ctx, "token auto-refresh failed",
				"event", "tokenstore.token.refresh.error",
				"shop_id", shopID,
				"error", rfErr,
			)
		} else if newAccess != "" {
			newExpiry := time.Now().Add(time.Duration(expireIn) * time.Second)
			if saveErr := s.Set(ctx, shopID, newAccess, newRefresh, newExpiry); saveErr != nil {
				logger.WarnContext(ctx, "save refreshed token failed",
					"event", "tokenstore.token.refresh.save_error",
					"shop_id", shopID,
					"error", saveErr,
				)
			} else {
				logger.InfoContext(ctx, "token auto-refreshed",
					"event", "tokenstore.token.refreshed",
					"shop_id", shopID,
					"expires_at", newExpiry.Format(time.RFC3339),
				)
				return newAccess, newRefresh, nil
			}
		}
	}

	// Warn when token is still expired after refresh attempt (or no refresher configured).
	if ttl <= 0 {
		logger.WarnContext(ctx, "token expired — please renew via Seller Center",
			"event", "tokenstore.token.expired",
			"shop_id", shopID,
			"expired_at", rec.ExpiredAt.Format(time.RFC3339),
		)
	} else if ttl < warnThreshold {
		logger.WarnContext(ctx, "token expiring soon",
			"event", "tokenstore.token.expiring_soon",
			"shop_id", shopID,
			"expires_in_minutes", int(ttl.Minutes()),
			"expired_at", rec.ExpiredAt.Format(time.RFC3339),
		)
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
