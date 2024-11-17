package med

type MedicalWorker struct {
	ID              uint64 `json:"id" validate:"required"`
	FirstName       string `json:"first_name" validate:"required"`
	MiddleName      string `json:"middle_name,omitempty"`
	LastName        string `json:"last_name" validate:"required"`
	MedOrganization string `json:"med_organization" validate:"required"`
	Job             string `json:"job" validate:"required"`
	IsRemoteWorker  bool   `json:"is_remote_worker" validate:"required"`
	ExpertDetails   string `json:"expert_details,omitempty"`
}

type Patient struct {
	ID            uint64 `json:"id" validate:"required"`
	FirstName     string `json:"first_name" validate:"required"`
	LastName      string `json:"last_name" validate:"required"`
	FatherName    string `json:"father_name,omitempty"`
	MedicalPolicy string `json:"medical_policy" validate:"required"`
	Email         string `json:"email" validate:"required,email"`
	IsActive      bool   `json:"is_active" validate:"required"`
}

type PatientCardList struct {
	Cards []PatientCard `json:"cards" validate:"required,dive"`
	Count uint64        `json:"count" validate:"required"`
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
	Patient       *Patient       `json:"patient" validate:"required"`
	Card          *PatientCard   `json:"card" validate:"required"`
	MedicalWorker *MedicalWorker `json:"medical_worker" validate:"required"`
}

type MedicalWorkerList struct {
	Workers []MedicalWorker `json:"workers" validate:"required,dive"`
	Count   uint64          `json:"count" validate:"required"`
}

type MedicalWorkerUpdateRequest struct {
	FirstName       string `json:"first_name" validate:"required"`
	LastName        string `json:"last_name" validate:"required"`
	MiddleName      string `json:"middle_name,omitempty"`
	MedOrganization string `json:"med_organization" validate:"required"`
	Job             string `json:"job" validate:"required"`
	IsRemoteWorker  bool   `json:"is_remote_worker" validate:"required"`
	ExpertDetails   string `json:"expert_details,omitempty"`
}

type MedicalWorkerPartialUpdateRequest struct {
	FirstName       *string `json:"first_name,omitempty"`
	LastName        *string `json:"last_name,omitempty"`
	MiddleName      *string `json:"middle_name,omitempty"`
	MedOrganization *string `json:"med_organization,omitempty"`
	Job             *string `json:"job,omitempty"`
	IsRemoteWorker  *bool   `json:"is_remote_worker,omitempty"`
	ExpertDetails   *string `json:"expert_details,omitempty"`
}

type AddMedicalWorkerRequest struct {
	FirstName       string `json:"first_name" validate:"required"`
	MiddleName      string `json:"middle_name,omitempty"`
	LastName        string `json:"last_name" validate:"required"`
	MedOrganization string `json:"med_organization" validate:"required"`
	Job             string `json:"job" validate:"required"`
	IsRemoteWorker  bool   `json:"is_remote_worker" validate:"required"`
	ExpertDetails   string `json:"expert_details,omitempty"`
}

type MedicalWorkerWithPatients struct {
	MedWorker MedicalWorker `json:"med_worker" validate:"required"`
	Patients  []PatientCard `json:"patients" validate:"required,dive"`
}
