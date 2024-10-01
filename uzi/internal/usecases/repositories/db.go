package repositories

import (
	"context"
	"yir/uzi/internal/entity"

	"github.com/google/uuid"
)

type UziRepo interface {
	GetDevicesList(ctx context.Context) ([]entity.Device, error)

	CreateTirads(ctx context.Context, tirads *entity.Tirads) (int, error)
	GetTiradsByID(ctx context.Context, id int) (*entity.Tirads, error)
	UpdateTirads(ctx context.Context, id int, tirads *entity.Tirads) error

	InsertUzi(ctx context.Context, uzi *entity.Uzi) error
	GetUziByID(ctx context.Context, id uuid.UUID) (*entity.Uzi, error)
	UpdateUzi(ctx context.Context, id uuid.UUID, uzi *entity.Uzi) error

	InsertImages(ctx context.Context, images []entity.Image) error
	GetUziImages(ctx context.Context, uziID uuid.UUID) ([]entity.Image, error)

	InsertFormations(ctx context.Context, formations []entity.Formation) error
	GetUziFormations(ctx context.Context, uziID uuid.UUID) ([]entity.Formation, error)
	GetFormationByID(ctx context.Context, id uuid.UUID) (*entity.Formation, error)

	InsertSegments(ctx context.Context, segments []entity.Segment) error
	GetUziSegments(ctx context.Context, uziID uuid.UUID) ([]entity.Segment, error)
	GetFormationSegments(ctx context.Context, formationID uuid.UUID) ([]entity.Segment, error)
}
