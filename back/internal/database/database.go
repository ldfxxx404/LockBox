package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func InitDB() *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", "./app.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT UNIQUE NOT NULL,
			name TEXT NOT NULL,
			password TEXT NOT NULL,
			is_admin BOOLEAN NOT NULL DEFAULT 0,
			storage_limit INTEGER NOT NULL DEFAULT 10485760
		);`,
		`CREATE TABLE IF NOT EXISTS sessions (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			token TEXT UNIQUE NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);`,
		`CREATE TABLE IF NOT EXISTS files (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			filename TEXT NOT NULL,
			original_name TEXT NOT NULL,
			size INTEGER NOT NULL,
			mime_type TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);`,
	}

	for _, q := range queries {
		db.MustExec(q)
	}

	return db
}
