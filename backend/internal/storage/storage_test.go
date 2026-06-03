package storage

import (
	"context"
	"errors"
	"io"
	"strings"
	"testing"
)

// MockStorageClient is a mock for StorageClient interface
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

// TestPutObject checks object upload
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
			name:        "successful text file upload",
			bucket:      "test-bucket",
			objectName:  "1/file.txt",
			data:        "hello world",
			size:        11,
			contentType: "text/plain",
		},
		{
			name:        "JSON file upload",
			bucket:      "test-bucket",
			objectName:  "2/data.json",
			data:        `{"key": "value"}`,
			size:        16,
			contentType: "application/json",
		},
		{
			name:        "image upload",
			bucket:      "test-bucket",
			objectName:  "1/image.png",
			data:        "\x89PNG\r\n\x1a\n",
			size:        8,
			contentType: "image/png",
		},
		{
			name:          "empty bucket",
			bucket:        "",
			objectName:    "1/file.txt",
			size:          10,
			contentType:   "text/plain",
			shouldFail:    true,
			expectedError: "bucket name required",
		},
		{
			name:          "empty object name",
			bucket:        "test-bucket",
			objectName:    "",
			size:          10,
			contentType:   "text/plain",
			shouldFail:    true,
			expectedError: "object name required",
		},
		{
			name:          "invalid object path",
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

			_, err := mock.PutObject(context.Background(), tt.bucket, tt.objectName, strings.NewReader(tt.data), tt.size, tt.contentType)

			if tt.shouldFail {
				if err == nil {
					t.Errorf("expected error: %s, but got success", tt.expectedError)
				}
				if !strings.Contains(err.Error(), tt.expectedError) {
					t.Errorf("expected error %s, got %s", tt.expectedError, err.Error())
				}
			}
		})
	}
}

// TestGetObject checks object retrieval
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
			name:       "successful file retrieval",
			bucket:     "test-bucket",
			objectName: "1/file.txt",
			data:       "test content",
		},
		{
			name:       "large file retrieval",
			bucket:     "test-bucket",
			objectName: "2/large.bin",
			data:       strings.Repeat("x", 10000),
		},
		{
			name:          "object not found",
			bucket:        "test-bucket",
			objectName:    "1/nonexistent.txt",
			shouldFail:    true,
			expectedError: "not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &MockStorageClient{
				getObjectFunc: func(ctx context.Context, bucket, objectName string) (io.ReadCloser, error) {
					if strings.Contains(objectName, "nonexistent") {
						return nil, errors.New("not found")
					}
					return io.NopCloser(strings.NewReader(tt.data)), nil
				},
			}

			reader, err := mock.GetObject(context.Background(), tt.bucket, tt.objectName)

			if tt.shouldFail {
				if err == nil {
					t.Errorf("expected error: %s", tt.expectedError)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			data, readErr := io.ReadAll(reader)
			if readErr != nil {
				t.Errorf("read error: %v", readErr)
			}

			if err := reader.Close(); err != nil {
				t.Errorf("close error: %v", err)
			}

			if string(data) != tt.data {
				t.Errorf("expected %q, got %q", tt.data, string(data))
			}
		})
	}
}
