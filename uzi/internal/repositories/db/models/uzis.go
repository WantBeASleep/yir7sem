package models

const UziName = "uzis"

type Uzi struct {
	Uuid        string `gorm:"primaryKey"`
	Url         string
	Projection  string
	PatientUUID string

	DeviceID uint64  `gorm:"not null"`
	Devices  Device `gorm:"foreignKey:DeviceID;references:Id"`
}

func (Uzi) TableName() string {
	return UziName
}
