package entity

import (
	"time"

	"uzi/internal/domain"

	"github.com/google/uuid"
)

type Uzi struct {
	Id         uuid.UUID `db:"id"`
	Projection string    `db:"projection"`
	Checked    bool      `db:"checked"`
	PatientID  uuid.UUID `db:"patient_id"`
	DeviceID   int       `db:"device_id"`
	CreateAt   time.Time `db:"create_at"`
}

func (Uzi) FromDomain(d domain.Uzi) Uzi {
	return Uzi{
		Id:         d.Id,
		Projection: d.Projection,
		Checked:    d.Checked,
		PatientID:  d.PatientID,
		DeviceID:   d.DeviceID,
		CreateAt:   d.CreateAt,
	}
}

func (d Uzi) ToDomain() domain.Uzi {
	return domain.Uzi{
		Id:         d.Id,
		Projection: d.Projection,
		Checked:    d.Checked,
		PatientID:  d.PatientID,
		DeviceID:   d.DeviceID,
		CreateAt:   d.CreateAt,
	}
}
