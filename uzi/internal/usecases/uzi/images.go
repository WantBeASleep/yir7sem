package uzi

import (
	"context"
	"fmt"
	"yir/uzi/internal/entity"
	"yir/uzi/internal/entity/dto"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (u *UziUseCase) InsertImages(ctx context.Context, images []entity.Image) error {
	u.logger.Debug("[Request] Insert images")
	if err := u.uziRepo.InsertImages(ctx, images); err != nil {
		u.logger.Error("Insert images", zap.Error(err))
		return fmt.Errorf("insert images: %w", err)
	}
	u.logger.Debug("[Response] Inserted images")

	return nil
}

func (u *UziUseCase) GetImageWithSegmentsFormations(ctx context.Context, id uuid.UUID) (*dto.ImageWithSegmentsFormations, error) {
	u.logger.Debug("[Request] Get image by ID", zap.Any("id", id))
	image, err := u.uziRepo.GetImageByID(ctx, id)
	if err != nil {
		u.logger.Error("Get image by ID", zap.Error(err))
		return nil, fmt.Errorf("get image by ID: %w", err)
	}
	u.logger.Debug("[Response] Got image by ID", zap.Any("image", image))

	u.logger.Debug("[Request] Get image formations", zap.Any("id", id))
	formations, err := u.uziRepo.GetImageFormations(ctx, id)
	if err != nil {
		u.logger.Error("Get image formations", zap.Error(err))
		return nil, fmt.Errorf("get image formations: %w", err)
	}
	u.logger.Debug("[Response] Got image formations", zap.Any("formations", formations))

	u.logger.Debug("[Request] Get image segments", zap.Any("id", id))
	segments, err := u.uziRepo.GetImageSegments(ctx, id)
	if err != nil {
		u.logger.Error("Get image segments", zap.Error(err))
		return nil, fmt.Errorf("get image segments: %w", err)
	}
	u.logger.Debug("[Response] Got image segments", zap.Any("segments", segments))

	dtoFormations, err := u.GetDTOFormation(ctx, formations)
	if err != nil {
		return nil, fmt.Errorf("get dto formations: %w", err)
	}

	dtoSegments, err := u.GetDTOSegments(ctx, segments)
	if err != nil {
		return nil, fmt.Errorf("get dto segments: %w", err)
	}

	return &dto.ImageWithSegmentsFormations{
		Image:      image,
		Formations: dtoFormations,
		Segments:   dtoSegments,
	}, nil
}
