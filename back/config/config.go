package config

import "os"

var (
	ServerPort         = getEnv("SERVER_PORT", "0.0.0.0:5000")
	JWTSecret          = getEnv("JWT_SECRET", "supersecretkey")
	StorageDir         = getEnv("STORAGE_DIR", "./storage")
	StorageLimit int64 = 10485760
)

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
