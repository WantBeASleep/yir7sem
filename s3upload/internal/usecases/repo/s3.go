package repo

import (
	"context"
	"io"
)

type S3 interface {
	Upload(ctx context.Context, path string, filename string, data []byte) error

	// стримим объект поэтому тут ReadCloser
	Get(ctx context.Context, path string) (io.ReadCloser, error)
}
