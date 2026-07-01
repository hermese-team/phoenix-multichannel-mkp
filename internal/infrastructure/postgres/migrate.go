package postgres

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5" // pgx5:// driver
	_ "github.com/golang-migrate/migrate/v4/source/file"     // file:// source

	"github.com/ascend/phoenix-multichannel-mkp/config"
)

// RunMigrations applies all pending up migrations from source (e.g. "file://migrations").
// Returns nil when there are no new migrations to apply.
func RunMigrations(cfg config.PostgresConfig, source string) error {
	m, err := migrate.New(source, cfg.MigrateDSN())
	if err != nil {
		return fmt.Errorf("create migrator: %w", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("run migrations: %w", err)
	}
	return nil
}
