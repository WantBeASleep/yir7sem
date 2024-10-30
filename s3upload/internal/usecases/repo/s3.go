package repo

import (
	"context"
	"io"
	"yir/s3upload/internal/entity"
)

type S3 interface {
	Upload(ctx context.Context, file *entity.File) error

	// стримим объект поэтому тут ReadCloser
	Get(ctx context.Context, path string, opts ...entity.GetOption) (*entity.FileMeta, io.Reader, error)
}
