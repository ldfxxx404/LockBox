package utils

import (
	"back/config"

	"github.com/gofiber/fiber/v2/log"
)

func ParseLoglevel() {
	level := config.LogLevel

	switch level {
	case "debug":
		log.SetLevel(log.LevelDebug)
	case "info":
		log.SetLevel(log.LevelInfo)
	case "error":
		log.SetLevel(log.LevelError)
	default:
		log.SetLevel(log.LevelWarn)
	}
}
