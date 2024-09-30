package repositories

import (
	"context"
	"yir/uzi/internal/entity"

	"github.com/google/uuid"
)

type UziRepo interface {
	GetDevicesList(ctx context.Context) ([]entity.Device, error)

	CreateTirads(ctx context.Context, tirads *entity.Tirads) (int, error)
	GetTirads(ctx context.Context, id int) (*entity.Tirads, error)

	InsertUzi(ctx context.Context, uzi *entity.Uzi) error
	GetUzi(ctx context.Context, id uuid.UUID) (*entity.Uzi, error)

	InsertImages(ctx context.Context, images []entity.Image) error
	GetUziImages(ctx context.Context, uziID uuid.UUID) ([]entity.Image, error)

	InsertFormations(ctx context.Context, formations []entity.Formation) error
	GetUziFormations(ctx context.Context, uziID uuid.UUID) ([]entity.Formation, error)

	InsertSegments(ctx context.Context, segments []entity.Segment) error
	GetUziSegments(ctx context.Context, uziID uuid.UUID) ([]entity.Segment, error)
}
