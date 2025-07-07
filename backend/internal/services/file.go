package services

import (
	"back/internal/models"
	"back/internal/repositories"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/gofiber/fiber/v2/log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type FileService struct {
	FileRepo repositories.FileRepoInterface
	UserRepo repositories.UserRepoInterface
	Minio    *minio.Client
	Bucket   string
}

func NewFileService(
	fileRepo repositories.FileRepoInterface,
	userRepo repositories.UserRepoInterface,
	endpoint, accessKey, secretKey, bucket string,
	useSSL bool,
) (*FileService, error) {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Error("minio openssl connect err", err)
		return nil, err
	}

	ctx := context.Background()
	exists, err := minioClient.BucketExists(ctx, bucket)
	if err != nil {
		log.Error("bucket exist check error", err)
		return nil, err
	}
	if !exists {
		err = minioClient.MakeBucket(ctx, bucket, minio.MakeBucketOptions{})
		if err != nil {
			log.Error("make bucket error", err)
			return nil, err
		}
	}

	return &FileService{
		FileRepo: fileRepo,
		UserRepo: userRepo,
		Minio:    minioClient,
		Bucket:   bucket,
	}, nil
}

func (s *FileService) UploadFile(userID int, fileHeader *multipart.FileHeader) error {
	log.Infof("Starting file upload: user_id=%d, filename=%s, size=%d", userID, fileHeader.Filename, fileHeader.Size)

	file, err := fileHeader.Open()
	if err != nil {
		log.Error("Failed to open file header:", err)
		return err
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			log.Error("Failed to close file:", cerr)
		}
	}()

	objectName := fmt.Sprintf("%d/%s", userID, fileHeader.Filename)
	contentType := fileHeader.Header.Get("Content-Type")

	uploadInfo, err := s.Minio.PutObject(context.Background(),
		s.Bucket,
		objectName,
		file,
		fileHeader.Size,
		minio.PutObjectOptions{
			ContentType: contentType,
		})
	if err != nil {
		log.Error("MinIO PutObject error:", err)
		return err
	}
	log.Infof("File uploaded to MinIO: location=%s, etag=%s", uploadInfo.Location, uploadInfo.ETag)

	meta := &models.File{
		UserID:       userID,
		Filename:     fileHeader.Filename,
		OriginalName: fileHeader.Filename,
		Size:         fileHeader.Size,
		MimeType:     contentType,
	}
	if err := s.FileRepo.Create(meta); err != nil {
		log.Error("Failed to save file metadata to DB:", err)
		return err
	}
	log.Infof("File metadata saved: user_id=%d, filename=%s", userID, fileHeader.Filename)

	return nil
}

func (s *FileService) ListFiles(userID int) ([]models.File, error) {
	log.Infof("Listing files for user_id=%d", userID)
	files, err := s.FileRepo.GetFilesByUser(userID)
	if err != nil {
		log.Error("Failed to list files:", err)
		return nil, err
	}
	log.Infof("Found %d files for user_id=%d", len(files), userID)
	return files, nil
}

func (s *FileService) GetFile(userID int, filename string) ([]byte, error) {
	objectName := fmt.Sprintf("%d/%s", userID, filename)
	log.Infof("Fetching file: user_id=%d, filename=%s", userID, filename)

	obj, err := s.Minio.GetObject(context.Background(),
		s.Bucket,
		objectName,
		minio.GetObjectOptions{})
	if err != nil {
		log.Error("Failed to get object from MinIO:", err)
		return nil, fmt.Errorf("get object error: %w", err)
	}
	defer func() { _ = obj.Close() }()

	data, err := io.ReadAll(obj)
	if err != nil {
		log.Error("Failed to read object data:", err)
		return nil, fmt.Errorf("failed to read object: %w", err)
	}
	log.Infof("File read successfully: user_id=%d, filename=%s, size=%d", userID, filename, len(data))
	return data, nil
}

func (s *FileService) DeleteFile(userID int, filename string) error {
	objectName := fmt.Sprintf("%d/%s", userID, filename)
	log.Infof("Deleting file: user_id=%d, filename=%s", userID, filename)

	if err := s.Minio.RemoveObject(context.Background(),
		s.Bucket,
		objectName,
		minio.RemoveObjectOptions{}); err != nil {
		log.Error("Failed to delete file from MinIO:", err)
		return errors.New("failed to delete file")
	}
	if err := s.FileRepo.DeleteFile(userID, filename); err != nil {
		log.Error("Failed to delete file from DB:", err)
		return err
	}
	log.Infof("File deleted successfully: user_id=%d, filename=%s", userID, filename)
	return nil
}

func (s *FileService) GetStorageInfo(userID int) (usedMB int64, limitMB int, err error) {
	log.Infof("Fetching storage info for user_id=%d", userID)

	prefix := fmt.Sprintf("%d/", userID)
	opts := minio.ListObjectsOptions{Prefix: prefix, Recursive: true}

	var totalSize int64
	for object := range s.Minio.ListObjects(context.Background(), s.Bucket, opts) {
		if object.Err != nil {
			log.Error("Failed to list object:", object.Err)
			return 0, 0, object.Err
		}
		totalSize += object.Size
	}
	usedMB = totalSize / (1024 * 1024)
	log.Infof("Used storage: user_id=%d, usedMB=%d", userID, usedMB)

	user, err := s.UserRepo.GetByID(userID)
	if err != nil {
		log.Error("Failed to fetch user data:", err)
		return 0, 0, err
	}
	limitMB = user.StorageLimit
	log.Infof("Storage info: user_id=%d, limitMB=%d", userID, limitMB)

	return usedMB, limitMB, nil
}
