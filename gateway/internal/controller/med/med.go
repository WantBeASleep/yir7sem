package med

import (
	"context"
	"yir/gateway/internal/entity/cardmodel"
	"yir/gateway/internal/entity/patientmodel"
	"yir/gateway/internal/entity/workermodel"
)

type MedService interface {
	// patients
	AddPatient(ctx context.Context, PatientInfo *patientmodel.PatientInformation) error
	UpdatePatient(ctx context.Context, PatientInfo *patientmodel.PatientInformation) error
	GetPatientList(ctx context.Context) ([]patientmodel.Patient, error)
	GetPatientInfoByID(ctx context.Context, ID uint64) (*patientmodel.PatientInformation, error)

	//workers
	GetMedWorkers(ctx context.Context, limit, offset uint64) (*workermodel.MedicalWorkerList, error)
	GetMedWorkerByID(ctx context.Context, id uint64) (*workermodel.MedicalWorker, error)

	UpdateMedWorker(ctx context.Context, id uint64, updateData *workermodel.MedicalWorkerUpdateRequest) (*workermodel.MedicalWorker, error)
	AddMedWorker(ctx context.Context, createData *workermodel.AddMedicalWorkerRequest) (*workermodel.MedicalWorker, error)
	GetPatientsByMedWorker(ctx context.Context, medWorkerID uint64) (*workermodel.MedicalWorkerWithPatients, error)

	// cards
	GetCards(ctx context.Context, limit, offset uint64) (*cardmodel.PatientCardList, error)
	PostCard(ctx context.Context, card *cardmodel.PatientInformation) error
	GetCardByID(ctx context.Context, id uint64) (*cardmodel.PatientCard, error)
	PutCard(ctx context.Context, card *cardmodel.PatientCard) (*cardmodel.PatientCard, error)
	DeleteCard(ctx context.Context, id uint64) error
}

type MedController struct {
	Service MedService
}

func New(s MedService) *MedController {
	return &MedController{
		Service: s,
	}
}

// функции разнесены по своим доменам
