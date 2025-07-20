package registry

import (
	"back/config"
	"back/internal/database"
	"back/internal/handlers"
	handlersv2 "back/internal/handlers_v2"
	"back/internal/repositories"
	"back/internal/services"
)

type Container struct {
	AuthHandler    *handlers.AuthHandler
	FileHandler    *handlers.FileHandler
	ProfileHandler *handlers.ProfileHandler
	ProfileV2      *handlersv2.ProfileHandlerV2
	AdminHandler   *handlers.AdminHandler
}

func NewContainer() *Container {
	db := database.InitDB()

	userRepo := repositories.NewUserRepo(db)
	fileRepo := repositories.NewFileRepo(db)

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
		panic(err)
	}

	profileService := services.NewProfileService(userRepo, fileService)
	adminService := services.NewAdminService(userRepo)

	return &Container{
		AuthHandler:    handlers.NewAuthHandler(authService),
		FileHandler:    handlers.NewFileHandler(fileService),
		ProfileHandler: handlers.NewProfileHandler(profileService),
		ProfileV2:      handlersv2.NewProfileHandlerV2(profileService, fileService),
		AdminHandler:   handlers.NewAdminHandler(adminService),
	}
}
