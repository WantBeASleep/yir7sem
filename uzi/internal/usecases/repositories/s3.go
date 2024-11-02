package repositories

import (
	"context"
	"yir/uzi/internal/entity/imagesplitter"
)

type S3 interface {
	Upload(ctx context.Context, path string, file *imagesplitter.File) error

	GetFile(ctx context.Context, path string) (*imagesplitter.File, error)
}
