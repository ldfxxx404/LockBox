package repositories

import (
	"back/internal/models"
	"github.com/jmoiron/sqlx"
)

type FileRepo struct {
	DB *sqlx.DB
}

func NewFileRepo(db *sqlx.DB) *FileRepo {
	return &FileRepo{DB: db}
}

func (r *FileRepo) Create(file *models.File) error {
	query := `INSERT INTO files (user_id, filename, original_name, size, mime_type) VALUES (:user_id, :filename, :original_name, :size, :mime_type)`
	_, err := r.DB.NamedExec(query, file)
	return err
}

func (r *FileRepo) GetFilesByUser(userID int) ([]models.File, error) {
	var files []models.File
	query := `SELECT * FROM files WHERE user_id = ?`
	err := r.DB.Select(&files, query, userID)
	return files, err
}

func (r *FileRepo) DeleteFile(userID int, filename string) error {
	query := `DELETE FROM files WHERE user_id = ? AND filename = ?`
	_, err := r.DB.Exec(query, userID, filename)
	return err
}
