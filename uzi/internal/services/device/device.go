package device

import (
	"context"
	"fmt"

	"yirv2/uzi/internal/domain"
	"yirv2/uzi/internal/repository"
	"yirv2/uzi/internal/repository/entity"
)

type Service interface {
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

func (s *service) GetDeviceList(ctx context.Context) ([]domain.Device, error) {
	devices, err := s.dao.NewDeviceQuery(ctx).GetDeviceList()
	if err != nil {
		return nil, fmt.Errorf("get device list: %w", err)
	}

	return entity.Device{}.SliceToDomain(devices), nil
}
