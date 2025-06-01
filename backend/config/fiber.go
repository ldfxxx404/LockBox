package config

import (
	"github.com/gofiber/fiber/v2"
)

var FiberConfig = fiber.Config{
	BodyLimit: 20 * 1024 * 1024,
}
