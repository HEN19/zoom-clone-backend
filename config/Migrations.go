package config

import (
	"log"

	"github.com/gobuffalo/packr/v2"
	migrate "github.com/rubenv/sql-migrate"
)

func RunMigrations() {
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "../migrations"), // Path to migration files
	}

	// Apply all migrations
	n, err := migrate.Exec(DB, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("Could not run migrations: %v", err)
	}

	log.Printf("Applied %d migrations!\n", n)
}
