package entity

import "uzi/internal/domain"

type Device struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

func (Device) FromDomain(d domain.Device) Device {
	return Device{
		Id:   d.Id,
		Name: d.Name,
	}
}

func (d Device) ToDomain() domain.Device {
	return domain.Device{
		Id:   d.Id,
		Name: d.Name,
	}
}

func (Device) SliceToDomain(slice []Device) []domain.Device {
	res := make([]domain.Device, 0, len(slice))
	for _, v := range slice {
		res = append(res, v.ToDomain())
	}
	return res
}
