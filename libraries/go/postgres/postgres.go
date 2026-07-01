package postgres

import (
	"context"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PoolOptions configures the pgx connection pool.
type PoolOptions struct {
	DSN             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

// Connect creates a pgxpool.Pool and verifies connectivity with Ping.
// Callers must defer pool.Close().
func Connect(ctx context.Context, opts PoolOptions) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(opts.DSN)
	if err != nil {
		return nil, err
	}
	if opts.MaxOpenConns > 0 {
		cfg.MaxConns = int32(opts.MaxOpenConns)
	}
	if opts.ConnMaxLifetime > 0 {
		cfg.MaxConnLifetime = opts.ConnMaxLifetime
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}
	return pool, nil
}

// RunMigrations applies all pending up-migrations.
// migrateDSN is a postgres:// URL; source is typically "file://migrations".
// ErrNoChange is treated as success.
func RunMigrations(migrateDSN, source string) error {
	m, err := migrate.New(source, migrateDSN)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
