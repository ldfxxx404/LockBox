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
		log.Error("minio openssl connect err", "err", err)
		return nil, err
	}

	ctx := context.Background()
	exists, err := minioClient.BucketExists(ctx, bucket)
	if err != nil {
		log.Error("bucket exist check error", "err", err)
		return nil, err
	}
	if !exists {
		err = minioClient.MakeBucket(ctx, bucket, minio.MakeBucketOptions{})
		if err != nil {
			log.Error("make bucket error", "err", err)
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
	file, err := fileHeader.Open()
	if err != nil {
		log.Error("open file header", "err", err)
		return err
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Error("file close error", "err", err)
		}
	}()

	objectName := fmt.Sprintf("%d/%s", userID, fileHeader.Filename)
	size := fileHeader.Size
	contentType := fileHeader.Header.Get("Content-Type")

	uploadInfo, err := s.Minio.PutObject(context.Background(),
		s.Bucket,
		objectName,
		file,
		size,
		minio.PutObjectOptions{
			ContentType: contentType,
		})
	log.Debug("put object minio:", "upload file info", uploadInfo)
	if err != nil {
		log.Error("put object error", "err", err)
		return err
	}

	meta := &models.File{
		UserID:       userID,
		Filename:     fileHeader.Filename,
		OriginalName: fileHeader.Filename,
		Size:         size,
		MimeType:     contentType,
	}
	return s.FileRepo.Create(meta)
}

func (s *FileService) ListFiles(userID int) ([]models.File, error) {
	return s.FileRepo.GetFilesByUser(userID)
}

func (s *FileService) GetFile(userID int, filename string) ([]byte, error) {
	objectName := fmt.Sprintf("%d/%s", userID, filename)

	obj, err := s.Minio.GetObject(context.Background(),
		s.Bucket,
		objectName,
		minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("get object error: %w", err)
	}

	defer func() { _ = obj.Close() }()
	data, err := io.ReadAll(obj)
	if err != nil {
		return nil, fmt.Errorf("failed to read object: %w", err)
	}

	return data, nil
}

func (s *FileService) DeleteFile(userID int, filename string) error {
	objectName := fmt.Sprintf("%d/%s", userID, filename)
	err := s.Minio.RemoveObject(context.Background(),
		s.Bucket,
		objectName,
		minio.RemoveObjectOptions{})
	if err != nil {
		log.Error("error of delete file", "err", err)
		return errors.New("failed to delete file")
	}
	return s.FileRepo.DeleteFile(userID, filename)
}

func (s *FileService) GetStorageInfo(userID int) (usedMB int64, limitMB int, err error) {
	prefix := fmt.Sprintf("%d/", userID)
	opts := minio.ListObjectsOptions{Prefix: prefix, Recursive: true}

	var totalSize int64
	for object := range s.Minio.ListObjects(context.Background(), s.Bucket, opts) {
		if object.Err != nil {
			return 0, 0, object.Err
		}
		totalSize += object.Size
	}
	usedMB = totalSize / (1024 * 1024)

	user, err := s.UserRepo.GetByID(userID)
	if err != nil {
		log.Error("error of use repo", "err", err)
		return 0, 0, err
	}
	limitMB = user.StorageLimit
	return usedMB, limitMB, nil
}
