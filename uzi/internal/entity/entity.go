package entity

import (
	"github.com/google/uuid"
)

// Не увидел смысла декомпозировать в отдельные таблички meta инфы

type Uzi struct {
	Id  uuid.UUID
	Url string

	// meta
	Projection string
	PatientID  uuid.UUID
	DeviceID   int
}

type Image struct {
	Id    uuid.UUID
	Url   string
	UziID uuid.UUID

	Page int
}

type Formation struct {
	Id uuid.UUID

	// meta
	Ai       bool
	TiradsID int
}

type Segment struct {
	Id          int
	ImageID     uuid.UUID
	FormationID uuid.UUID

	// meta
	ContorURL string
	TiradsID  int
}

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
