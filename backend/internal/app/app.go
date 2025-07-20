package app

import (
	"back/config"
	"back/internal/registry"
	"back/internal/utils"
	"io"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func NewApp() *fiber.App {
	utils.ParseLoglevel()

	file, _ := os.OpenFile(config.LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	log.SetOutput(io.MultiWriter(os.Stdout, file))

	app := fiber.New(config.FiberConfig)
	log.Info("Fiber initialized")

	container := registry.NewContainer()
	SetupRoutes(app, container)

	return app
}
