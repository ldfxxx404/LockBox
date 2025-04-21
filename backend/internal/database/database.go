package database

import (
	"back/config"
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func InitDB() *sqlx.DB {
	db, _ := sqlx.Connect(config.DbName, config.PostgresLink)

	runMigrations(db.DB, config.MigrationsDir)
	return db
}

func runMigrations(db *sql.DB, migrationsDir string) {
	goose.SetDialect(config.DbName)
	goose.Up(db, migrationsDir)
}
