package database

import (
	"database/sql"
	"embed"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var EmbedMigrations embed.FS

func DBMigrations(db *sql.DB) error {

	goose.SetBaseFS(EmbedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	if err := goose.Up(db, "migrations"); err != nil {
		return err
	}
	if err := goose.Status(db, "migrations"); err != nil {
		return err
	}

	return nil
}
