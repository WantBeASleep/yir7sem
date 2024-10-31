package uzi

import (
	"context"
	"fmt"
	"yir/uzi/internal/entity"
	"yir/uzi/internal/entity/dto"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (u *UziUseCase) CreateUzi(ctx context.Context, req *entity.Uzi) (uuid.UUID, error) {
	u.logger.Debug("[Request] Create Uzi", zap.Any("Data", req))
	uziID, err := u.uziRepo.CreateUzi(ctx, req)
	if err != nil {
		u.logger.Error("Create Uzi", zap.Error(err))
		return fmt.Errorf("create uzi: %w", err)
	}
	u.logger.Debug("[Response] Created Uzi")
	
	return uziID, nil
}

func (u *UziUseCase) GetUzi(ctx context.Context, id uuid.UUID) (*dto.Uzi, error) {
	u.logger.Debug("[Request] Get Uzi", zap.Any("uzi id", id))
	uzi, err := u.uziRepo.GetUziByID(ctx, id)
	if err != nil {
		u.logger.Error("Get uzi", zap.Error(err))
		return nil, fmt.Errorf("get uzi: %w", err)
	}
	u.logger.Debug("[Response] Got uzi", zap.Any("Uzi", uzi))

	u.logger.Debug("[Request] Get uzi images", zap.Any("uzi id", id))
	images, err := u.uziRepo.GetUziImages(ctx, id)
	if err != nil {
		u.logger.Error("Get uzi images", zap.Error(err))
		return nil, fmt.Errorf("get uzi images: %w", err)
	}
	u.logger.Debug("[Response] Got uzi images", zap.Any("Images", images))

	u.logger.Debug("[Request] Get uzi formations", zap.Any("uzi id", id))
	formations, err := u.uziRepo.GetUziFormations(ctx, id)
	if err != nil {
		u.logger.Error("Get uzi formations", zap.Error(err))
		return nil, fmt.Errorf("get uzi formations: %w", err)
	}
	u.logger.Debug("[Response] Got uzi formations", zap.Any("Formations", formations))

	u.logger.Debug("[Request] Get uzi segments", zap.Any("uzi id", id))
	segments, err := u.uziRepo.GetUziSegments(ctx, id)
	if err != nil {
		u.logger.Error("Get uzi segments", zap.Error(err))
		return nil, fmt.Errorf("get uzi segments: %w", err)
	}
	u.logger.Debug("[Response] Got uzi segments", zap.Any("Segments", segments))

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
