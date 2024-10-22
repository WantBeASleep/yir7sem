package repo

import (
	"context"
	"io"
	"yir/s3upload/internal/entity"
)

type S3 interface {
	Upload(ctx context.Context, path string, data io.Reader, opts ...entity.LoadOption) error

	// стримим объект поэтому тут ReadCloser
	Get(ctx context.Context, path string, opts ...entity.GetOption) (io.Reader, error)
}
