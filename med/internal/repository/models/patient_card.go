package models

type PatientCard struct {
	ID              uint64 `gorm:"primaryKey"`
	AppointmentTime string
	HasNodules      bool
	Diagnosis       string
	MedWorkerID     uint64 `gorm:"foreignKey:MedWorkerID"`
	PatientID       uint64 `gorm:"foreignKey:PatientID"`
}

func (PatientCard) TableName() string {
	return PatientsCardsName
}
