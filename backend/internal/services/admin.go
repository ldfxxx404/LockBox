package services

import (
	"back/internal/models"
	"back/internal/repositories"
	"errors"

	"github.com/gofiber/fiber/v2/log"
)

type AdminService struct {
	UserRepo repositories.UserRepoInterface
}

func NewAdminService(repo repositories.UserRepoInterface) *AdminService {
	return &AdminService{UserRepo: repo}
}

func (s *AdminService) GetAllUsers() ([]models.User, error) {
	model, err := s.UserRepo.GetAll()
	if err != nil {
		log.Debug(err)
		log.Error("get all users fetch", "err", err)
		return nil, err
	}
	return model, nil
}

func (s *AdminService) UpdateStorageLimit(userID, newLimit int) error {
	if newLimit <= 0 {
		return errors.New("invalid storage limit")
	}
	return s.UserRepo.UpdateStorageLimit(userID, newLimit)
}

func (s *AdminService) MakeAdmin(userID int) error {
	return s.UserRepo.UpdateAdmin(userID, true)
}

func (s *AdminService) RevokeAdmin(userID int) error {
	return s.UserRepo.UpdateAdmin(userID, false)
}
