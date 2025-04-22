// @title        LockBox API
// @version      1.0
// @description  This is the API documentation for LockBox SaaS app.
// @host         localhost:5000
// @BasePath     /api
// @schemes      https
package main

import (
	"back/config"
	"back/internal/database"
	"back/internal/handlers"
	"back/internal/middleware"
	"back/internal/repositories"
	"back/internal/services"
	"log"

	_ "back/cmd/swagger"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func main() {
	db := database.InitDB()

	app := fiber.New()

	userRepo := repositories.NewUserRepo(db)
	fileRepo := repositories.NewFileRepo(db)

	authService := services.NewAuthService(userRepo)
	fileService := services.NewFileService(fileRepo, userRepo)
	profileService := services.NewProfileService(userRepo, fileService)
	adminService := services.NewAdminService(userRepo)

	authHandler := handlers.NewAuthHandler(authService)
	fileHandler := handlers.NewFileHandler(fileService)
	profileHandler := handlers.NewProfileHandler(profileService)
	adminHandler := handlers.NewAdminHandler(adminService)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	app.Post("/api/register", authHandler.Register)
	app.Post("/api/login", authHandler.Login)
	app.Post("/api/logout", authHandler.Logout)

	api := app.Group("/api", middleware.AuthRequired())
	api.Post("/upload", fileHandler.Upload)
	api.Get("/storage", fileHandler.ListFiles)
	api.Get("/storage/:filename", fileHandler.Download)
	api.Delete("/delete/:filename", fileHandler.Delete)
	api.Get("/profile", profileHandler.GetProfile)

	admin := api.Group("/admin", middleware.AdminRequired())
	admin.Get("/users", adminHandler.GetAllUsers)
	admin.Post("/update_limit", adminHandler.UpdateStorageLimit)
	admin.Post("/make_admin/:user_id", adminHandler.MakeAdmin)
	admin.Post("/revoke_admin/:user_id", adminHandler.RevokeAdmin)

	log.Fatal(app.Listen(config.ServerPort))
}
