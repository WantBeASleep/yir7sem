package models

import "github.com/google/uuid"

type PatientCardInfo struct {
	ID              uuid.UUID `gorm:"primaryKey"`
	AppointmentTime string
	HasNodules      bool
	Diagnosis       string
	MedWorkerID     uuid.UUID `gorm:"foreignKey:MedWorkerID"`
	PatientID       uuid.UUID `gorm:"foreignKey:PatientID"`
}

func (PatientCardInfo) TableName() string {
	return PatientsCardsName
}
