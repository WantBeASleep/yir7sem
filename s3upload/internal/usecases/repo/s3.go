package repo

import (
	"context"
)

type S3 interface {
	Upload(ctx context.Context, path string, filename string, data []byte) error
}
