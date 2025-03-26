package database

import (
	"back/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

// sest autodel bakcn
func InitDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", config.PostgresLink)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        email TEXT UNIQUE NOT NULL,
        name TEXT NOT NULL,
        password TEXT NOT NULL,
        is_admin BOOLEAN NOT NULL DEFAULT FALSE,
        storage_limit INTEGER NOT NULL DEFAULT 10
    );`,
		`CREATE TABLE IF NOT EXISTS sessions (
        id SERIAL PRIMARY KEY,
        user_id INTEGER NOT NULL,
        token TEXT UNIQUE NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    );`,
		`CREATE TABLE IF NOT EXISTS files (
        id SERIAL PRIMARY KEY,
        user_id INTEGER NOT NULL,
        filename TEXT NOT NULL,
        original_name TEXT NOT NULL,
        size INTEGER NOT NULL,
        mime_type TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    );`,
	}

	for _, q := range queries {
		db.MustExec(q)
	}

	return db
}
