package usecases

import (
	"context"
	"yir/uzi/internal/entity"
	"yir/uzi/internal/entity/dto"

	"github.com/google/uuid"
)

type Uzi interface {
	InsertUzi(ctx context.Context, req *dto.Uzi) error
	CreateUziInfo(ctx context.Context, req *entity.Uzi) error

	GetUzi(ctx context.Context, id uuid.UUID) (*dto.Uzi, error)
	GetUziInfo(ctx context.Context, id uuid.UUID) (*entity.Uzi, error)
	UpdateUziInfo(ctx context.Context, id uuid.UUID, req *entity.Uzi) error

	GetImageWithSegmentsFormations(ctx context.Context, id uuid.UUID) (*dto.ImageWithSegmentsFormations, error)

	InsertFormationWithSegments(ctx context.Context, req *dto.FormationWithSegments) error
	GetFormationWithSegments(ctx context.Context, id uuid.UUID) (*dto.FormationWithSegments, error)
	UpdateFormation(ctx context.Context, id uuid.UUID, req *dto.Formation) error

	DeviceList(ctx context.Context) ([]entity.Device, error)

	SplitLoadSaveUzi(ctx context.Context, uziID uuid.UUID) (uuid.UUIDs, error)
}
