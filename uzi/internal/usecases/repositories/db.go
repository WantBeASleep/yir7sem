package repositories

import (
	"context"
	"yir/uzi/internal/entity"
)

type UziRepo interface {
	GetDevicesList(ctx context.Context) ([]*entity.Device, error)
	InsertUzi(ctx context.Context, uzi *entity.Uzi) error
}
