package models

type PatientCardInfo struct {
	ID              uint64 `gorm:"primaryKey"`
	AppointmentTime string
	HasNodules      bool
	Diagnosis       string
	MedWorkerID     uint64 `gorm:"foreignKey:MedWorkerID"`
	PatientID       uint64 `gorm:"foreignKey:PatientID"`
}

func (PatientCardInfo) TableName() string {
	return PatientsCardsName
}
