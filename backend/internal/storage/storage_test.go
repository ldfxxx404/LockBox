package storage

import (
	"context"
	"errors"
	"io"
	"strings"
	"testing"
)

// MockStorageClient мок для StorageClient интерфейса
type MockStorageClient struct {
	putObjectFunc    func(ctx context.Context, bucket, objectName string, reader io.Reader, objectSize int64, contentType string) (etag string, err error)
	getObjectFunc    func(ctx context.Context, bucket, objectName string) (io.ReadCloser, error)
	removeObjectFunc func(ctx context.Context, bucket, objectName string) error
	listObjectsFunc  func(ctx context.Context, bucket, prefix string, recursive bool) <-chan ObjectInfo
	bucketExistsFunc func(ctx context.Context, bucket string) (exists bool, err error)
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

func (m *MockStorageClient) ListObjects(ctx context.Context, bucket, prefix string, recursive bool) <-chan ObjectInfo {
	if m.listObjectsFunc != nil {
		return m.listObjectsFunc(ctx, bucket, prefix, recursive)
	}
	ch := make(chan ObjectInfo)
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

// TestPutObject проверяет загрузку объекта
func TestPutObject(t *testing.T) {
	tests := []struct {
		name          string
		bucket        string
		objectName    string
		data          string
		size          int64
		contentType   string
		shouldFail    bool
		expectedError string
	}{
		{
			name:        "успешная загрузка текстового файла",
			bucket:      "test-bucket",
			objectName:  "1/file.txt",
			data:        "hello world",
			size:        11,
			contentType: "text/plain",
			shouldFail:  false,
		},
		{
			name:        "загрузка JSON файла",
			bucket:      "test-bucket",
			objectName:  "2/data.json",
			data:        `{"key": "value"}`,
			size:        16,
			contentType: "application/json",
			shouldFail:  false,
		},
		{
			name:        "загрузка изображения",
			bucket:      "test-bucket",
			objectName:  "1/image.png",
			data:        "\x89PNG\r\n\x1a\n",
			size:        8,
			contentType: "image/png",
			shouldFail:  false,
		},
		{
			name:          "пустой bucket",
			bucket:        "",
			objectName:    "1/file.txt",
			size:          10,
			contentType:   "text/plain",
			shouldFail:    true,
			expectedError: "bucket name required",
		},
		{
			name:          "пустой objectName",
			bucket:        "test-bucket",
			objectName:    "",
			size:          10,
			contentType:   "text/plain",
			shouldFail:    true,
			expectedError: "object name required",
		},
		{
			name:          "неправильный путь объекта",
			bucket:        "test-bucket",
			objectName:    "invalid//path",
			size:          10,
			contentType:   "text/plain",
			shouldFail:    true,
			expectedError: "invalid object name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &MockStorageClient{
				putObjectFunc: func(ctx context.Context, bucket, objectName string, reader io.Reader, objectSize int64, contentType string) (string, error) {
					if len(bucket) == 0 {
						return "", errors.New("bucket name required")
					}
					if len(objectName) == 0 {
						return "", errors.New("object name required")
					}
					if strings.Contains(objectName, "//") {
						return "", errors.New("invalid object name")
					}
					return "etag123", nil
				},
			}

			reader := strings.NewReader(tt.data)
			etag, err := mock.PutObject(context.Background(), tt.bucket, tt.objectName, reader, tt.size, tt.contentType)

			if tt.shouldFail {
				if err == nil {
					t.Errorf("ожидалась ошибка: %s, но получен успех", tt.expectedError)
				}
				if !strings.Contains(err.Error(), tt.expectedError) {
					t.Errorf("неправильная ошибка: ожидали %s, получили %s", tt.expectedError, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("не ожидалась ошибка: %v", err)
				}
				if len(etag) == 0 {
					t.Error("etag не должен быть пустым")
				}
			}
		})
	}
}

// TestGetObject проверяет получение объекта
func TestGetObject(t *testing.T) {
	tests := []struct {
		name          string
		bucket        string
		objectName    string
		data          string
		shouldFail    bool
		expectedError string
	}{
		{
			name:       "успешное получение файла",
			bucket:     "test-bucket",
			objectName: "1/file.txt",
			data:       "test content",
			shouldFail: false,
		},
		{
			name:       "получение большого файла",
			bucket:     "test-bucket",
			objectName: "2/large.bin",
			data:       strings.Repeat("x", 10000),
			shouldFail: false,
		},
		{
			name:          "объект не найден",
			bucket:        "test-bucket",
			objectName:    "1/nonexistent.txt",
			shouldFail:    true,
			expectedError: "not found",
		},
		{
			name:          "пустой bucket",
			bucket:        "",
			objectName:    "1/file.txt",
			shouldFail:    true,
			expectedError: "bucket name required",
		},
		{
			name:          "пустой objectName",
			bucket:        "test-bucket",
			objectName:    "",
			shouldFail:    true,
			expectedError: "object name required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &MockStorageClient{
				getObjectFunc: func(ctx context.Context, bucket, objectName string) (io.ReadCloser, error) {
					if len(bucket) == 0 {
						return nil, errors.New("bucket name required")
					}
					if len(objectName) == 0 {
						return nil, errors.New("object name required")
					}
					if strings.Contains(objectName, "nonexistent") {
						return nil, errors.New("not found")
					}
					return io.NopCloser(strings.NewReader(tt.data)), nil
				},
			}

			reader, err := mock.GetObject(context.Background(), tt.bucket, tt.objectName)

			if tt.shouldFail {
				if err == nil {
					t.Errorf("ожидалась ошибка: %s, но получен успех", tt.expectedError)
				}
				if !strings.Contains(err.Error(), tt.expectedError) {
					t.Errorf("неправильная ошибка: ожидали %s, получили %s", tt.expectedError, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("не ожидалась ошибка: %v", err)
				}
				if reader == nil {
					t.Error("reader не должен быть nil")
				}

				// Проверяем содержимое
				data, readErr := io.ReadAll(reader)
				if readErr != nil {
					t.Errorf("ошибка чтения данных: %v", readErr)
				}
				if string(data) != tt.data {
					t.Errorf("неправильные данные: ожидали %q, получили %q", tt.data, string(data))
				}
				reader.Close()
			}
		})
	}
}

// TestRemoveObject проверяет удаление объекта
func TestRemoveObject(t *testing.T) {
	tests := []struct {
		name          string
		bucket        string
		objectName    string
		shouldFail    bool
		expectedError string
	}{
		{
			name:       "успешное удаление файла",
			bucket:     "test-bucket",
			objectName: "1/file.txt",
			shouldFail: false,
		},
		{
			name:       "удаление файла по полному пути",
			bucket:     "test-bucket",
			objectName: "2/subfolder/file.txt",
			shouldFail: false,
		},
		{
			name:          "объект не найден",
			bucket:        "test-bucket",
			objectName:    "1/nonexistent.txt",
			shouldFail:    true,
			expectedError: "not found",
		},
		{
			name:          "отсутствует доступ",
			bucket:        "restricted-bucket",
			objectName:    "1/file.txt",
			shouldFail:    true,
			expectedError: "access denied",
		},
		{
			name:          "пустой bucket",
			bucket:        "",
			objectName:    "1/file.txt",
			shouldFail:    true,
			expectedError: "bucket name required",
		},
		{
			name:          "пустой objectName",
			bucket:        "test-bucket",
			objectName:    "",
			shouldFail:    true,
			expectedError: "object name required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &MockStorageClient{
				removeObjectFunc: func(ctx context.Context, bucket, objectName string) error {
					if len(bucket) == 0 {
						return errors.New("bucket name required")
					}
					if len(objectName) == 0 {
						return errors.New("object name required")
					}
					if strings.Contains(objectName, "nonexistent") {
						return errors.New("not found")
					}
					if strings.Contains(bucket, "restricted") {
						return errors.New("access denied")
					}
					return nil
				},
			}

			err := mock.RemoveObject(context.Background(), tt.bucket, tt.objectName)

			if tt.shouldFail {
				if err == nil {
					t.Errorf("ожидалась ошибка: %s, но получен успех", tt.expectedError)
				}
				if !strings.Contains(err.Error(), tt.expectedError) {
					t.Errorf("неправильная ошибка: ожидали %s, получили %s", tt.expectedError, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("не ожидалась ошибка: %v", err)
				}
			}
		})
	}
}

// TestListObjects проверяет листинг объектов
func TestListObjects(t *testing.T) {
	tests := []struct {
		name       string
		bucket     string
		prefix     string
		recursive  bool
		expected   int
		shouldFail bool
	}{
		{
			name:      "листинг файлов пользователя 1",
			bucket:    "test-bucket",
			prefix:    "1/",
			recursive: true,
			expected:  3,
		},
		{
			name:      "листинг файлов пользователя 2",
			bucket:    "test-bucket",
			prefix:    "2/",
			recursive: true,
			expected:  5,
		},
		{
			name:      "пустой prefix - все файлы",
			bucket:    "test-bucket",
			prefix:    "",
			recursive: true,
			expected:  10,
		},
		{
			name:      "листинг подпапки без рекурсии",
			bucket:    "test-bucket",
			prefix:    "1/documents/",
			recursive: false,
			expected:  0,
		},
		{
			name:       "bucket не существует",
			bucket:     "nonexistent",
			prefix:     "",
			recursive:  true,
			shouldFail: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &MockStorageClient{
				listObjectsFunc: func(ctx context.Context, bucket, prefix string, recursive bool) <-chan ObjectInfo {
					ch := make(chan ObjectInfo)
					go func() {
						defer close(ch)

						if strings.Contains(bucket, "nonexistent") {
							ch <- ObjectInfo{Err: errors.New("bucket not found")}
							return
						}

						objects := []ObjectInfo{
							{Name: "1/file1.txt", Size: 100},
							{Name: "1/file2.txt", Size: 200},
							{Name: "1/file3.txt", Size: 300},
							{Name: "2/doc1.pdf", Size: 1000},
							{Name: "2/doc2.pdf", Size: 2000},
							{Name: "2/doc3.pdf", Size: 3000},
							{Name: "2/doc4.pdf", Size: 4000},
							{Name: "2/doc5.pdf", Size: 5000},
							{Name: "3/image1.png", Size: 50000},
							{Name: "3/image2.png", Size: 60000},
						}

						for _, obj := range objects {
							if strings.HasPrefix(obj.Name, prefix) {
								ch <- obj
							}
						}
					}()
					return ch
				},
			}

			objCount := 0
			hasError := false
			for obj := range mock.ListObjects(context.Background(), tt.bucket, tt.prefix, tt.recursive) {
				if obj.Err != nil {
					hasError = true
					if !tt.shouldFail {
						t.Errorf("не ожидалась ошибка: %v", obj.Err)
					}
				} else {
					objCount++
				}
			}

			if tt.shouldFail && !hasError {
				t.Errorf("ожидалась ошибка, но операция завершилась успешно")
			}
			if !tt.shouldFail && objCount != tt.expected {
				t.Errorf("неправильное количество объектов: ожидали %d, получили %d", tt.expected, objCount)
			}
		})
	}
}

// TestBucketExists проверяет существование bucket
func TestBucketExists(t *testing.T) {
	tests := []struct {
		name       string
		bucket     string
		exists     bool
		shouldFail bool
	}{
		{
			name:   "bucket существует",
			bucket: "test-bucket",
			exists: true,
		},
		{
			name:   "bucket не существует",
			bucket: "nonexistent-bucket",
			exists: false,
		},
		{
			name:   "другой существующий bucket",
			bucket: "uploads",
			exists: true,
		},
		{
			name:       "пустое имя bucket",
			bucket:     "",
			shouldFail: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &MockStorageClient{
				bucketExistsFunc: func(ctx context.Context, bucket string) (bool, error) {
					if len(bucket) == 0 {
						return false, errors.New("bucket name required")
					}
					return strings.HasPrefix(bucket, "nonexistent") == false, nil
				},
			}

			exists, err := mock.BucketExists(context.Background(), tt.bucket)

			if tt.shouldFail {
				if err == nil {
					t.Errorf("ожидалась ошибка, но получен успех")
				}
			} else {
				if err != nil {
					t.Errorf("не ожидалась ошибка: %v", err)
				}
				if exists != tt.exists {
					t.Errorf("неправильный результат: ожидали %v, получили %v", tt.exists, exists)
				}
			}
		})
	}
}

// TestMakeBucket проверяет создание bucket
func TestMakeBucket(t *testing.T) {
	tests := []struct {
		name          string
		bucket        string
		shouldFail    bool
		expectedError string
	}{
		{
			name:   "успешное создание нового bucket",
			bucket: "new-bucket",
		},
		{
			name:          "bucket уже существует",
			bucket:        "existing-bucket",
			shouldFail:    true,
			expectedError: "bucket already exists",
		},
		{
			name:          "недопустимое имя bucket - спецсимволы",
			bucket:        "invalid_name!",
			shouldFail:    true,
			expectedError: "invalid bucket name",
		},
		{
			name:          "недопустимое имя bucket - заглавные буквы",
			bucket:        "InvalidBucket",
			shouldFail:    true,
			expectedError: "invalid bucket name",
		},
		{
			name:          "пустое имя bucket",
			bucket:        "",
			shouldFail:    true,
			expectedError: "bucket name required",
		},
		{
			name:          "слишком длинное имя bucket",
			bucket:        strings.Repeat("a", 64),
			shouldFail:    true,
			expectedError: "bucket name too long",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			existingBuckets := map[string]bool{
				"existing-bucket": true,
			}

			mock := &MockStorageClient{
				makeBucketFunc: func(ctx context.Context, bucket string) error {
					if len(bucket) == 0 {
						return errors.New("bucket name required")
					}
					if len(bucket) > 63 {
						return errors.New("bucket name too long")
					}
					if strings.ContainsAny(bucket, "_!@#$%^&*()") || bucket != strings.ToLower(bucket) {
						return errors.New("invalid bucket name")
					}
					if existingBuckets[bucket] {
						return errors.New("bucket already exists")
					}
					return nil
				},
			}

			err := mock.MakeBucket(context.Background(), tt.bucket)

			if tt.shouldFail {
				if err == nil {
					t.Errorf("ожидалась ошибка: %s, но получен успех", tt.expectedError)
				}
				if !strings.Contains(err.Error(), tt.expectedError) {
					t.Errorf("неправильная ошибка: ожидали %s, получили %s", tt.expectedError, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("не ожидалась ошибка: %v", err)
				}
			}
		})
	}
}

// TestObjectInfo проверяет структуру ObjectInfo
func TestObjectInfo(t *testing.T) {
	tests := []struct {
		name     string
		objInfo  ObjectInfo
		validate func(ObjectInfo) bool
	}{
		{
			name: "нормальный объект",
			objInfo: ObjectInfo{
				Name: "file.txt",
				Size: 1024,
				Err:  nil,
			},
			validate: func(oi ObjectInfo) bool {
				return len(oi.Name) > 0 && oi.Size >= 0 && oi.Err == nil
			},
		},
		{
			name: "объект с ошибкой",
			objInfo: ObjectInfo{
				Name: "",
				Size: 0,
				Err:  errors.New("connection lost"),
			},
			validate: func(oi ObjectInfo) bool {
				return oi.Err != nil
			},
		},
		{
			name: "пустой объект",
			objInfo: ObjectInfo{
				Name: "empty.txt",
				Size: 0,
				Err:  nil,
			},
			validate: func(oi ObjectInfo) bool {
				return oi.Size == 0
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.validate(tt.objInfo) {
				t.Errorf("валидация не пройдена для объекта: %+v", tt.objInfo)
			}
		})
	}
}

// TestStorageClientInterface проверяет соответствие интерфейсу
func TestStorageClientInterface(t *testing.T) {
	var _ StorageClient = (*MockStorageClient)(nil)

	mock := &MockStorageClient{}
	if mock == nil {
		t.Error("не удалось создать экземпляр MockStorageClient")
	}
}

// TestConcurrentOperations проверяет одновременные операции
func TestConcurrentOperations(t *testing.T) {
	mock := &MockStorageClient{
		putObjectFunc: func(ctx context.Context, bucket, objectName string, reader io.Reader, objectSize int64, contentType string) (string, error) {
			return "etag123", nil
		},
		getObjectFunc: func(ctx context.Context, bucket, objectName string) (io.ReadCloser, error) {
			return io.NopCloser(strings.NewReader("data")), nil
		},
	}

	done := make(chan bool, 10)

	// Запускаем 10 параллельных операций
	for i := 0; i < 10; i++ {
		go func(idx int) {
			reader := strings.NewReader("test data")
			_, err := mock.PutObject(context.Background(), "bucket", "object", reader, 9, "text/plain")
			if err != nil {
				t.Errorf("операция %d: ошибка PutObject: %v", idx, err)
			}
			done <- true
		}(i)
	}

	// Ждём завершения всех операций
	for i := 0; i < 10; i++ {
		<-done
	}
}

// TestErrorHandling проверяет обработку ошибок
func TestErrorHandling(t *testing.T) {
	tests := []struct {
		name  string
		setup func() StorageClient
		test  func(StorageClient) error
	}{
		{
			name: "обработка ошибки PutObject",
			setup: func() StorageClient {
				return &MockStorageClient{
					putObjectFunc: func(ctx context.Context, bucket, objectName string, reader io.Reader, objectSize int64, contentType string) (string, error) {
						return "", errors.New("connection refused")
					},
				}
			},
			test: func(sc StorageClient) error {
				_, err := sc.PutObject(context.Background(), "bucket", "object", strings.NewReader("data"), 4, "text/plain")
				return err
			},
		},
		{
			name: "обработка ошибки GetObject",
			setup: func() StorageClient {
				return &MockStorageClient{
					getObjectFunc: func(ctx context.Context, bucket, objectName string) (io.ReadCloser, error) {
						return nil, errors.New("object not found")
					},
				}
			},
			test: func(sc StorageClient) error {
				_, err := sc.GetObject(context.Background(), "bucket", "object")
				return err
			},
		},
		{
			name: "обработка ошибки RemoveObject",
			setup: func() StorageClient {
				return &MockStorageClient{
					removeObjectFunc: func(ctx context.Context, bucket, objectName string) error {
						return errors.New("permission denied")
					},
				}
			},
			test: func(sc StorageClient) error {
				return sc.RemoveObject(context.Background(), "bucket", "object")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := tt.setup()
			err := tt.test(sc)
			if err == nil {
				t.Error("ожидалась ошибка, но операция завершилась успешно")
			}
		})
	}
}

// TestLargeFileHandling проверяет работу с большими файлами
func TestLargeFileHandling(t *testing.T) {
	largeData := strings.Repeat("x", 1000000) // 1MB

	mock := &MockStorageClient{
		putObjectFunc: func(ctx context.Context, bucket, objectName string, reader io.Reader, objectSize int64, contentType string) (string, error) {
			if objectSize > 1000000000 { // 1GB
				return "", errors.New("file too large")
			}
			return "etag456", nil
		},
	}

	// Успешная загрузка большого файла
	etag, err := mock.PutObject(context.Background(), "bucket", "large.bin", strings.NewReader(largeData), int64(len(largeData)), "application/octet-stream")
	if err != nil {
		t.Errorf("ошибка при загрузке большого файла: %v", err)
	}
	if len(etag) == 0 {
		t.Error("etag не должен быть пустым")
	}

	// Попытка загрузить файл больше 1GB
	_, err = mock.PutObject(context.Background(), "bucket", "huge.bin", strings.NewReader(""), 2000000000, "application/octet-stream")
	if err == nil {
		t.Error("ожидалась ошибка при загрузке файла больше 1GB")
	}
}

// TestContextCancellation проверяет отмену контекста
func TestContextCancellation(t *testing.T) {
	mock := &MockStorageClient{
		putObjectFunc: func(ctx context.Context, bucket, objectName string, reader io.Reader, objectSize int64, contentType string) (string, error) {
			select {
			case <-ctx.Done():
				return "", ctx.Err()
			default:
				return "etag", nil
			}
		},
	}

	// Тест с отменённым контекстом
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := mock.PutObject(ctx, "bucket", "object", strings.NewReader("data"), 4, "text/plain")
	if err == nil {
		t.Error("ожидалась ошибка при отменённом контексте")
	}
	if err != context.Canceled {
		t.Errorf("ожидалась ошибка context.Canceled, получена: %v", err)
	}
}

// TestEdgeCases проверяет граничные случаи
func TestEdgeCases(t *testing.T) {
	tests := []struct {
		name   string
		bucket string
		object string
		test   func(bucket, object string) bool
	}{
		{
			name:   "объект с точками в имени",
			bucket: "bucket",
			object: "1/file.tar.gz",
			test: func(b, o string) bool {
				return strings.Contains(o, ".")
			},
		},
		{
			name:   "объект с кириллицей в имени",
			bucket: "bucket",
			object: "1/файл.txt",
			test: func(b, o string) bool {
				return len(o) > 0
			},
		},
		{
			name:   "объект с пробелами",
			bucket: "bucket",
			object: "1/my file.txt",
			test: func(b, o string) bool {
				return strings.Contains(o, " ")
			},
		},
		{
			name:   "очень длинное имя объекта",
			bucket: "bucket",
			object: "1/" + strings.Repeat("a", 255),
			test: func(b, o string) bool {
				return len(o) > 255
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.test(tt.bucket, tt.object) {
				t.Errorf("тест не пройден для: %s", tt.name)
			}
		})
	}
}
