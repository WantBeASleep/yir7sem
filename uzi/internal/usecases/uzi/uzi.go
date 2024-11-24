package uzi

import (
	"context"
	"fmt"
	"yir/uzi/internal/entity"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (u *UziUseCase) CreateUzi(ctx context.Context, req *entity.Uzi) (uuid.UUID, error) {
	u.logger.Debug("[Request] Create Uzi", zap.Any("Data", req))
	uziID := uuid.New()
	req.Id = uziID

	if err := u.uziRepo.InsertUzi(ctx, req); err != nil {
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

func (u *UziUseCase) UpdateUzi(ctx context.Context, id uuid.UUID, req *entity.Uzi) (*entity.Uzi, error) {
	u.logger.Debug("[Request] Update Uzi", zap.String("uzi id", id.String()))
	uzi, err := u.uziRepo.UpdateUzi(ctx, id, req)
	if err != nil {
		u.logger.Error("Update Uzi", zap.Error(err))
		return nil, fmt.Errorf("update uzi info: %w", err)
	}
	u.logger.Debug("[Response] Updated Uzi")

	return uzi, nil
}
