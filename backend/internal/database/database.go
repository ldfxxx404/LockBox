package database

import (
	"back/config"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", config.PostgresLink)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db
}
