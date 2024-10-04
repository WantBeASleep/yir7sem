package entity

type Patient struct {
	ID            uint64
	FirstName     string
	LastName      string
	FatherName    string
	MedicalPolicy string
	Email         string
	IsActive      bool
}

type PatientCard struct {
	ID              uint64
	AppointmentTime string
	HasNodules      bool
	Diagnosis       string
	MedWorkerID     uint64
	PatientID       uint64
}

type PatientInformation struct {
	Patient *Patient
	Card    *PatientCard
}
