package config

import "os"

var (
	ServerPort   = getEnv("SERVER_PORT", "0.0.0.0:5000")
	JWTSecret    = getEnv("JWT_SECRET", "supersecretkey")
	StorageDir   = getEnv("STORAGE_DIR", "./storage")
	PostgresLink = getEnv("POSTGRES_URL", "postgres://postgres:postgres@localhost:6969/lock_box?sslmode=disable")
	StorageLimit = 10
)

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
