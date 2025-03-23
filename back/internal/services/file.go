package services

import (
	"back/config"
	"back/internal/models"
	"back/internal/repositories"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type FileService struct {
	FileRepo *repositories.FileRepo
}

func NewFileService(repo *repositories.FileRepo) *FileService {
	return &FileService{FileRepo: repo}
}

func (s *FileService) UploadFile(userID int, fileHeader *multipart.FileHeader) error {
	userDir := fmt.Sprintf("%s/%d", config.StorageDir, userID)
	if err := os.MkdirAll(userDir, os.ModePerm); err != nil {
		return errors.New("failed to create user directory")
	}

	src, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dstPath := filepath.Join(userDir, fileHeader.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	size, err := io.Copy(dst, src)
	if err != nil {
		return err
	}

	file := &models.File{
		UserID:       userID,
		Filename:     fileHeader.Filename,
		OriginalName: fileHeader.Filename,
		Size:         size,
		MimeType:     fileHeader.Header.Get("Content-Type"),
	}
	return s.FileRepo.Create(file)
}

func (s *FileService) ListFiles(userID int) ([]models.File, error) {
	return s.FileRepo.GetFilesByUser(userID)
}

func (s *FileService) GetFilePath(userID int, filename string) (string, error) {
	filePath := fmt.Sprintf("%s/%d/%s", config.StorageDir, userID, filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", errors.New("file not found")
	}
	return filePath, nil
}

func (s *FileService) DeleteFile(userID int, filename string) error {
	filePath := fmt.Sprintf("%s/%d/%s", config.StorageDir, userID, filename)
	if err := os.Remove(filePath); err != nil {
		return errors.New("failed to delete file")
	}
	return s.FileRepo.DeleteFile(userID, filename)
}

func (s *FileService) GetStorageInfo(userID int) (usedMB int64, limitMB int, err error) {
	userDir := fmt.Sprintf("%s/%d", config.StorageDir, userID)
	var totalSize int64
	err = filepath.Walk(userDir, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			totalSize += info.Size()
		}
		return nil
	})
	if err != nil {
		return 0, 0, err
	}
	usedMB = totalSize / (1024 * 1024)
	limitMB = config.StorageLimit
	return usedMB, limitMB, nil
}
