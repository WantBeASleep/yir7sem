package models

import "github.com/google/uuid"

const UziName = "uzis"

type Uzi struct {
	Id         uuid.UUID `gorm:"primaryKey"`
	Url        string
	Projection string
	PatientID  uuid.UUID

	DeviceID int    `gorm:"not null"`
	Devices  Device `gorm:"foreignKey:DeviceID;references:Id" copier:"-"`
}

func (Uzi) TableName() string {
	return UziName
}
