package services

import (
	"back/config"
	"back/internal/models"
	"back/internal/repositories"
	"back/internal/utils"
	"errors"

	"github.com/gofiber/fiber/v2/log"
)

type AuthService struct {
	UserRepo repositories.UserRepoInterface
}

func NewAuthService(repo repositories.UserRepoInterface) *AuthService {
	return &AuthService{UserRepo: repo}
}

func (s *AuthService) Register(dto models.RegisterDTO) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		log.Error("AuthService: failed to hash password", "err", err)
		return nil, err
	}
	log.Debug("AuthService: password hashed successfully")

	user := &models.User{
		Email:        dto.Email,
		Name:         dto.Name,
		Password:     hashedPassword,
		IsAdmin:      false,
		StorageLimit: config.StorageLimit,
	}

	log.Debug("AuthService: creating user", "email", user.Email, "name", user.Name)

	err = s.UserRepo.Create(user)
	if err != nil {
		log.Error("AuthService: failed to create user", "email", user.Email, "err", err)
		return nil, err
	}

	log.Info("AuthService: user registered successfully", "user_id", user.ID, "email", user.Email)
	return user, nil
}

func (s *AuthService) Login(dto models.LoginDTO) (string, *models.User, error) {
	user, err := s.UserRepo.GetByEmail(dto.Email)
	if err != nil {
		log.Warn("AuthService: login failed - user not found", "email", dto.Email, "err", err)
		return "", nil, errors.New("invalid email or password")
	}

	if !utils.CheckPasswordHash(dto.Password, user.Password) {
		log.Warn("AuthService: login failed - invalid password", "email", dto.Email)
		return "", nil, errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(user.ID, user.IsAdmin)
	if err != nil {
		log.Error("AuthService: failed to generate token", "user_id", user.ID, "err", err)
		return "", nil, err
	}

	log.Info("AuthService: login successful", "user_id", user.ID, "email", user.Email)
	return token, user, nil
}

func (s *AuthService) Logout(token string) error {
	log.Info("AuthService: logout requested")
	return nil
}
