package uzi

import (
	"context"
	"fmt"
	"yir/pkg/mappers"
	"yir/uzi/internal/entity"
	"yir/uzi/internal/usecases/dto"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

// создаст id
func (u *UziUseCase) CreateDTOSegments(ctx context.Context, segments []dto.Segment) (uuid.UUIDs, error) {
	entitySegments := make([]entity.Segment, 0, len(segments))
	for _, seg := range segments {
		u.logger.Debug("[Request] Create tirads")
		tiradsID, err := u.uziRepo.CreateTirads(ctx, seg.Tirads)
		if err != nil {
			u.logger.Error("create tirads", zap.Error(err))
			return nil, fmt.Errorf("create tirads: %w", err)
		}
		u.logger.Debug("[Response] Created tirads")

		entSeg := mappers.MustTransformObj[dto.Segment, entity.Segment](&seg)
		entSeg.TiradsID = tiradsID
		entitySegments = append(entitySegments, *entSeg)
	}

	// generate uuid
	resp := make([]uuid.UUID, 0, len(entitySegments))
	for i := range entitySegments {
		id := uuid.New()
		resp = append(resp, id)
		entitySegments[i].Id = id
	}

	u.logger.Debug("[Request] Insert segments")
	if err := u.uziRepo.InsertSegments(ctx, entitySegments); err != nil {
		u.logger.Error("Insert segments", zap.Error(err))
		return nil, fmt.Errorf("insert segments: %w", err)
	}
	u.logger.Debug("[Response] Inserted segments")

	return resp, nil
}

func (u *UziUseCase) InsertDTOSegments(ctx context.Context, segments []dto.Segment) error {
	entitySegments := make([]entity.Segment, 0, len(segments))
	for _, seg := range segments {
		u.logger.Debug("[Request] Create tirads")
		tiradsID, err := u.uziRepo.CreateTirads(ctx, seg.Tirads)
		if err != nil {
			u.logger.Error("create tirads", zap.Error(err))
			return fmt.Errorf("create tirads: %w", err)
		}
		u.logger.Debug("[Response] Created tirads")

		entSeg := mappers.MustTransformObj[dto.Segment, entity.Segment](&seg)
		entSeg.TiradsID = tiradsID
		entitySegments = append(entitySegments, *entSeg)
	}

	u.logger.Debug("[Request] Insert segments")
	if err := u.uziRepo.InsertSegments(ctx, entitySegments); err != nil {
		u.logger.Error("Insert segments", zap.Error(err))
		return fmt.Errorf("insert segments: %w", err)
	}
	u.logger.Debug("[Response] Inserted segments")

	return nil
}

func (u *UziUseCase) GetDTOSegmentFromEntity(ctx context.Context, segment *entity.Segment) (*dto.Segment, error) {
	u.logger.Debug("[Request] Get segment tirads", zap.Int("tirads id", segment.TiradsID))
	tirads, err := u.uziRepo.GetTiradsByID(ctx, segment.TiradsID)
	if err != nil {
		return nil, fmt.Errorf("get tirads [id %q]: %w", segment.TiradsID, err)
	}
	u.logger.Debug("[Response] Got segment tirads")

	dtoSegment := mappers.MustTransformObj[entity.Segment, dto.Segment](segment)
	dtoSegment.Tirads = tirads

	return dtoSegment, nil
}

func (u *UziUseCase) GetDTOSegmentsFromEntity(ctx context.Context, segments []entity.Segment) ([]dto.Segment, error) {
	dtoSegments := make([]dto.Segment, 0, len(segments))
	for _, seg := range segments {
		dtoSeg, err := u.GetDTOSegmentFromEntity(ctx, &seg)
		if err != nil {
			return nil, fmt.Errorf("get dto segment: %w", err)
		}

		dtoSegments = append(dtoSegments, *dtoSeg)
	}

	return dtoSegments, nil
}

func (u *UziUseCase) UpdateSegment(ctx context.Context, id uuid.UUID, segment *dto.Segment) (*dto.Segment, error) {
	u.logger.Debug("[Request] Create tirads")
	tiradsID, err := u.uziRepo.CreateTirads(ctx, segment.Tirads)
	if err != nil {
		u.logger.Error("Create tirads", zap.Error(err))
		return nil, fmt.Errorf("create tirads: %w", err)
	}
	u.logger.Debug("[Response] Created tirads", zap.Int("id", tiradsID))

	entSegment := mappers.MustTransformObj[dto.Segment, entity.Segment](segment)
	entSegment.TiradsID = tiradsID

	u.logger.Debug("[Request] Update segment", zap.String("segment id", id.String()))
	updateSegment, err := u.uziRepo.UpdateSegment(ctx, id, entSegment)
	if err != nil {
		u.logger.Error("Update segment", zap.Error(err))
		return nil, fmt.Errorf("update segment: %w", err)
	}
	u.logger.Debug("[Response] Updated segment")

	dtoUpdateSegment, err := u.GetDTOSegmentFromEntity(ctx, updateSegment)
	if err != nil {
		return nil, fmt.Errorf("get updated segment: %w", err)
	}

	return dtoUpdateSegment, nil
}
