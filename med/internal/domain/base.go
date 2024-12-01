package domain

import (
	"time"

	"github.com/google/uuid"
)

type Doctor struct {
	Id       uuid.UUID
	FullName string
	Org      string
	Job      string
	Desc     *string
}

type Patient struct {
	Id          uuid.UUID
	FullName    string
	Email       string
	Policy      string
	Active      bool
	Malignancy  bool
	LastUziDate *time.Time
}

type Card struct {
	DoctorID  uuid.UUID
	PatientID uuid.UUID
	Diagnosis *string
}
