package services

import (
	"back/internal/models"
	"back/internal/repositories"

	log "github.com/charmbracelet/log"
)

type ProfileService struct {
	UserRepo *repositories.UserRepo
	FileServ *FileService
}

func NewProfileService(userRepo *repositories.UserRepo, fileServ *FileService) *ProfileService {
	return &ProfileService{UserRepo: userRepo, FileServ: fileServ}
}

func (s *ProfileService) GetProfile(userID int) (*models.User, int64, int, error) {
	user, err := s.UserRepo.GetByID(userID)
	if err != nil {
		log.Error(err)
		return nil, 0, 0, err
	}
	usedMB, limitMB, err := s.FileServ.GetStorageInfo(userID)
	if err != nil {
		log.Error(err)
		return nil, 0, 0, err
	}
	return user, usedMB, limitMB, nil
}
