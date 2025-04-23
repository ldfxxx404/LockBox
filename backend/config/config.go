package config

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

var (
	ServerPort    = getEnv("SERVER_PORT", "0.0.0.0:5000")
	JWTSecret     = getEnv("JWT_SECRET", "supersecretkey")
	StorageDir    = getEnv("STORAGE_DIR", "./storage")
	PostgresLink  = getEnv("POSTGRES_URL", "postgres://postgres:postgres@localhost:6969/lock_box?sslmode=disable")
	StorageLimit  = 20
	MigrationsDir = "migrations"
	DbName        = "postgres"
	Limiter       = limiter.Config{
		Expiration: 10 * time.Second,
		Max:        3,
	}
	FiberConfig = fiber.Config{
		BodyLimit: 20 * 1024 * 1024,
	}
)

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
