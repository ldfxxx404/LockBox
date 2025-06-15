package repositories

import (
	"back/internal/models"

	"github.com/gofiber/fiber/v2/log"
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
	result, err := r.DB.Exec(CreateFileSql, file.UserID, file.Filename, file.OriginalName, file.Size, file.MimeType)
	log.Debug("create file", "result", result)
	if err != nil {
		log.Error("create file error", "err", err)
		return err
	}
	return nil
}

func (r *FileRepo) GetFilesByUser(userID int) ([]models.File, error) {
	var files []models.File
	err := r.DB.Select(&files, GetFileSql, userID)
	log.Debug("get files", "files", files)
	if err != nil {
		log.Error("get files by users", "err", err)
		return nil, err
	}
	return files, nil
}

func (r *FileRepo) DeleteFile(userID int, filename string) error {
	result, err := r.DB.Exec(DeleteFileSql, userID, filename)
	log.Debug("delete file", "result", result)
	if err != nil {
		log.Error("delete file error", "err", err)
		return err
	}
	return nil
}
