package uzi

import (
	"context"
	"fmt"
	"yir/uzi/internal/entity"
	"yir/uzi/internal/usecases/dto"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (u *UziUseCase) CreateImages(ctx context.Context, images []entity.Image) (uuid.UUIDs, error) {
	// generate uuid
	resp := make([]uuid.UUID, 0, len(images))
	for i := range images {
		id := uuid.New()
		resp = append(resp, id)
		images[i].Id = id
	}

	u.logger.Debug("[Request] Create images")
	if err := u.uziRepo.InsertImages(ctx, images); err != nil {
		u.logger.Error("Create images", zap.Error(err))
		return nil, fmt.Errorf("create images: %w", err)
	}
	u.logger.Debug("[Response] Created images")

	return resp, nil
}

// TODO: после formation/segments в базе
func (u *UziUseCase) GetImageWithFormationsSegments(ctx context.Context, id uuid.UUID) (*dto.ImageWithFormationsSegments, error) {
	u.logger.Debug("[Request] Get image by ID", zap.String("image id", id.String()))
	image, err := u.uziRepo.GetImageByID(ctx, id)
	if err != nil {
		u.logger.Error("Get image by ID", zap.Error(err))
		return nil, fmt.Errorf("get image by ID: %w", err)
	}
	u.logger.Debug("[Response] Got image by ID", zap.String("image id", id.String()))

	u.logger.Debug("[Request] Get image formations", zap.String("image id", id.String()))
	formations, err := u.uziRepo.GetFormationsByImageID(ctx, id)
	if err != nil {
		u.logger.Error("Get image formations", zap.Error(err))
		return nil, fmt.Errorf("get image formations: %w", err)
	}
	u.logger.Debug("[Response] Got image formations", zap.Int("count formations", len(formations)))

	formationsWithTirads, err := u.GetDTOFormationsFromEntity(ctx, formations)
	if err != nil {
		return nil, fmt.Errorf("get dto formations: %w", err)
	}

	u.logger.Debug("[Request] Get image segments", zap.String("image id", id.String()))
	segments, err := u.uziRepo.GetSegmentsByImageID(ctx, id)
	if err != nil {
		u.logger.Error("Get image segments", zap.Error(err))
		return nil, fmt.Errorf("get image segments: %w", err)
	}
	u.logger.Debug("[Response] Got image segments", zap.Int("count segments", len(segments)))

	segmentsWithTirads, err := u.GetDTOSegmentsFromEntity(ctx, segments)
	if err != nil {
		return nil, fmt.Errorf("get dto segments: %w", err)
	}

	return &dto.ImageWithFormationsSegments{
		Image:      image,
		Formations: formationsWithTirads,
		Segments:   segmentsWithTirads,
	}, nil
}
