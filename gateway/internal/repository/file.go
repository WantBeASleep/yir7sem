package repository

import (
	"context"
	"errors"
	"fmt"
	"io"

	"gateway/internal/domain"

	minio "github.com/minio/minio-go/v7"
)

var ErrFileNotFound = errors.New("file not found")

type FileRepo interface {
	GetFile(ctx context.Context, path string) (io.ReadCloser, error)
	LoadFile(ctx context.Context, path string, file domain.File) error
}

type fileRepo struct {
	s3     *minio.Client
	bucket string
}

func (r *fileRepo) GetFile(ctx context.Context, path string) (io.ReadCloser, error) {
	_, err := r.s3.StatObject(ctx, r.bucket, path, minio.GetObjectOptions{})
	if err != nil {
		errResp := minio.ToErrorResponse(err)
		if errResp.Code == "NoSuchKey" {
			return nil, ErrFileNotFound
		}
		return nil, fmt.Errorf("get stat of object: %w", err)
	}

	obj, err := r.s3.GetObject(ctx, r.bucket, path, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (r *fileRepo) LoadFile(ctx context.Context, path string, file domain.File) error {
	_, err := r.s3.PutObject(ctx, r.bucket, path, file.Buf, -1, minio.PutObjectOptions{ContentType: file.Format})
	if err != nil {
		return err
	}

	return nil
}
