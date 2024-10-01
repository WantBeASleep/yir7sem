package uzi

import (
	"context"
	"fmt"
	"yir/pkg/mappers"
	"yir/uzi/internal/entity"
	"yir/uzi/internal/entity/dto"

	"go.uber.org/zap"
)

func (u *UziUseCase) InsertDTOSegments(ctx context.Context, segments []dto.Segment) error {
	u.logger.Debug("[Request] Create segments tirads")
	entitySegments, err := mappers.TransformSlice[dto.Segment, entity.Segment](segments, func(src *dto.Segment, dst *entity.Segment) error {
		tiradsID, err := u.uziRepo.CreateTirads(ctx, src.Tirads)
		if err != nil {
			return fmt.Errorf("create tirads: %w", err)
		}

		dst.TiradsID = tiradsID
		return nil
	})
	if err != nil {
		u.logger.Error("Create segments tirads", zap.Error(err))
		return fmt.Errorf("create segments tirads: %w", err)
	}
	u.logger.Debug("[Response] Created segments tirads")

	u.logger.Debug("[Request] Insert segments")
	if err := u.uziRepo.InsertSegments(ctx, entitySegments); err != nil {
		u.logger.Error("Insert segments", zap.Error(err))
		return fmt.Errorf("insert segments: %w", err)
	}
	u.logger.Debug("[Response] Inserted segments")

	return nil
}

func (u *UziUseCase) GetDTOSegments(ctx context.Context, segments []entity.Segment) ([]dto.Segment, error) {
	u.logger.Debug("[Request] Get segments tirads")
	dtoSegments, err := mappers.TransformSlice[entity.Segment, dto.Segment](segments, func(src *entity.Segment, dst *dto.Segment) error {
		tirads, err := u.uziRepo.GetTiradsByID(ctx, src.TiradsID)
		if err != nil {
			return fmt.Errorf("get tirads [id %q]: %w", src.TiradsID, err)
		}

		dst.Tirads = tirads
		return nil
	})
	if err != nil {
		u.logger.Error("Get segments tirads", zap.Error(err))
		return nil, fmt.Errorf("get segments tirads: %w", err)
	}
	u.logger.Debug("[Response] Got segments tirads")

	return dtoSegments, nil
}