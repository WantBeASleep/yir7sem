package mapper

import (
	"service/all/internal/entity"
	"service/all/internal/repository/models"
)

func PatientCardToEntity(pc *models.PatientCardInfo) (*entity.PatientCard, error) {
	return &entity.PatientCard{
		ID:              pc.ID,
		AppointmentTime: pc.AppointmentTime,
		HasNodules:      pc.HasNodules,
		Diagnosis:       pc.Diagnosis,
		MedWorkerID:     pc.MedWorkerID,
		PatientID:       pc.PatientID,
	}, nil
}

func PatientCardToModels(pc *entity.PatientCard) *models.PatientCardInfo {
	return &models.PatientCardInfo{
		ID:              pc.ID,
		AppointmentTime: pc.AppointmentTime,
		HasNodules:      pc.HasNodules,
		Diagnosis:       pc.Diagnosis,
		MedWorkerID:     pc.MedWorkerID,
		PatientID:       pc.PatientID,
	}
}

func PatientToEntity(p *models.PatientInfo) *entity.Patient {
	return &entity.Patient{
		ID:            p.ID,
		FirstName:     p.FirstName,
		LastName:      p.LastName,
		FatherName:    p.FatherName,
		MedicalPolicy: p.MedicalPolicy,
		Email:         p.Email,
		IsActive:      p.IsActive,
	}
}

func PatientToModels(p *entity.Patient) *models.PatientInfo {
	return &models.PatientInfo{
		ID:            p.ID,
		FirstName:     p.FirstName,
		LastName:      p.LastName,
		FatherName:    p.FatherName,
		MedicalPolicy: p.MedicalPolicy,
		Email:         p.Email,
		IsActive:      p.IsActive,
	}
}

// func MedicalWorkerToEntity(m *models.MedWorkerInfo) *entity.MedicalWorker {
// 	return &entity.MedicalWorker{
// 		ID:              m.ID,
// 		FirstName:       m.FirstName,
// 		MiddleName:      m.MiddleName,
// 		LastName:        m.LastName,
// 		MedOrganization: m.MedOrganization,
// 		Job:             m.Job,
// 		IsRemoteWorker:  m.IsRemoteWorker,
// 		ExpertDetails:   m.ExpertDetails,
// 	}
// }

//	func MedicalWorkerToModels(m *entity.MedicalWorker) *models.MedWorkerInfo {
//		return &models.MedWorkerInfo{
//			ID:              m.ID,
//			FirstName:       m.FirstName,
//			MiddleName:      m.MiddleName,
//			LastName:        m.LastName,
//			MedOrganization: m.MedOrganization,
//			Job:             m.Job,
//			IsRemoteWorker:  m.IsRemoteWorker,
//			ExpertDetails:   m.ExpertDetails,
//		}
//	}
func ToMedWorkerModel(e *entity.MedicalWorker) (*models.MedWorkerInfo, error) {
	return &models.MedWorkerInfo{
		ID:              e.ID,
		FirstName:       e.FirstName,
		MiddleName:      e.MiddleName,
		LastName:        e.LastName,
		MedOrganization: e.MedOrganization,
		Job:             e.Job,
		IsRemoteWorker:  e.IsRemoteWorker,
		ExpertDetails:   e.ExpertDetails,
	}, nil
}

func ToMedWorkerEntity(m *models.MedWorkerInfo) (*entity.MedicalWorker, error) {
	return &entity.MedicalWorker{
		ID:              m.ID,
		FirstName:       m.FirstName,
		MiddleName:      m.MiddleName,
		LastName:        m.LastName,
		MedOrganization: m.MedOrganization,
		Job:             m.Job,
		IsRemoteWorker:  m.IsRemoteWorker,
		ExpertDetails:   m.ExpertDetails,
	}, nil
}

//Доп мапперы в виде костыля
