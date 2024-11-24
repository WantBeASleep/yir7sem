package entity

import "github.com/google/uuid"

type PatientCard struct {
	ID              uuid.UUID
	AppointmentTime string
	HasNodules      bool
	Diagnosis       string
	MedWorkerID     uuid.UUID
	PatientID       uuid.UUID
}

type MedicalWorker struct {
	ID              uuid.UUID
	FirstName       string
	FathersName     string
	LastName        string
	MedOrganization string
	Job             string
	IsRemoteWorker  bool
	ExpertDetails   string
}

type Patient struct {
	ID            uuid.UUID
	FirstName     string
	LastName      string
	FatherName    string
	MedicalPolicy string
	Email         string
	IsActive      bool
}

type PatientInformation struct {
	Patient *Patient
}

type PatientCardList struct {
	Cards []PatientCard
	Count int //Кол-во карт
}

type MedicalWorkerUpdateRequest struct {
	FirstName       string
	LastName        string
	FathersName     string
	MedOrganization string
	Job             string
	IsRemoteWorker  bool
	ExpertDetails   string
}

// Указатели для передачи тех полей, которые нужно обновить
type MedicalWorkerPartialUpdateRequest struct {
	FirstName       *string
	LastName        *string
	FathersName     *string
	MedOrganization *string
	Job             *string
	IsRemoteWorker  *bool
	ExpertDetails   *string
}

type AddMedicalWorkerRequest struct {
	FirstName       string
	FathersName     string
	LastName        string
	MedOrganization string
	Job             string
	IsRemoteWorker  bool
	ExpertDetails   string
}

type MedicalWorkerList struct {
	Workers []MedicalWorker
	Count   int //Кол-во работников
}
