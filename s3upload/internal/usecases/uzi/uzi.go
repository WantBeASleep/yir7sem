package uzi

import (
	"context"
	"fmt"
	"io"

	"yir/s3upload/internal/entity"
	"yir/s3upload/internal/usecases/repo"

	"go.uber.org/zap"
)

type UziUseCase struct {
	s3 repo.S3

	logger *zap.Logger
}

func NewUziUseCase(
	s3 repo.S3,

	logger *zap.Logger,
) *UziUseCase {
	return &UziUseCase{
		s3:     s3,
		logger: logger,
	}
}

func (u *UziUseCase) UploadFile(ctx context.Context, file *entity.File) error {
	u.logger.Info("[Request] Upload file to S3", zap.Any("meta", file.Meta))
	if err := u.s3.Upload(ctx, file); err != nil {
		u.logger.Info("Upload file to S3", zap.Error(err))
		return fmt.Errorf("upload file to S3 [path %q]: %w", file.Meta.Path, err)
	}
	u.logger.Info("[Response] Uploaded file to S3")

	return nil
}

func (u *UziUseCase) GetFile(ctx context.Context, path string) (*entity.FileMeta, io.Reader, error) {
	u.logger.Info("[Request] Get file from S3", zap.String("path", path))
	meta, stream, err := u.s3.Get(ctx, path)
	if err != nil {
		u.logger.Error("Get file from S3", zap.Error(err))
		return nil, nil, fmt.Errorf("get file [path %q]: %w", path, err)
	}
	u.logger.Info("[Response] Got file from S3")

	return meta, stream, nil
}
