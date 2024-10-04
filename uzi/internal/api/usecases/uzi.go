package usecases

import (
	"context"
	"yir/uzi/internal/entity/dto"

	"github.com/google/uuid"
)

type Uzi interface {
	InsertUzi(ctx context.Context, req *dto.Uzi) error
	GetUzi(ctx context.Context, id uuid.UUID) (*dto.Uzi, error)
}
