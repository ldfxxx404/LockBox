package services

import (
	"back/internal/models"
	"back/internal/repositories"
	"errors"
)

type AdminService struct {
	UserRepo *repositories.UserRepo
}

func NewAdminService(repo *repositories.UserRepo) *AdminService {
	return &AdminService{UserRepo: repo}
}

func (s *AdminService) GetAllUsers() ([]models.User, error) {
	return s.UserRepo.GetAll()
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
