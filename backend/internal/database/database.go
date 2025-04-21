package database

import (
	"back/config"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func InitDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", config.PostgresLink)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := runMigrations(db.DB, config.MigrationsDir); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	return db
}

func runMigrations(db *sql.DB, migrationsDir string) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(db, migrationsDir); err != nil {
		return err
	}

	log.Println("DB migrations applied")
	return nil
}
