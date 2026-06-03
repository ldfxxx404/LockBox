package storage

import (
	"context"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinIOAdapter struct {
	client *minio.Client
}

func NewMinIOAdapter(endpoint, accessKey, secretKey string, useSSL bool) (*MinIOAdapter, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}
	return &MinIOAdapter{client: client}, nil
}

func (m *MinIOAdapter) PutObject(ctx context.Context, bucket, objectName string, reader io.Reader, objectSize int64, contentType string) (string, error) {
	info, err := m.client.PutObject(ctx, bucket, objectName, reader, objectSize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	return info.ETag, err
}

func (m *MinIOAdapter) GetObject(ctx context.Context, bucket, objectName string) (io.ReadCloser, error) {
	obj, err := m.client.GetObject(ctx, bucket, objectName, minio.GetObjectOptions{})
	return obj, err
}

func (m *MinIOAdapter) RemoveObject(ctx context.Context, bucket, objectName string) error {
	return m.client.RemoveObject(ctx, bucket, objectName, minio.RemoveObjectOptions{})
}

func (m *MinIOAdapter) ListObjects(ctx context.Context, bucket, prefix string, recursive bool) <-chan ObjectInfo {
	out := make(chan ObjectInfo)

	go func() {
		defer close(out)
		for objInfo := range m.client.ListObjects(ctx, bucket, minio.ListObjectsOptions{
			Prefix:    prefix,
			Recursive: recursive,
		}) {
			out <- ObjectInfo{
				Name: objInfo.Key,
				Size: objInfo.Size,
				Err:  objInfo.Err,
			}
		}
	}()

	return out
}

func (m *MinIOAdapter) BucketExists(ctx context.Context, bucket string) (bool, error) {
	return m.client.BucketExists(ctx, bucket)
}

func (m *MinIOAdapter) MakeBucket(ctx context.Context, bucket string) error {
	return m.client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{})
}
