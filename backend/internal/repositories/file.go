package repositories

import (
	"back/internal/models"
	"github.com/jmoiron/sqlx"
)

const (
	CreateFileSql = `INSERT INTO files (user_id, filename, original_name, size, mime_type) VALUES ($1, $2, $3, $4, $5)`
	GetFileSql    = `SELECT * FROM files WHERE user_id = $1`
	DeleteFileSql = `DELETE FROM files WHERE user_id = $1 AND filename = $2`
)

type FileRepo struct {
	DB *sqlx.DB
}

func NewFileRepo(db *sqlx.DB) *FileRepo {
	return &FileRepo{DB: db}
}

func (r *FileRepo) Create(file *models.File) error {
	_, err := r.DB.Exec(CreateFileSql, file.UserID, file.Filename, file.OriginalName, file.Size, file.MimeType)
	return err
}

func (r *FileRepo) GetFilesByUser(userID int) ([]models.File, error) {
	var files []models.File
	err := r.DB.Select(&files, GetFileSql, userID)
	return files, err
}

func (r *FileRepo) DeleteFile(userID int, filename string) error {
	_, err := r.DB.Exec(DeleteFileSql, userID, filename)
	return err
}
