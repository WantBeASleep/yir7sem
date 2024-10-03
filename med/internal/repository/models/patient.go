package models

type Patient struct {
	ID            uint64 `gorm:"primaryKey"`
	FirstName     string
	LastName      string
	FatherName    string
	MedicalPolicy string `gorm:"unique"`
	Email         string `gorm:"unique"`
	IsActive      bool
}

func (Patient) TableName() string {
	return PatientsName
}
