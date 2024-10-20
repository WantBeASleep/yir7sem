package entity

import "github.com/google/uuid"

type Uzi struct {
	ID         uuid.UUID `json:"id" validate:"required"`
	URL        string    `json:"url" validate:"required,url"`
	Projection string    `json:"projection" validate:"required"`
	PatientID  uuid.UUID `json:"patient_id" validate:"required"`
	DeviceID   int       `json:"device_id" validate:"required"`
}

type Image struct {
	ID    uuid.UUID `json:"id" validate:"required"`
	URL   string    `json:"url" validate:"required,url"`
	Page  int       `json:"page" validate:"required"`
	UziID uuid.UUID `json:"uzi_id" validate:"required"`
}

type Formation struct {
	ID uuid.UUID `json:"id" validate:"required"`

	// meta
	AI       bool `json:"ai" validate:"required"`
	TiradsID int  `json:"tirads_id" validate:"required"`
}

type Segment struct {
	ID          uuid.UUID `json:"id" validate:"required"`
	ImageID     uuid.UUID `json:"image_id" validate:"required"`
	FormationID uuid.UUID `json:"formation_id" validate:"required"`

	// meta
	ContorURL string `json:"contor_url" validate:"required,url"`
	TiradsID  int    `json:"tirads_id" validate:"required"`
}

type Device struct {
	ID   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type Tirads struct {
	Tirads1 float64 `json:"tirads1" validate:"required"`
	Tirads2 float64 `json:"tirads2" validate:"required"`
	Tirads3 float64 `json:"tirads3" validate:"required"`
	Tirads4 float64 `json:"tirads4" validate:"required"`
	Tirads5 float64 `json:"tirads5" validate:"required"`
}
