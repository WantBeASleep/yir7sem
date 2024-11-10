package patientmodel

type Patient struct {
	ID            uint64 `json:"id" validate:"required"`
	FirstName     string `json:"first_name" validate:"required"`
	LastName      string `json:"last_name" validate:"required"`
	FatherName    string `json:"father_name" validate:"required"`
	MedicalPolicy string `json:"medical_policy" validate:"required"`
	Email         string `json:"email" validate:"required,email"`
	IsActive      bool   `json:"is_active" validate:"required"`
}

type PatientCard struct {
	ID              uint64 `json:"id" validate:"required"`
	AppointmentTime string `json:"appointment_time" validate:"required,datetime"`
	HasNodules      bool   `json:"has_nodules" validate:"required"`
	Diagnosis       string `json:"diagnosis" validate:"required"`
	MedWorkerID     uint64 `json:"med_worker_id" validate:"required"`
	PatientID       uint64 `json:"patient_id" validate:"required"`
}

type PatientInformation struct {
	Patient *Patient     `json:"patient" validate:"required"`
	Card    *PatientCard `json:"patient_card" validate:"required"`
}
