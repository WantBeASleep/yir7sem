package mapper

import (
	"yir/med/internal/entity"
	"yir/med/internal/repository/models"

	"github.com/google/uuid"
)

func PatientToEntity(p *models.Patient) *entity.Patient {
	uuid, _ := uuid.Parse(p.UUID)
	return &entity.Patient{
		UUID:          uuid,
		FirstName:     p.FirstName,
		LastName:      p.LastName,
		FatherName:    p.FatherName,
		MedicalPolicy: p.MedicalPolicy,
		Email:         p.Email,
		IsActive:      p.IsActive,
	}
}

func PatientToModel(p *entity.Patient) *models.Patient {
	return &models.Patient{
		UUID:          p.UUID.String(),
		FirstName:     p.FirstName,
		LastName:      p.LastName,
		FatherName:    p.FatherName,
		MedicalPolicy: p.MedicalPolicy,
		Email:         p.Email,
		IsActive:      p.IsActive,
	}
}

// func PatientCardToEntity(pc *models.PatientCard) *entity.PatientCard {
// 	return &entity.PatientCard{
// 		ID:              pc.ID,
// 		AppointmentTime: pc.AppointmentTime,
// 		HasNodules:      pc.HasNodules,
// 		Diagnosis:       pc.Diagnosis,
// 		MedWorkerID:     pc.MedWorkerID,
// 		PatientID:       pc.PatientID,
// 	}
// }

// func PatientCardToModel(pc *entity.PatientCard) *models.PatientCard {
// 	return &models.PatientCard{
// 		ID:              pc.ID,
// 		AppointmentTime: pc.AppointmentTime,
// 		HasNodules:      pc.HasNodules,
// 		Diagnosis:       pc.Diagnosis,
// 		MedWorkerID:     pc.MedWorkerID,
// 		PatientID:       pc.PatientID,
// 	}
// }
