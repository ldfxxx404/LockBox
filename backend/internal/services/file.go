package services

import (
	"back/internal/models"
	"back/internal/repositories"
	"back/internal/storage"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2/log"
)

type FileService struct {
	FileRepo repositories.FileRepoInterface
	UserRepo repositories.UserRepoInterface
	Storage  storage.StorageClient
	Bucket   string
}

func NewFileService(
	fileRepo repositories.FileRepoInterface,
	userRepo repositories.UserRepoInterface,
	endpoint, accessKey, secretKey, bucket string,
	useSSL bool,
) (*FileService, error) {
	storageClient, err := storage.NewMinIOAdapter(endpoint, accessKey, secretKey, useSSL)
	if err != nil {
		log.Error("storage client initialization error:", err)
		return nil, err
	}

	ctx := context.Background()
	exists, err := storageClient.BucketExists(ctx, bucket)
	if err != nil {
		log.Error("bucket exist check error:", err)
		return nil, err
	}
	if !exists {
		err = storageClient.MakeBucket(ctx, bucket)
		if err != nil {
			log.Error("make bucket error:", err)
			return nil, err
		}
	}

	return &FileService{
		FileRepo: fileRepo,
		UserRepo: userRepo,
		Storage:  storageClient,
		Bucket:   bucket,
	}, nil
}

func (s *FileService) incrementNewName(fileHeader *multipart.FileHeader, userID int) (string, string) {
	base := strings.TrimSuffix(fileHeader.Filename, filepath.Ext(fileHeader.Filename))
	ext := filepath.Ext(fileHeader.Filename)

	newName := fileHeader.Filename
	counter := 1
	for {
		exists, _ := s.FileRepo.Exists(userID, newName)
		if !exists {
			break
		}
		newName = fmt.Sprintf("%s(%d)%s", base, counter, ext)
		counter++
	}
	return fmt.Sprintf("%d/%s", userID, newName), newName
}

func (s *FileService) UploadFile(userID int, fileHeader *multipart.FileHeader) error {
	log.Info("Starting file upload:", "user_id", userID, "filename", fileHeader.Filename, "size", fileHeader.Size)

	file, err := fileHeader.Open()
	if err != nil {
		log.Error("Failed to open file header:", err)
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Error("Failed to close file:", err)
		}
	}()

	objectName, newName := s.incrementNewName(fileHeader, userID)
	contentType := fileHeader.Header.Get("Content-Type")

	etag, err := s.Storage.PutObject(context.Background(),
		s.Bucket,
		objectName,
		file,
		fileHeader.Size,
		contentType)
	if err != nil {
		log.Error("Storage PutObject error:", err)
		return err
	}
	log.Info("File uploaded to storage:", "etag", etag)

	meta := &models.File{
		UserID:       userID,
		Filename:     newName,
		OriginalName: fileHeader.Filename,
		Size:         fileHeader.Size,
		MimeType:     contentType,
	}
	if err := s.FileRepo.Create(meta); err != nil {
		log.Error("Failed to save file metadata to DB:", err)
		return err
	}
	log.Info("File metadata saved:", "user_id", userID, "filename", fileHeader.Filename)

	return nil
}

func (s *FileService) ListFiles(userID int) ([]models.File, error) {
	log.Info("Listing files for user_id:", userID)
	files, err := s.FileRepo.GetFilesByUser(userID)
	if err != nil {
		log.Error("Failed to list files:", err)
		return nil, err
	}
	log.Info("Found files:", "count", len(files), "user_id", userID)
	return files, nil
}

func (s *FileService) GetFile(userID int, filename string) ([]byte, error) {
	objectName := fmt.Sprintf("%d/%s", userID, filename)
	log.Info("Fetching file:", "user_id", userID, "filename", filename)

	obj, err := s.Storage.GetObject(context.Background(), s.Bucket, objectName)
	if err != nil {
		log.Error("Failed to get object from storage:", err)
		return nil, fmt.Errorf("get object error: %w", err)
	}
	defer func() { _ = obj.Close() }()

	data, err := io.ReadAll(obj)
	if err != nil {
		log.Error("Failed to read object data:", err)
		return nil, fmt.Errorf("failed to read object: %w", err)
	}
	log.Info("File read successfully:", "user_id", userID, "filename", filename, "size", len(data))
	return data, nil
}

func (s *FileService) DeleteFile(userID int, filename string) error {
	objectName := fmt.Sprintf("%d/%s", userID, filename)
	log.Info("Deleting file:", "user_id", userID, "filename", filename)

	if err := s.Storage.RemoveObject(context.Background(), s.Bucket, objectName); err != nil {
		log.Error("Failed to delete file from storage:", err)
		return errors.New("failed to delete file")
	}
	if err := s.FileRepo.DeleteFile(userID, filename); err != nil {
		log.Error("Failed to delete file from DB:", err)
		return err
	}
	log.Info("File deleted successfully:", "user_id", userID, "filename", filename)
	return nil
}

func (s *FileService) GetStorageInfo(userID int) (usedMB int64, limitMB int, err error) {
	log.Info("Fetching storage info for user_id:", userID)

	prefix := fmt.Sprintf("%d/", userID)
	var totalSize int64
	for objInfo := range s.Storage.ListObjects(context.Background(), s.Bucket, prefix, true) {
		if objInfo.Err != nil {
			log.Error("Failed to list object:", objInfo.Err)
			return 0, 0, objInfo.Err
		}
		totalSize += objInfo.Size
	}
	usedMB = totalSize / (1024 * 1024)
	log.Info("Used storage:", "user_id", userID, "usedMB", usedMB)

	user, err := s.UserRepo.GetByID(userID)
	if err != nil {
		log.Error("Failed to fetch user data:", err)
		return 0, 0, err
	}
	limitMB = user.StorageLimit
	log.Info("Storage info:", "user_id", userID, "limitMB", limitMB)

	return usedMB, limitMB, nil
}
