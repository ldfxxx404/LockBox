package database

import (
	"back/config"
	"database/sql"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func InitDB() *sqlx.DB {
	db, err := sqlx.Connect(config.DbName, config.PostgresLink)
	if err != nil {
		log.Fatalf("error of connect DB", err)
	}

	runMigrations(db.DB, config.MigrationsDir)
	return db
}

func runMigrations(db *sql.DB, migrationsDir string) {
	err := goose.SetDialect(config.DbName)
	if err != nil {
		log.Fatalf("error of migrate", err)
	}
	err = goose.Up(db, migrationsDir)
	if err != nil {
		log.Fatalf("error of init migrate", err)
	}

}
