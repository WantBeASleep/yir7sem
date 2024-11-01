package repositories

import (
	"context"
	"yir/uzi/internal/entity"

	"github.com/google/uuid"
)

type UziRepo interface {
	CreateUzi(ctx context.Context, uzi *entity.Uzi) (uuid.UUID, error)
	GetUziByID(ctx context.Context, id uuid.UUID) (*entity.Uzi, error)
}
