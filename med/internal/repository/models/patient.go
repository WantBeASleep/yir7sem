package models

import "github.com/google/uuid"

type PatientInfo struct {
	ID            uuid.UUID `gorm:"primaryKey"`
	FirstName     string
	LastName      string
	FatherName    string
	MedicalPolicy string `gorm:"unique"`
	Email         string `gorm:"unique"`
	IsActive      bool
}

func (PatientInfo) TableName() string {
	return PatientsName
}
