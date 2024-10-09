package usecases

import (
	"context"

	"github.com/google/uuid"
)

type Uzi interface {
	UploadAndSplitUziFile(ctx context.Context, img []byte) (uuid.UUID, uuid.UUIDs, error)
}
