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
	log.Debug("hashed passwod of latest user", "Password hash", hashedPassword)
	if err != nil {
		log.Error("hash password check", "err", err)
		return nil, err
	}
	user := &models.User{
		Email:        dto.Email,
		Name:         dto.Name,
		Password:     hashedPassword,
		IsAdmin:      false,
		StorageLimit: config.StorageLimit,
	}
	log.Debug("current user info", "userModel", user)
	err = s.UserRepo.Create(user)
	if err != nil {
		log.Error("error of create user", "err", err)
		return nil, err
	}
	return user, nil
}

func (s *AuthService) Login(dto models.LoginDTO) (string, *models.User, error) {
	user, err := s.UserRepo.GetByEmail(dto.Email)
	if err != nil {
		log.Error("get user by email error", "err", err)
		return "", nil, errors.New("invalid email or password")
	}
	if !utils.CheckPasswordHash(dto.Password, user.Password) {
		return "", nil, errors.New("invalid email or password")
	}
	token, err := utils.GenerateToken(user.ID, user.IsAdmin)
	if err != nil {
		log.Error("generate token error", "err", err)
		return "", nil, err
	}

	log.Debug("new logn info", "new login user", user, "user JWT token", token)
	return token, user, nil
}

func (s *AuthService) Logout(token string) error {
	return nil
}
