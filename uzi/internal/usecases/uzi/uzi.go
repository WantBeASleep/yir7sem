package uzi

import (
	"context"
	"fmt"
	"yir/uzi/internal/entity"
	"yir/uzi/internal/entity/dto"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (u *UziUseCase) InsertUzi(ctx context.Context, req *dto.Uzi) error {
	u.logger.Debug("[Request] Insert Uzi", zap.Any("Data", req.UziInfo))
	if err := u.uziRepo.InsertUzi(ctx, req.UziInfo); err != nil {
		u.logger.Error("Insert Uzi", zap.Error(err))
		return fmt.Errorf("insert uzi: %w", err)
	}
	u.logger.Debug("[Response] Inserted Uzi")

	u.logger.Debug("[Request] Insert images")
	if err := u.uziRepo.InsertImages(ctx, req.Images); err != nil {
		u.logger.Error("Insert images", zap.Error(err))
		return fmt.Errorf("insert images: %w", err)
	}
	u.logger.Debug("[Response] Inserted images")

	if err := u.InsertDTOFormations(ctx, req.Formations); err != nil {
		return fmt.Errorf("insert dto formations: %w", err)
	}

	if err := u.InsertDTOSegments(ctx, req.Segments); err != nil {
		return fmt.Errorf("insert dto segments: %w", err)
	}

	return nil
}

func (u *UziUseCase) GetUzi(ctx context.Context, id uuid.UUID) (*dto.Uzi, error) {
	u.logger.Debug("[Request] Get Uzi", zap.Any("uzi id", id))
	uzi, err := u.uziRepo.GetUziByID(ctx, id)
	if err != nil {
		u.logger.Error("Get uzi", zap.Error(err))
		return nil, fmt.Errorf("get uzi: %w", err)
	}
	u.logger.Debug("[Response] Get uzi", zap.Any("Uzi", uzi))

	u.logger.Debug("[Request] Get uzi images", zap.Any("uzi id", id))
	images, err := u.uziRepo.GetUziImages(ctx, id)
	if err != nil {
		u.logger.Error("Get uzi images", zap.Error(err))
		return nil, fmt.Errorf("get uzi images: %w", err)
	}
	u.logger.Debug("[Response] Get uzi images", zap.Any("Images", images))

	u.logger.Debug("[Request] Get uzi formations", zap.Any("uzi id", id))
	formations, err := u.uziRepo.GetUziFormations(ctx, id)
	if err != nil {
		u.logger.Error("Get uzi formations", zap.Error(err))
		return nil, fmt.Errorf("get uzi formations: %w", err)
	}
	u.logger.Debug("[Response] Get uzi formations", zap.Any("Formations", formations))

	u.logger.Debug("[Request] Get uzi segments", zap.Any("uzi id", id))
	segments, err := u.uziRepo.GetUziSegments(ctx, id)
	if err != nil {
		u.logger.Error("Get uzi segments", zap.Error(err))
		return nil, fmt.Errorf("get uzi segments: %w", err)
	}
	u.logger.Debug("[Response] Get uzi segments", zap.Any("Segments", segments))

	dtoFormations, err := u.GetDTOFormation(ctx, formations)
	if err != nil {
		return nil, fmt.Errorf("get dto formations: %w", err)
	}

	dtoSegments, err := u.GetDTOSegments(ctx, segments)
	if err != nil {
		return nil, fmt.Errorf("get dto segments: %w", err)
	}

	return &dto.Uzi{
		UziInfo:    uzi,
		Images:     images,
		Formations: dtoFormations,
		Segments:   dtoSegments,
	}, nil
}

func (u *UziUseCase) GetUziInfo(ctx context.Context, id uuid.UUID) (*entity.Uzi, error) {
	u.logger.Debug("[Request] Get Uzi", zap.Any("uzi id", id))
	uzi, err := u.uziRepo.GetUziByID(ctx, id)
	if err != nil {
		u.logger.Error("Get uzi", zap.Error(err))
		return nil, fmt.Errorf("get uzi: %w", err)
	}
	u.logger.Debug("[Response] Get uzi", zap.Any("Uzi", uzi))

	return uzi, nil
}

func (u *UziUseCase) UpdateUziInfo(ctx context.Context, id uuid.UUID, req *entity.Uzi) error {
	u.logger.Debug("[Request] Update UziInfo", zap.Any("Requset", req))
	if err := u.uziRepo.UpdateUzi(ctx, id, req); err != nil {
		u.logger.Error("Update UziInfo", zap.Error(err))
		return fmt.Errorf("update uzi info: %w", err)
	}
	u.logger.Debug("[Response] Updated Uzi")

	return nil
}
