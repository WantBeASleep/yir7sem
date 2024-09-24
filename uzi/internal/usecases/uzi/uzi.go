package uzi

import (
	"context"
	"fmt"
	"yir/uzi/internal/entity"

	"go.uber.org/zap"
)

func (u *UziUseCase) InsertUzi(ctx context.Context, req *entity.InsertUziRequest) error {
	u.logger.Debug("Insert Uzi Request")

	u.logger.Debug("[Request] Insert Uzi")
	if err := u.uziRepo.InsertUzi(ctx, &req.Uzi); err != nil {
		u.logger.Error("Insert Uzi", zap.Error(err))
		return fmt.Errorf("Inser")
	}

	return nil
}
