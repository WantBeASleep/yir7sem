package repositories

import (
	"context"
)

type S3 interface {
	Upload(ctx context.Context, path string, data []byte) error

	FullGetByPath(ctx context.Context, path string) ([]byte, error)
}
