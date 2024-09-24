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
	Id         uuid.UUID
	Url        string
	Projection string
	PatientID  uuid.UUID

	DeviceID int
}

type Image struct {
	Id   uuid.UUID
	Url  string
	Page int

	UziID uuid.UUID
}

type Formation struct {
	Id uuid.UUID
	Ai bool

	TiradsId int
}

type ImageFormation struct {
	Id          int
	ContorURL   string
	FormationID uuid.UUID
	ImageID     uuid.UUID
	TiradsId    int
}
