package main

import (
	"back/config"
	"back/internal/app"

	_ "back/cmd/swagger"

	"github.com/gofiber/fiber/v2/log"
)

// @title        LockBox API
// @version      0.2.2
// @description  This is the API documentation for LockBox SaaS app.
// @host         localhost:5000
// @BasePath     /api
// TODO: fix http to https when it need!!!
// @schemes      http
func main() {
	app := app.NewApp()
	log.Fatal(app.Listen(config.ServerPort))
}
