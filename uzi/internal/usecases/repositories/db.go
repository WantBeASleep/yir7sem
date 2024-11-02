package repositories

import (
	"context"
	"yir/uzi/internal/entity"

	"github.com/google/uuid"
)

type UziRepo interface {
	// uzi
	CreateUzi(ctx context.Context, uzi *entity.Uzi) (uuid.UUID, error)
	GetUziByID(ctx context.Context, id uuid.UUID) (*entity.Uzi, error)
	UpdateUzi(ctx context.Context, id uuid.UUID, uzi *entity.Uzi) (*entity.Uzi, error)

	// image
	CreateImages(ctx context.Context, images []entity.Image) (uuid.UUIDs, error)
	GetImagesByUziID(ctx context.Context, uziID uuid.UUID) ([]entity.Image, error)
	GetImageByID(ctx context.Context, id uuid.UUID) (*entity.Image, error)

	// formations
	CreateFormation(ctx context.Context, formation *entity.Formation) (uuid.UUID, error)
	GetFormationByID(ctx context.Context, id uuid.UUID) (*entity.Formation, error)
	GetFormationsByUziID(ctx context.Context, uziID uuid.UUID) ([]entity.Formation, error)
	GetFormationsByImageID(ctx context.Context, imageID uuid.UUID) ([]entity.Formation, error)
	UpdateFormation(ctx context.Context, id uuid.UUID, formation *entity.Formation) (*entity.Formation, error)

	// segments
	CreateSegments(ctx context.Context, segments []entity.Segment) (uuid.UUIDs, error)
	GetSegmentsByUziID(ctx context.Context, uziID uuid.UUID) ([]entity.Segment, error)
	GetSegmentsByImageID(ctx context.Context, imageID uuid.UUID) ([]entity.Segment, error)
	GetSegmentsByFormationID(ctx context.Context, formationID uuid.UUID) ([]entity.Segment, error)
	UpdateSegment(ctx context.Context, id uuid.UUID, segment *entity.Segment) (*entity.Segment, error)

	// tirads
	CreateTirads(ctx context.Context, tirads *entity.Tirads) (int, error)
	GetTiradsByID(ctx context.Context, id int) (*entity.Tirads, error)

	// device
	GetDevicesList(ctx context.Context) ([]entity.Device, error)
}
