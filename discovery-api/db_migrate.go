package main

import (
	"database/sql"
	"embed"
	"log/slog"

	"github.com/pressly/goose/v3"
)

//go:embed sql/schema/*.sql
var embedMigrations embed.FS

func migrateDb(db *sql.DB) error {
	slog.Info("Checking for migrations to run...")

	goose.SetBaseFS(embedMigrations)

	err := goose.SetDialect("postgres")
	if err != nil {
		slog.Error("Failed to set postgres dialect for goose migrations", "err", err)
		return err
	}

	err = goose.Up(db, "sql/schema")
	if err != nil {
		slog.Error("Error when running goose UP migrations", "err", err)
		return err
	}

	return nil
}
