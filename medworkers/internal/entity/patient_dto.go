package entity

type PatientCardDTO struct {
	ID              uint64
	AppointmentTime string
	HasNodules      bool
	Diagnosis       string
	Patient         PatientDTO
}

type PatientDTO struct {
	ID            uint64
	FirstName     string
	LastName      string
	FatherName    string
	MedicalPolicy string
	Email         string
	IsActive      bool
}
