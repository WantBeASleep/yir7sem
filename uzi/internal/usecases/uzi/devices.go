package uzi

import (
	"context"
	"fmt"
	"yir/uzi/internal/entity"

	"go.uber.org/zap"
)

// Прикрутить сюда пагинацию позже
func (u *UziUseCase) DeviceList(ctx context.Context) ([]*entity.Device, error) {
	u.logger.Info("[Request] Get device list")
	devices, err := u.uziRepo.GetDevicesList(ctx)
	if err != nil {
		u.logger.Error("Get device list", zap.Error(err))
		return nil, fmt.Errorf("get device list: %w", err)
	}
	u.logger.Info("[Response] Get device list", zap.Any("Devices", devices))

	return devices, nil
}
