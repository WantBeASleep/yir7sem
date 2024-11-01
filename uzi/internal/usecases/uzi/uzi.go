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
		return uuid.Nil, fmt.Errorf("create uzi: %w", err)
	}
	u.logger.Debug("[Response] Created Uzi")
	
	return uziID, nil
}

func (u *UziUseCase) GetUzi(ctx context.Context, id uuid.UUID) (*entity.Uzi, error) {
	u.logger.Debug("[Request] Get Uzi", zap.String("uzi id", id.String()))
	uzi, err := u.uziRepo.GetUziByID(ctx, id)
	if err != nil {
		u.logger.Error("Get uzi", zap.Error(err))
		return nil, fmt.Errorf("get uzi: %w", err)
	}
	u.logger.Debug("[Response] Got uzi", zap.Any("Uzi", uzi))

	return uzi, nil
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
