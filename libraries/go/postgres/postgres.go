// Package postgres provides a shared pgx connection pool factory and
// golang-migrate runner for all phoenix-multichannel-mkp services.
//
// Usage:
//
//	pool, err := postgres.Connect(ctx, postgres.PoolOptions{
//	    DSN:             cfg.Postgres.DSN(),
//	    MaxOpenConns:    cfg.Postgres.MaxOpenConns,
//	    MaxIdleConns:    cfg.Postgres.MaxIdleConns,
//	    ConnMaxLifetime: cfg.Postgres.ConnMaxLifetime,
//	})
//	defer pool.Close()
//
//	err = postgres.RunMigrations(cfg.Postgres.MigrateDSN(), "file://migrations")
package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5" // pgx5:// driver
	_ "github.com/golang-migrate/migrate/v4/source/file"     // file:// source
	"github.com/jackc/pgx/v5/pgxpool"
)

// PoolOptions configures the pgx connection pool.
type PoolOptions struct {
	// DSN is a libpq-style connection string or pgx URL.
	// Example: "host=localhost port=5432 user=app password=secret dbname=mydb sslmode=disable"
	DSN string

	// MaxOpenConns is the maximum number of open connections (pgxpool MaxConns).
	MaxOpenConns int

	// MaxIdleConns is kept as the minimum number of idle connections (pgxpool MinConns).
	MaxIdleConns int

	// ConnMaxLifetime is the maximum lifetime of a connection before it is recycled.
	ConnMaxLifetime time.Duration
}

// Connect creates a pgx pool, applies pool sizing, and verifies with a 5-second ping.
// Callers must defer pool.Close().
func Connect(ctx context.Context, opts PoolOptions) (*pgxpool.Pool, error) {
	poolCfg, err := pgxpool.ParseConfig(opts.DSN)
	if err != nil {
		return nil, fmt.Errorf("postgres: parse config: %w", err)
	}
	if opts.MaxOpenConns > 0 {
		poolCfg.MaxConns = int32(opts.MaxOpenConns)
	}
	if opts.MaxIdleConns > 0 {
		poolCfg.MinConns = int32(opts.MaxIdleConns)
	}
	if opts.ConnMaxLifetime > 0 {
		poolCfg.MaxConnLifetime = opts.ConnMaxLifetime
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("postgres: create pool: %w", err)
	}

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := pool.Ping(pingCtx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("postgres: ping: %w", err)
	}
	return pool, nil
}

// RunMigrations applies all pending up migrations from source.
// migrateDSN must use the pgx5:// scheme (e.g. "pgx5://user:pass@host/db?sslmode=disable").
// Returns nil when there are no new migrations to apply (ErrNoChange is swallowed).
func RunMigrations(migrateDSN, source string) error {
	m, err := migrate.New(source, migrateDSN)
	if err != nil {
		return fmt.Errorf("postgres: create migrator: %w", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("postgres: run migrations: %w", err)
	}
	return nil
}
