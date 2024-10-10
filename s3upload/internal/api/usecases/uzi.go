package usecases

import (
	"context"
	"io"

	"github.com/google/uuid"
)

type Uzi interface {
	UploadAndSplitUziFile(ctx context.Context, img []byte) (uuid.UUID, uuid.UUIDs, error)

	GetByPath(ctx context.Context, path string) (io.ReadCloser, error)
}
