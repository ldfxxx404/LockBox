package repositories

import "back/internal/models"

type UserRepoInterface interface {
	Create(user *models.User) error
	GetByEmail(email string) (*models.User, error)
	GetByID(id int) (*models.User, error)
	GetAll() ([]models.User, error)
	GetAllAdmins() ([]models.User, error)
	UpdateStorageLimit(userID, newLimit int) error
	UpdateAdmin(userID int, isAdmin bool) error
}

type FileRepoInterface interface {
	Create(file *models.File) error
	GetFilesByUser(userID int) ([]models.File, error)
	DeleteFile(userID int, filename string) error
}
