// Package tokenstore persists Lazada OAuth tokens in Redis so the access and
// refresh tokens survive restarts and can be rotated by the scheduler.
package tokenstore

import (
	"context"
	"time"

	"github.com/okdev/marketplace-sync/internal/infrastructure/redis"
)

const (
	keyAccess  = "lazada:access_token"
	keyRefresh = "lazada:refresh_token"
)

type Store struct {
	rdb *redis.Client
}

func New(rdb *redis.Client) *Store { return &Store{rdb: rdb} }

// Get returns the stored access/refresh tokens. Missing keys yield empty strings.
func (s *Store) Get(ctx context.Context) (access, refresh string, err error) {
	access, _, err = s.rdb.GetString(ctx, keyAccess)
	if err != nil {
		return "", "", err
	}
	refresh, _, err = s.rdb.GetString(ctx, keyRefresh)
	if err != nil {
		return "", "", err
	}
	return access, refresh, nil
}

// Set stores the access/refresh tokens. A ttl <= 0 means no expiry.
func (s *Store) Set(ctx context.Context, access, refresh string, ttl time.Duration) error {
	if err := s.rdb.Set(ctx, keyAccess, access, ttl); err != nil {
		return err
	}
	return s.rdb.Set(ctx, keyRefresh, refresh, ttl)
}
