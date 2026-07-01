package postgres

import (
	"github.com/ascend/phoenix-multichannel-mkp/config"
	pglib "github.com/hermese-team/phoenix-multichannel-mkp/libraries/go/postgres"
)

// RunMigrations delegates to the shared postgres library.
// source is typically "file://migrations".
func RunMigrations(cfg config.PostgresConfig, source string) error {
	return pglib.RunMigrations(cfg.MigrateDSN(), source)
}
