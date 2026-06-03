package services

import (
	"back/internal/models"
	"back/internal/storage"
	"context"
	"errors"
	"io"
	"strings"
	"testing"
)

type MockFileRepo struct {
	createFunc func(file *models.File) error
	getFunc    func(userID int) ([]models.File, error)
	deleteFunc func(userID int, filename string) error
	existsFunc func(userID int, filename string) (bool, error)
}

func (m *MockFileRepo) Create(file *models.File) error {
	if m.createFunc != nil {
		return m.createFunc(file)
	}
	return nil
}

func (m *MockFileRepo) GetFilesByUser(userID int) ([]models.File, error) {
	if m.getFunc != nil {
		return m.getFunc(userID)
	}
	return []models.File{}, nil
}

func (m *MockFileRepo) DeleteFile(userID int, filename string) error {
	if m.deleteFunc != nil {
		return m.deleteFunc(userID, filename)
	}
	return nil
}

func (m *MockFileRepo) Exists(userID int, filename string) (bool, error) {
	if m.existsFunc != nil {
		return m.existsFunc(userID, filename)
	}
	return false, nil
}

type MockUserRepo struct {
	getByIDFunc func(id int) (*models.User, error)
}

func (m *MockUserRepo) Create(user *models.User) error {
	return nil
}

func (m *MockUserRepo) GetByEmail(email string) (*models.User, error) {
	return nil, nil
}

func (m *MockUserRepo) GetByID(id int) (*models.User, error) {
	if m.getByIDFunc != nil {
		return m.getByIDFunc(id)
	}
	return &models.User{
		ID:           id,
		StorageLimit: 20,
	}, nil
}

func (m *MockUserRepo) GetAll() ([]models.User, error) {
	return []models.User{}, nil
}

func (m *MockUserRepo) GetAllAdmins() ([]models.User, error) {
	return []models.User{}, nil
}

func (m *MockUserRepo) UpdateStorageLimit(userID, newLimit int) error {
	return nil
}

func (m *MockUserRepo) UpdateAdmin(userID int, isAdmin bool) error {
	return nil
}

type MockStorageClient struct {
	putObjectFunc    func(ctx context.Context, bucket, objectName string, reader io.Reader, objectSize int64, contentType string) (string, error)
	getObjectFunc    func(ctx context.Context, bucket, objectName string) (io.ReadCloser, error)
	removeObjectFunc func(ctx context.Context, bucket, objectName string) error
	listObjectsFunc  func(ctx context.Context, bucket, prefix string, recursive bool) <-chan storage.ObjectInfo
	bucketExistsFunc func(ctx context.Context, bucket string) (bool, error)
	makeBucketFunc   func(ctx context.Context, bucket string) error
}

func (m *MockStorageClient) PutObject(ctx context.Context, bucket, objectName string, reader io.Reader, objectSize int64, contentType string) (string, error) {
	if m.putObjectFunc != nil {
		return m.putObjectFunc(ctx, bucket, objectName, reader, objectSize, contentType)
	}
	return "", nil
}

func (m *MockStorageClient) GetObject(ctx context.Context, bucket, objectName string) (io.ReadCloser, error) {
	if m.getObjectFunc != nil {
		return m.getObjectFunc(ctx, bucket, objectName)
	}
	return io.NopCloser(strings.NewReader("")), nil
}

func (m *MockStorageClient) RemoveObject(ctx context.Context, bucket, objectName string) error {
	if m.removeObjectFunc != nil {
		return m.removeObjectFunc(ctx, bucket, objectName)
	}
	return nil
}

func (m *MockStorageClient) ListObjects(ctx context.Context, bucket, prefix string, recursive bool) <-chan storage.ObjectInfo {
	if m.listObjectsFunc != nil {
		return m.listObjectsFunc(ctx, bucket, prefix, recursive)
	}
	ch := make(chan storage.ObjectInfo)
	close(ch)
	return ch
}

func (m *MockStorageClient) BucketExists(ctx context.Context, bucket string) (bool, error) {
	if m.bucketExistsFunc != nil {
		return m.bucketExistsFunc(ctx, bucket)
	}
	return true, nil
}

func (m *MockStorageClient) MakeBucket(ctx context.Context, bucket string) error {
	if m.makeBucketFunc != nil {
		return m.makeBucketFunc(ctx, bucket)
	}
	return nil
}

func TestListFiles(t *testing.T) {
	tests := []struct {
		name       string
		userID     int
		fileCount  int
		shouldFail bool
		errorMsg   string
	}{
		{
			name:      "user file list",
			userID:    1,
			fileCount: 3,
		},
		{
			name:      "user without files",
			userID:    2,
			fileCount: 0,
		},
		{
			name:      "large number of files",
			userID:    3,
			fileCount: 100,
		},
		{
			name:       "error while fetching files",
			userID:     99,
			shouldFail: true,
			errorMsg:   "database connection failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileRepo := &MockFileRepo{
				getFunc: func(userID int) ([]models.File, error) {
					if userID == 99 {
						return nil, errors.New("database connection failed")
					}
					files := make([]models.File, tt.fileCount)
					for i := 0; i < tt.fileCount; i++ {
						files[i] = models.File{
							ID:       i + 1,
							UserID:   userID,
							Filename: "file" + string(rune(48+i%10)) + ".txt",
						}
					}
					return files, nil
				},
			}

			fileService := &FileService{
				FileRepo: fileRepo,
				UserRepo: &MockUserRepo{},
				Storage:  &MockStorageClient{},
				Bucket:   "test-bucket",
			}

			files, err := fileService.ListFiles(tt.userID)

			if tt.shouldFail {
				if err == nil {
					t.Errorf("expected error: %s, but got success", tt.errorMsg)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if len(files) != tt.fileCount {
					t.Errorf("wrong file count: expected %d, got %d", tt.fileCount, len(files))
				}
			}
		})
	}
}

// TestGetFile checks file retrieval
func TestGetFile(t *testing.T) {
	tests := []struct {
		name       string
		userID     int
		filename   string
		data       string
		shouldFail bool
		errorMsg   string
	}{
		{
			name:     "successful file retrieval",
			userID:   1,
			filename: "document.txt",
			data:     "file content here",
		},
		{
			name:     "large file retrieval",
			userID:   1,
			filename: "large.bin",
			data:     strings.Repeat("x", 10000),
		},
		{
			name:       "file not found",
			userID:     1,
			filename:   "nonexistent.txt",
			shouldFail: true,
			errorMsg:   "not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storageClient := &MockStorageClient{
				getObjectFunc: func(ctx context.Context, bucket, objectName string) (io.ReadCloser, error) {
					if strings.Contains(objectName, "nonexistent") {
						return nil, errors.New("not found")
					}
					return io.NopCloser(strings.NewReader(tt.data)), nil
				},
			}

			fileService := &FileService{
				FileRepo: &MockFileRepo{},
				UserRepo: &MockUserRepo{},
				Storage:  storageClient,
				Bucket:   "test-bucket",
			}

			data, err := fileService.GetFile(tt.userID, tt.filename)

			if tt.shouldFail {
				if err == nil {
					t.Errorf("expected error: %s, but got success", tt.errorMsg)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if string(data) != tt.data {
					t.Errorf("wrong data: expected %q, got %q", tt.data, string(data))
				}
			}
		})
	}
}

// TestDeleteFile checks file deletion
func TestDeleteFile(t *testing.T) {
	tests := []struct {
		name       string
		userID     int
		filename   string
		shouldFail bool
		errorMsg   string
	}{
		{
			name:     "successful file deletion",
			userID:   1,
			filename: "document.txt",
		},
		{
			name:       "file not found in DB",
			userID:     1,
			filename:   "db-missing.txt",
			shouldFail: true,
			errorMsg:   "database error",
		},
		{
			name:       "file not found in storage",
			userID:     1,
			filename:   "storage-missing.txt",
			shouldFail: true,
			errorMsg:   "storage error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileRepo := &MockFileRepo{
				deleteFunc: func(userID int, filename string) error {
					if strings.Contains(filename, "db-missing") {
						return errors.New("database error")
					}
					return nil
				},
			}

			storageClient := &MockStorageClient{
				removeObjectFunc: func(ctx context.Context, bucket, objectName string) error {
					if strings.Contains(objectName, "storage-missing") {
						return errors.New("storage error")
					}
					return nil
				},
			}

			fileService := &FileService{
				FileRepo: fileRepo,
				UserRepo: &MockUserRepo{},
				Storage:  storageClient,
				Bucket:   "test-bucket",
			}

			err := fileService.DeleteFile(tt.userID, tt.filename)

			if tt.shouldFail {
				if err == nil {
					t.Errorf("expected error: %s, but got success", tt.errorMsg)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}
		})
	}
}

// TestGetStorageInfo checks storage information
func TestGetStorageInfo(t *testing.T) {
	tests := []struct {
		name         string
		userID       int
		storageLimit int
		objectCount  int
		totalSize    int64
		shouldFail   bool
		errorMsg     string
	}{
		{
			name:         "storage info retrieval",
			userID:       1,
			storageLimit: 20,
			objectCount:  3,
			totalSize:    5242880,
		},
		{
			name:         "empty storage",
			userID:       2,
			storageLimit: 20,
			objectCount:  0,
			totalSize:    0,
		},
		{
			name:       "user not found",
			userID:     99,
			shouldFail: true,
			errorMsg:   "user not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storageClient := &MockStorageClient{
				listObjectsFunc: func(ctx context.Context, bucket, prefix string, recursive bool) <-chan storage.ObjectInfo {
					ch := make(chan storage.ObjectInfo)
					go func() {
						defer close(ch)

						if tt.objectCount == 0 {
							return
						}

						sizePerObject := tt.totalSize / int64(tt.objectCount)
						for i := 0; i < tt.objectCount; i++ {
							ch <- storage.ObjectInfo{
								Name: "file" + string(rune(48+i)) + ".txt",
								Size: sizePerObject,
							}
						}
					}()
					return ch
				},
			}

			userRepo := &MockUserRepo{
				getByIDFunc: func(id int) (*models.User, error) {
					if id == 99 {
						return nil, errors.New("user not found")
					}
					return &models.User{
						ID:           id,
						StorageLimit: tt.storageLimit,
					}, nil
				},
			}

			fileService := &FileService{
				FileRepo: &MockFileRepo{},
				UserRepo: userRepo,
				Storage:  storageClient,
				Bucket:   "test-bucket",
			}

			used, limit, err := fileService.GetStorageInfo(tt.userID)

			if tt.shouldFail {
				if err == nil {
					t.Errorf("expected error: %s, but got success", tt.errorMsg)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if limit != tt.storageLimit {
					t.Errorf("wrong limit: expected %d, got %d", tt.storageLimit, limit)
				}
				if used < 0 {
					t.Errorf("used storage cannot be negative: %d", used)
				}
			}
		})
	}
}

// TestFileServiceInitialization checks service initialization
func TestFileServiceInitialization(t *testing.T) {
	fileService := &FileService{
		FileRepo: &MockFileRepo{},
		UserRepo: &MockUserRepo{},
		Storage:  &MockStorageClient{},
		Bucket:   "test-bucket",
	}

	if fileService.FileRepo == nil {
		t.Error("FileRepo must not be nil")
	}
	if fileService.UserRepo == nil {
		t.Error("UserRepo must not be nil")
	}
	if fileService.Storage == nil {
		t.Error("Storage must not be nil")
	}
	if fileService.Bucket != "test-bucket" {
		t.Errorf("wrong bucket: expected test-bucket, got %s", fileService.Bucket)
	}
}

// TestRepositoryIntegration checks repository integration
func TestRepositoryIntegration(t *testing.T) {
	t.Run("file repo is called on listing", func(t *testing.T) {
		called := false
		fileRepo := &MockFileRepo{
			getFunc: func(userID int) ([]models.File, error) {
				called = true
				return []models.File{}, nil
			},
		}

		fileService := &FileService{
			FileRepo: fileRepo,
			UserRepo: &MockUserRepo{},
			Storage:  &MockStorageClient{},
			Bucket:   "test-bucket",
		}

		_, _ = fileService.ListFiles(1)
		if !called {
			t.Error("FileRepo.GetFilesByUser must be called")
		}
	})

	t.Run("user repo is called on storage info", func(t *testing.T) {
		called := false
		userRepo := &MockUserRepo{
			getByIDFunc: func(id int) (*models.User, error) {
				called = true
				return &models.User{ID: id, StorageLimit: 20}, nil
			},
		}

		fileService := &FileService{
			FileRepo: &MockFileRepo{},
			UserRepo: userRepo,
			Storage:  &MockStorageClient{},
			Bucket:   "test-bucket",
		}

		_, _, _ = fileService.GetStorageInfo(1)
		if !called {
			t.Error("UserRepo.GetByID must be called")
		}
	})

	t.Run("file repo is called on delete", func(t *testing.T) {
		called := false
		fileRepo := &MockFileRepo{
			deleteFunc: func(userID int, filename string) error {
				called = true
				return nil
			},
		}

		fileService := &FileService{
			FileRepo: fileRepo,
			UserRepo: &MockUserRepo{},
			Storage:  &MockStorageClient{},
			Bucket:   "test-bucket",
		}

		_ = fileService.DeleteFile(1, "file.txt")
		if !called {
			t.Error("FileRepo.DeleteFile must be called")
		}
	})

	t.Run("storage client is called on file fetch", func(t *testing.T) {
		called := false
		storageClient := &MockStorageClient{
			getObjectFunc: func(ctx context.Context, bucket, objectName string) (io.ReadCloser, error) {
				called = true
				return io.NopCloser(strings.NewReader("data")), nil
			},
		}

		fileService := &FileService{
			FileRepo: &MockFileRepo{},
			UserRepo: &MockUserRepo{},
			Storage:  storageClient,
			Bucket:   "test-bucket",
		}

		_, _ = fileService.GetFile(1, "file.txt")
		if !called {
			t.Error("Storage.GetObject must be called")
		}
	})
}

// TestErrorHandling checks repository error handling
func TestErrorHandling(t *testing.T) {
	t.Run("error on file listing", func(t *testing.T) {
		fileRepo := &MockFileRepo{
			getFunc: func(userID int) ([]models.File, error) {
				return nil, errors.New("database error")
			},
		}

		fileService := &FileService{
			FileRepo: fileRepo,
			UserRepo: &MockUserRepo{},
			Storage:  &MockStorageClient{},
			Bucket:   "test-bucket",
		}

		_, err := fileService.ListFiles(1)
		if err == nil {
			t.Error("error expected")
		}
		if !strings.Contains(err.Error(), "database error") {
			t.Errorf("error should contain 'database error', got: %v", err)
		}
	})

	t.Run("error on storage info", func(t *testing.T) {
		userRepo := &MockUserRepo{
			getByIDFunc: func(id int) (*models.User, error) {
				return nil, errors.New("user not found")
			},
		}

		fileService := &FileService{
			FileRepo: &MockFileRepo{},
			UserRepo: userRepo,
			Storage:  &MockStorageClient{},
			Bucket:   "test-bucket",
		}

		_, _, err := fileService.GetStorageInfo(1)
		if err == nil {
			t.Error("error expected")
		}
	})

	t.Run("error on file delete", func(t *testing.T) {
		fileRepo := &MockFileRepo{
			deleteFunc: func(userID int, filename string) error {
				return errors.New("delete failed")
			},
		}

		fileService := &FileService{
			FileRepo: fileRepo,
			UserRepo: &MockUserRepo{},
			Storage:  &MockStorageClient{},
			Bucket:   "test-bucket",
		}

		err := fileService.DeleteFile(1, "file.txt")
		if err == nil {
			t.Error("error expected")
		}
	})

	t.Run("error on storage get object", func(t *testing.T) {
		storageClient := &MockStorageClient{
			getObjectFunc: func(ctx context.Context, bucket, objectName string) (io.ReadCloser, error) {
				return nil, errors.New("connection error")
			},
		}

		fileService := &FileService{
			FileRepo: &MockFileRepo{},
			UserRepo: &MockUserRepo{},
			Storage:  storageClient,
			Bucket:   "test-bucket",
		}

		_, err := fileService.GetFile(1, "file.txt")
		if err == nil {
			t.Error("error expected")
		}
	})
}

// BenchmarkListFiles benchmarks file listing
func BenchmarkListFiles(b *testing.B) {
	files := make([]models.File, 100)
	for i := 0; i < 100; i++ {
		files[i] = models.File{
			ID:       i,
			UserID:   1,
			Filename: "file" + string(rune(48+i%10)) + ".txt",
		}
	}

	fileService := &FileService{
		FileRepo: &MockFileRepo{
			getFunc: func(userID int) ([]models.File, error) {
				return files, nil
			},
		},
		UserRepo: &MockUserRepo{},
		Storage:  &MockStorageClient{},
		Bucket:   "test-bucket",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = fileService.ListFiles(1)
	}
}

// BenchmarkGetFile benchmarks file retrieval
func BenchmarkGetFile(b *testing.B) {
	fileService := &FileService{
		FileRepo: &MockFileRepo{},
		UserRepo: &MockUserRepo{},
		Storage: &MockStorageClient{
			getObjectFunc: func(ctx context.Context, bucket, objectName string) (io.ReadCloser, error) {
				return io.NopCloser(strings.NewReader("test data")), nil
			},
		},
		Bucket: "test-bucket",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = fileService.GetFile(1, "file.txt")
	}
}
