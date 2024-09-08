package entity

import (
	"github.com/google/uuid"
)

type Device struct {
	Id   int
	Name string
}

type Tirads struct {
	Tirads1 float64
	Tirads2 float64
	Tirads3 float64
	Tirads4 float64
	Tirads5 float64
}

type Uzi struct {
	Uuid        uuid.UUID
	Url         string
	Projection  string
	PatientUUID uuid.UUID

	DeviceID int
}

type Image struct {
	Uuid uuid.UUID
	Url  string
	Page int

	UziUUID uuid.UUID
}
