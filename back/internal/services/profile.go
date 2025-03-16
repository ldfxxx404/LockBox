package services

import (
	"back/internal/models"
	"back/internal/repositories"
)

type ProfileService struct {
	UserRepo *repositories.UserRepo
	FileServ *FileService
}

func NewProfileService(userRepo *repositories.UserRepo, fileServ *FileService) *ProfileService {
	return &ProfileService{UserRepo: userRepo, FileServ: fileServ}
}

func (s *ProfileService) GetProfile(userID int) (*models.User, int64, error) {
	user, err := s.UserRepo.GetByID(userID)
	if err != nil {
		return nil, 0, err
	}
	used, err := s.FileServ.GetStorageInfo(userID)
	if err != nil {
		return nil, 0, err
	}
	return user, used, nil
}
