// Преобразование entities в db models и наоборот
package mappers

import (
	"service/internal/entity"
	"service/internal/repositories/models"
)

// Сущность медворкера в модель базы данных
func ToMedWorkerModel(e *entity.MedicalWorker) (*models.MedWorkerInfo, error) {
	return &models.MedWorkerInfo{
		ID:              uint64(e.ID),
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
		ID:              int(m.ID),
		FirstName:       m.FirstName,
		MiddleName:      m.MiddleName,
		LastName:        m.LastName,
		MedOrganization: m.MedOrganization,
		Job:             m.Job,
		IsRemoteWorker:  m.IsRemoteWorker,
		ExpertDetails:   m.ExpertDetails,
	}, nil
}
