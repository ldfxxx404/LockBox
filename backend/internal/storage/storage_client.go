package storage

import (
	"context"
	"io"
)

type StorageClient interface {
	PutObject(ctx context.Context, bucket, objectName string, reader io.Reader, objectSize int64, contentType string) (etag string, err error)

	GetObject(ctx context.Context, bucket, objectName string) (io.ReadCloser, error)

	RemoveObject(ctx context.Context, bucket, objectName string) error

	ListObjects(ctx context.Context, bucket, prefix string, recursive bool) <-chan ObjectInfo

	BucketExists(ctx context.Context, bucket string) (exists bool, err error)

	MakeBucket(ctx context.Context, bucket string) error
}

type ObjectInfo struct {
	Name string
	Size int64
	Err  error
}
