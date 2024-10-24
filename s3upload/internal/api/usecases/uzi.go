package usecases

import (
	"context"
	"io"
	"yir/s3upload/internal/entity"
)

type Uzi interface {
	UploadFile(ctx context.Context, file *entity.File) error

	GetFile(ctx context.Context, path string) (io.Reader, error)
}
