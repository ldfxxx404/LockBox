package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
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
	MinioUseSSL    = getEnv("MINIO_USE_SSL", false)
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
	case time.Duration:
		d, err := time.ParseDuration(val)
		if err == nil {
			return any(d).(T)
		}
	default:
		fmt.Printf("Unsupported type for key %s\n", key)
	}
	return fallback
}
