package config

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2/log"
)

var (
	ServerPort     = getEnv("SERVER_PORT", "0.0.0.0:5000")
	JWTSecret      = getEnv("JWT_SECRET", "supersecretkey")
	PostgresLink   = getEnv("POSTGRES_URL", "postgres://postgres:postgres@localhost:6969/lock_box?sslmode=disable")
	MigrationsDir  = getEnv("MIGRATIONS_DIR", "migrations")
	DbName         = getEnv("DB_NAME", "postgres")
	StorageLimit   = getEnv("STORAGE_LIMIT", 20)
	MinioEndpoint  = getEnv("MINIO_ENDPOINT", "localhost:9000")
	MinioAccessKey = getEnv("MINIO_ACCESS_KEY", "minioadmin")
	MinioSecretKey = getEnv("MINIO_SECRET_KEY", "minioadmin123")
	MinioBucket    = getEnv("MINIO_BUCKET", "uploads")
	LogLevel       = getEnv("LOG_LEVEL", "debug")
	LogFilePath    = getEnv("BACK_LOG_FILE", "/app/logs/")
	MinioUseSSL    = getEnv("MINIO_USE_SSL", false)
	FiberLimitBody = getEnv("BODY_LIMIT_MB", 500)
	FiberLimitReq  = getEnv("FIBER_LIMIT_REQ", 500)
)

func getEnv[T any](key string, fallback T) T {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	switch any(fallback).(type) {
	case string:
		return any(val).(T)
	case int:
		i, err := strconv.Atoi(val)
		if err == nil {
			return any(i).(T)
		}
	case float64:
		f, err := strconv.ParseFloat(val, 64)
		if err == nil {
			return any(f).(T)
		}
	case bool:
		b, err := strconv.ParseBool(val)
		if err == nil {
			return any(b).(T)
		}
	default:
		log.Fatalf("Unsupported type for key %s\n", key)
	}
	return fallback
}
