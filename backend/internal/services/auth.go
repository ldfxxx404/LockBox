package services

import (
	"back/config"
	"back/internal/models"
	"back/internal/repositories"
	"back/internal/utils"
	"errors"
)

type AuthService struct {
	UserRepo *repositories.UserRepo
}

func NewAuthService(repo *repositories.UserRepo) *AuthService {
	return &AuthService{UserRepo: repo}
}

func (s *AuthService) Register(dto models.RegisterDTO) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		Email:        dto.Email,
		Name:         dto.Name,
		Password:     hashedPassword,
		IsAdmin:      false,
		StorageLimit: config.StorageLimit,
	}
	err = s.UserRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AuthService) Login(dto models.LoginDTO) (string, *models.User, error) {
	user, err := s.UserRepo.GetByEmail(dto.Email)
	if err != nil {
		return "", nil, errors.New("invalid email or password")
	}
	if !utils.CheckPasswordHash(dto.Password, user.Password) {
		return "", nil, errors.New("invalid email or password")
	}
	token, err := utils.GenerateToken(user.ID, user.IsAdmin)
	if err != nil {
		return "", nil, err
	}
	return token, user, nil
}

func (s *AuthService) Logout(token string) error {
	return nil
}
