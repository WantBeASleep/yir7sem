package repositories

import (
	"context"
	"yir/uzi/internal/entity"

	"github.com/google/uuid"
)

type UziRepo interface {
	GetDevicesList(ctx context.Context) ([]*entity.Device, error)
	InsertUzi(ctx context.Context, uzi *entity.Uzi) error
	InsertImages(ctx context.Context, images []entity.Image) error
	InsertFormationsWithImageFormations(ctx context.Context, formations []entity.DBFormation) error
	GetUzi(ctx context.Context, id uuid.UUID) (*entity.Uzi, error)
	GetDevice(ctx context.Context, id int) (*entity.Device, error)
	GetUziImages(ctx context.Context, uziID uuid.UUID) ([]entity.Image, error)
}
