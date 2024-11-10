package usecases

import (
	"context"
	"yir/uzi/internal/entity"
	"yir/uzi/internal/usecases/dto"

	"github.com/google/uuid"
)

type Uzi interface {
	// report
	GetReport(ctx context.Context, id uuid.UUID) (*dto.Report, error)

	// uzi
	CreateUzi(ctx context.Context, req *entity.Uzi) (uuid.UUID, error)
	GetUzi(ctx context.Context, id uuid.UUID) (*entity.Uzi, error)
	UpdateUzi(ctx context.Context, id uuid.UUID, req *entity.Uzi) (*entity.Uzi, error)

	// images
	GetImageWithFormationsSegments(ctx context.Context, id uuid.UUID) (*dto.ImageWithFormationsSegments, error)

	// formations
	CreateFormationWithSegments(ctx context.Context, req *dto.FormationWithSegments) (uuid.UUID, uuid.UUIDs, error)
	InsertFormationsAndSegemetsSeparately(ctx context.Context, formations []dto.Formation, segments []dto.Segment) error
	GetFormationWithSegments(ctx context.Context, id uuid.UUID) (*dto.FormationWithSegments, error)
	UpdateFormation(ctx context.Context, id uuid.UUID, req *dto.Formation) (*dto.Formation, error)

	// segments
	UpdateSegment(ctx context.Context, id uuid.UUID, segment *dto.Segment) (*dto.Segment, error)

	// device
	DeviceList(ctx context.Context) ([]entity.Device, error)

	// splitted
	SplitLoadSaveUzi(ctx context.Context, uziID uuid.UUID) (uuid.UUIDs, error)
}
