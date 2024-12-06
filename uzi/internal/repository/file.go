package repository

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"uzi/internal/domain"

	minio "github.com/minio/minio-go/v7"
)

var ErrFileNotFound = errors.New("file not found")

type FileRepo interface {
	GetFileViaTemp(ctx context.Context, path string) (domain.File, func() error, error)
	LoadFile(ctx context.Context, path string, file domain.File) error
}

type fileRepo struct {
	s3     *minio.Client
	bucket string
}

var emptyCloser = func() error { return nil }

func (r *fileRepo) GetFileViaTemp(ctx context.Context, path string) (domain.File, func() error, error) {
	_, err := r.s3.StatObject(ctx, r.bucket, path, minio.GetObjectOptions{})
	if err != nil {
		errResp := minio.ToErrorResponse(err)
		if errResp.Code == "NoSuchKey" {
			return domain.File{}, emptyCloser, ErrFileNotFound
		}
		return domain.File{}, emptyCloser, fmt.Errorf("get stat of object: %w", err)
	}

	obj, err := r.s3.GetObject(ctx, r.bucket, path, minio.GetObjectOptions{})
	if err != nil {
		return domain.File{}, emptyCloser, fmt.Errorf("get file from s3: %w", err)
	}
	defer obj.Close()

	objInfo, err := obj.Stat()
	if err != nil {
		return domain.File{}, emptyCloser, fmt.Errorf("get s3 obj info: %w", err)
	}

	file, err := os.CreateTemp("/tmp", "fnus-*.tmp")
	if err != nil {
		return domain.File{}, emptyCloser, err
	}
	closer := func() error { return errors.Join(file.Close(), os.Remove(file.Name())) }

	if _, err := io.Copy(file, obj); err != nil {
		closer()
		return domain.File{}, emptyCloser, fmt.Errorf("copy to tmp: %w", err)
	}

	stat, err := file.Stat()
	if err != nil {
		closer()
		return domain.File{}, emptyCloser, fmt.Errorf("get stat of file: %w", err)
	}

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		closer()
		return domain.File{}, emptyCloser, fmt.Errorf("seek file: %w", err)
	}

	return domain.File{
		Format: objInfo.ContentType,
		Size:   stat.Size(),
		Buf:    file,
	}, closer, nil
}

func (r *fileRepo) LoadFile(ctx context.Context, path string, file domain.File) error {
	sizeUploaded, err := r.s3.PutObject(ctx, r.bucket, path, file.Buf, file.Size, minio.PutObjectOptions{ContentType: file.Format})
	if err != nil {
		return fmt.Errorf("put obj to s3: %w", err)
	}

	if file.Size != sizeUploaded.Size {
		return errors.New("file not upload to s3 completely")
	}

	return nil
}
