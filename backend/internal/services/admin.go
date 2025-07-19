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
	users, err := s.UserRepo.GetAll()
	if err != nil {
		log.Error("AdminService: failed to get all users", "err", err)
		return nil, err
	}
	log.Info("AdminService: successfully fetched all users", "count", len(users))
	return users, nil
}

func (s *AdminService) GetAllAdminUsers() ([]models.User, error) {
	admins, err := s.UserRepo.GetAllAdmins()
	if err != nil {
		log.Error("AdminService: faild to get all admin users", "err", err)
	}
	return admins, nil
}

func (s *AdminService) UpdateStorageLimit(userID, newLimit int) error {
	if newLimit <= 0 {
		err := errors.New("invalid storage limit: must be greater than zero")
		log.Error("AdminService: update storage limit failed", "user_id", userID, "new_limit", newLimit, "err", err)
		return err
	}

	err := s.UserRepo.UpdateStorageLimit(userID, newLimit)
	if err != nil {
		log.Error("AdminService: failed to update storage limit", "user_id", userID, "new_limit", newLimit, "err", err)
		return err
	}

	log.Info("AdminService: updated storage limit", "user_id", userID, "new_limit", newLimit)
	return nil
}

func (s *AdminService) MakeAdmin(userID int) error {
	err := s.UserRepo.UpdateAdmin(userID, true)
	if err != nil {
		log.Error("AdminService: failed to make user admin", "user_id", userID, "err", err)
		return err
	}

	log.Info("AdminService: user promoted to admin", "user_id", userID)
	return nil
}

func (s *AdminService) RevokeAdmin(userID int) error {
	err := s.UserRepo.UpdateAdmin(userID, false)
	if err != nil {
		log.Error("AdminService: failed to revoke admin rights", "user_id", userID, "err", err)
		return err
	}

	log.Info("AdminService: admin rights revoked", "user_id", userID)
	return nil
}
