package repo

import (
	"context"
	"io"
	"yir/s3upload/internal/entity"
)

type S3 interface {
	Upload(ctx context.Context, path string, filename string, data []byte, meta *entity.ImageMetaData) error

	// стримим объект поэтому тут ReadCloser
	Get(ctx context.Context, path string) (io.ReadCloser, error)
}
