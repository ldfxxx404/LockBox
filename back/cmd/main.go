package main

import (
	"back/config"
	"back/internal/database"
	"back/internal/handlers"
	"back/internal/middleware"
	"back/internal/repositories"
	"back/internal/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := database.InitDB()

	app := fiber.New()

	// Инициализируем репозитории
	userRepo := repositories.NewUserRepo(db)
	fileRepo := repositories.NewFileRepo(db)

	// Инициализируем сервисы
	authService := services.NewAuthService(userRepo)
	fileService := services.NewFileService(fileRepo)
	profileService := services.NewProfileService(userRepo, fileService)
	adminService := services.NewAdminService(userRepo)

	// Инициализируем хэндлеры
	authHandler := handlers.NewAuthHandler(authService)
	fileHandler := handlers.NewFileHandler(fileService)
	profileHandler := handlers.NewProfileHandler(profileService)
	adminHandler := handlers.NewAdminHandler(adminService)

	// Публичные маршруты
	app.Post("/api/register", authHandler.Register)
	app.Post("/api/login", authHandler.Login)
	app.Post("/api/logout", authHandler.Logout)

	// Защищённые маршруты (требуется JWT)
	api := app.Group("/api", middleware.AuthRequired())
	api.Post("/upload", fileHandler.Upload)
	api.Get("/storage", fileHandler.ListFiles)
	api.Get("/storage/:filename", fileHandler.Download)
	api.Delete("/delete/:filename", fileHandler.Delete)
	api.Get("/profile", profileHandler.GetProfile)

	// Маршруты для администраторов
	admin := api.Group("/admin", middleware.AdminRequired())
	admin.Get("/users", adminHandler.GetAllUsers)
	admin.Post("/update_limit", adminHandler.UpdateStorageLimit)
	admin.Post("/make_admin/:user_id", adminHandler.MakeAdmin)
	admin.Post("/revoke_admin/:user_id", adminHandler.RevokeAdmin)

	log.Fatal(app.Listen(config.ServerPort))
}
