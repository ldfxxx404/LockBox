package services

import (
	"back/internal/models"
	"back/internal/repositories"

	"github.com/gofiber/fiber/v2/log"
)

type ProfileService struct {
	UserRepo repositories.UserRepoInterface
	FileServ *FileService
}

func NewProfileService(userRepo repositories.UserRepoInterface, fileServ *FileService) *ProfileService {
	return &ProfileService{UserRepo: userRepo, FileServ: fileServ}
}

func (s *ProfileService) GetProfile(userID int) (*models.User, int64, int, error) {
	user, err := s.UserRepo.GetByID(userID)
	if err != nil {
		log.Error("use repo error", "err", err)
		return nil, 0, 0, err
	}
	usedMB, limitMB, err := s.FileServ.GetStorageInfo(userID)
	if err != nil {
		log.Error("use get storage", "err", err)
		return nil, 0, 0, err
	}
	log.Debug("User profile info", "Used storage MB", usedMB, "User info", user, "Limit MB", limitMB)
	return user, usedMB, limitMB, nil
}
