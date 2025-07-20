package app

import (
	"back/config"
	"back/internal/middleware"
	"back/internal/registry"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	swaber "github.com/swaggo/fiber-swagger"
)

func SetupRoutes(app *fiber.App, c *registry.Container) {
	app.Get("/docs/*", swaber.WrapHandler)
	app.Use(limiter.New(config.Limiter))

	app.Post("/api/register", c.AuthHandler.Register)
	app.Post("/api/login", c.AuthHandler.Login)
	app.Post("/api/logout", c.AuthHandler.Logout)

	api := app.Group("/api", middleware.AuthRequired())
	api.Get("/storage", c.FileHandler.ListFiles)
	api.Get("/storage/:filename", c.FileHandler.Download)
	api.Get("/profile", c.ProfileHandler.GetProfile)
	api.Post("/upload", c.FileHandler.Upload)
	api.Delete("/delete/:filename", c.FileHandler.Delete)

	admin := api.Group("/admin", middleware.AdminRequired())
	admin.Get("/users", c.AdminHandler.GetAllUsers)
	admin.Get("/admins", c.AdminHandler.GetAllAdminUsers)
	admin.Put("/update_limit", c.AdminHandler.UpdateStorageLimit)
	admin.Put("/make_admin/:user_id", c.AdminHandler.MakeAdmin)
	admin.Put("/revoke_admin/:user_id", c.AdminHandler.RevokeAdmin)

	v2 := app.Group("/api/v2", middleware.AuthRequired())
	v2.Get("/profile", c.ProfileV2.GetProfile)
}
