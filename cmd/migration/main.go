package main

import (
	"context"

	"github.com/scylladb/gocqlx/v2/migrate"

	"github.com/amin-mir/reporting/config"
	"github.com/amin-mir/reporting/scylla"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	manager := scylla.NewManager(cfg)

	if err = manager.CreateKeyspace(cfg.ScyllaKeyspace); err != nil {
		panic(err)
	}

	session, err := manager.Connect()
	if err != nil {
		panic(err)
	}

	if err = migrate.Migrate(ctx, session, cfg.ScyllaMigrationsDir); err != nil {
		panic(err)
	}
}
