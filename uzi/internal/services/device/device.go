package device

import (
	"context"
	"fmt"

	"uzi/internal/domain"
	"uzi/internal/repository"
	"uzi/internal/repository/entity"
)

type Service interface {
	CreateDevice(ctx context.Context, deviceName string) (int, error)
	GetDeviceList(ctx context.Context) ([]domain.Device, error)
}

type service struct {
	dao repository.DAO
}

func New(
	dao repository.DAO,
) Service {
	return &service{
		dao: dao,
	}
}

func (s *service) CreateDevice(ctx context.Context, deviceName string) (int, error) {
	id, err := s.dao.NewDeviceQuery(ctx).CreateDevice(deviceName)
	// Здесь не нужно оборачивать в строку, так как по вызову сверху и так будет обернуто в строку
	return id, err
}

func (s *service) GetDeviceList(ctx context.Context) ([]domain.Device, error) {
	devices, err := s.dao.NewDeviceQuery(ctx).GetDeviceList()
	if err != nil {
		return nil, fmt.Errorf("get device list: %w", err)
	}

	return entity.Device{}.SliceToDomain(devices), nil
}
