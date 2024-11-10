package workermodel

import (
	"yir/gateway/internal/entity/patientmodel"
)

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
	MedWorker MedicalWorker              `json:"med_worker" validate:"required"`
	Patients  []patientmodel.PatientCard `json:"patients" validate:"required,dive"`
}
