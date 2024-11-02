package uzi

import (
	"context"
	"fmt"
	"yir/uzi/internal/usecases/dto"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (u *UziUseCase) GetReport(ctx context.Context, id uuid.UUID) (*dto.Report, error) {
	u.logger.Debug("[Request] Get Uzi", zap.String("uzi id", id.String()))
	uzi, err := u.uziRepo.GetUziByID(ctx, id)
	if err != nil {
		u.logger.Error("Get uzi", zap.Error(err))
		return nil, fmt.Errorf("get uzi: %w", err)
	}
	u.logger.Debug("[Response] Got uzi", zap.Any("Uzi", uzi))

	u.logger.Debug("[Request] Get uzi images", zap.String("uzi id", id.String()))
	images, err := u.uziRepo.GetImagesByUziID(ctx, id)
	if err != nil {
		u.logger.Error("Get uzi images", zap.Error(err))
		return nil, fmt.Errorf("get uzi images: %w", err)
	}
	u.logger.Debug("[Response] Got uzi images", zap.Int("Count images", len(images)))

	u.logger.Debug("[Request] Get uzi formations", zap.String("uzi id", id.String()))
	formations, err := u.uziRepo.GetFormationsByUziID(ctx, id)
	if err != nil {
		u.logger.Error("Get uzi formations", zap.Error(err))
		return nil, fmt.Errorf("get uzi formations: %w", err)
	}
	u.logger.Debug("[Response] Got uzi formations", zap.Int("Count formations", len(formations)))

	formationsWithTirads, err := u.GetDTOFormations(ctx, formations)
	if err != nil {
		return nil, fmt.Errorf("get dto formations: %w", err)
	}

	u.logger.Debug("[Request] Get uzi segments", zap.String("uzi id", id.String()))
	segments, err := u.uziRepo.GetSegmentsByUziID(ctx, id)
	if err != nil {
		u.logger.Error("Get uzi segments", zap.Error(err))
		return nil, fmt.Errorf("get uzi segments: %w", err)
	}
	u.logger.Debug("[Response] Got uzi segments", zap.Int("Count segments", len(segments)))

	segmentsWithTirads, err := u.GetDTOSegments(ctx, segments)
	if err != nil {
		return nil, fmt.Errorf("get dto segments: %w", err)
	}

	return &dto.Report{
		Uzi:        uzi,
		Images:     images,
		Formations: formationsWithTirads,
		Segments:   segmentsWithTirads,
	}, nil
}
