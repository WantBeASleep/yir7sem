package models

type PatientInfo struct {
	ID            uint64 `gorm:"primaryKey"`
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
