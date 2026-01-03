// Package db handles database connnection and migrations.
package db

import (
	"embed"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

// Connect establishes a connection to the SQLite database at the given file path,
// applies any pending migrations, and returns the database handle.
func Connect(dbFilePath string) (ModelsFactory, error) {
	db, err := sqlx.Connect("sqlite3", dbFilePath)
	if err != nil {
		return nil, err
	}

	// Apply migrations
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("sqlite3"); err != nil {
		return nil, err
	}
	if err := goose.Up(db.DB, "migrations"); err != nil {
		return nil, err
	}

	return newModelsFactory(db), nil
}
