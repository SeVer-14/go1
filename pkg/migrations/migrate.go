package migrations

import (
	"database/sql"
	"log"

	"github.com/pressly/goose/v3"
)

func RunMigrations(db *sql.DB, dialect string, dir string) error {
	goose.SetBaseFS(nil) // Используем файловую систему ОС

	if err := goose.SetDialect(dialect); err != nil {
		return err
	}

	if err := goose.Up(db, dir); err != nil {
		return err
	}

	return nil
}

func Status(db *sql.DB, dir string) {
	if err := goose.Status(db, dir); err != nil {
		log.Fatalf("failed to get migration status: %v", err)
	}
}
