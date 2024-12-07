package repository

import (
	"github.com/WantBeASleep/goooool/daolib"

	"uzi/internal/repository/entity"
)

const deviceTable = "device"

type DeviceQuery interface {
	GetDeviceList() ([]entity.Device, error)
}

type deviceQuery struct {
	*daolib.BaseQuery
}

func (q *deviceQuery) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}

func (q *deviceQuery) GetDeviceList() ([]entity.Device, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"name",
		).
		From(deviceTable)

	var devices []entity.Device
	if err := q.Runner().Selectx(q.Context(), &devices, query); err != nil {
		return nil, err
	}

	return devices, nil
}
