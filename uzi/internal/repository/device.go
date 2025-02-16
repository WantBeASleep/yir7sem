package repository

import (
	"uzi/internal/repository/entity"

	daolib "github.com/WantBeASleep/med_ml_lib/dao"
)

const deviceTable = "device"

type DeviceQuery interface {
	// ID device будет сгенерирован
	CreateDevice(name string) (int, error)
	GetDeviceList() ([]entity.Device, error)
}

type deviceQuery struct {
	*daolib.BaseQuery
}

func (q *deviceQuery) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}

func (q *deviceQuery) CreateDevice(name string) (int, error) {
	query := q.QueryBuilder().
		Insert(deviceTable).
		Columns(
			"name",
		).
		Values(
			name,
		).
		Suffix("RETURNING id")

	var id int
	if err := q.Runner().Getx(q.Context(), &id, query); err != nil {
		return 0, err
	}

	return id, nil
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
