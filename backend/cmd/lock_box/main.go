package main

import (
	"back/config"
	"back/internal/database"
	"back/internal/handlers"
	"back/internal/middleware"
	"back/internal/repositories"
	"back/internal/services"
	"back/internal/utils"
	"io"
	"os"

	_ "back/cmd/swagger"

	"github.com/gofiber/fiber/v2/log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	swaber "github.com/swaggo/fiber-swagger"
)

// @title        LockBox API
// @version      0.1.1
// @description  This is the API documentation for LockBox SaaS app.
// @host         localhost:5000
// @BasePath     /api
// TODO: fix http to https when it need!!!
// @schemes      http
func main() {
	utils.ParseLoglevel()
	
	file, _ := os.OpenFile(config.LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	iw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(iw)

	db := database.InitDB()

	app := fiber.New(config.FiberConfig)
	log.Info("create new fiber config with settings")

	userRepo := repositories.NewUserRepo(db)
	fileRepo := repositories.NewFileRepo(db)
	log.Info("database init")

	authService := services.NewAuthService(userRepo)
	fileService, err := services.NewFileService(
		fileRepo,
		userRepo,
		config.MinioEndpoint,
		config.MinioAccessKey,
		config.MinioSecretKey,
		config.MinioBucket,
		config.MinioUseSSL,
	)
	if err != nil {
		log.Fatalf("failed to initialize FileService: %v", err)
	}

	profileService := services.NewProfileService(userRepo, fileService)
	adminService := services.NewAdminService(userRepo)
	log.Info("init services")

	authHandler := handlers.NewAuthHandler(authService)
	fileHandler := handlers.NewFileHandler(fileService)
	profileHandler := handlers.NewProfileHandler(profileService)
	adminHandler := handlers.NewAdminHandler(adminService)
	log.Info("init handlers")

	app.Get("/docs/*", swaber.WrapHandler)

	app.Use(limiter.New(config.Limiter))

	app.Post("/api/register", authHandler.Register)
	app.Post("/api/login", authHandler.Login)
	app.Post("/api/logout", authHandler.Logout)

	api := app.Group("/api", middleware.AuthRequired())
	api.Get("/storage", fileHandler.ListFiles)
	api.Get("/storage/:filename", fileHandler.Download)
	api.Get("/profile", profileHandler.GetProfile)
	api.Post("/upload", fileHandler.Upload)
	api.Delete("/delete/:filename", fileHandler.Delete)

	admin := api.Group("/admin", middleware.AdminRequired())
	admin.Get("/users", adminHandler.GetAllUsers)
	admin.Put("/update_limit", adminHandler.UpdateStorageLimit)
	admin.Put("/make_admin/:user_id", adminHandler.MakeAdmin)
	admin.Put("/revoke_admin/:user_id", adminHandler.RevokeAdmin)
	log.Info("init routes")

	log.Fatal(app.Listen(config.ServerPort))
}
