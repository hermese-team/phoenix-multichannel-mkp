package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ascend/phoenix-multichannel-mkp/config"
	pglib "github.com/hermese-team/phoenix-multichannel-mkp/libraries/go/postgres"
)

// Connect delegates to the shared postgres library.
// Callers must defer pool.Close().
func Connect(ctx context.Context, cfg config.PostgresConfig) (*pgxpool.Pool, error) {
	return pglib.Connect(ctx, pglib.PoolOptions{
		DSN:             cfg.DSN(),
		MaxOpenConns:    cfg.MaxOpenConns,
		MaxIdleConns:    cfg.MaxIdleConns,
		ConnMaxLifetime: cfg.ConnMaxLifetime,
	})
}
