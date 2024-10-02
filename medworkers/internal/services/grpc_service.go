package services

import (
	"context"
	"log"

	"service/internal/entity"
	pb "service/med/api/patient" // Импортируем gRPC клиент для пациентов

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCPatientService struct {
	client pb.MedPatientClient
}

// NewGRPCPatientService создает новый gRPC клиент для сервиса пациентов
func NewGRPCPatientService(conn *grpc.ClientConn) *GRPCPatientService {
	return &GRPCPatientService{
		client: pb.NewMedPatientClient(conn),
	}
}

// GetPatientsByMedWorker вызывает gRPC метод сервиса пациентов для получения всех пациентов,
// затем фильтрует их по ID медработника через поле `med_worker_id` в `PatientCard`
func (s *GRPCPatientService) GetPatientsByMedWorker(ctx context.Context, medWorkerID uint64) ([]*entity.PatientCardDTO, error) {
	req := &emptypb.Empty{} // Используем пустой запрос для метода GetPatientList

	// Получаем список пациентов с сервиса пациентов
	res, err := s.client.GetPatientList(ctx, req)
	if err != nil {
		log.Printf("Error calling GetPatientList: %v", err)
		return nil, err
	}

	// Преобразуем полученные данные и фильтруем пациентов по ID медработника через поле `med_worker_id` в карточке пациента
	patientCards := []*entity.PatientCardDTO{}
	for _, patient := range res.Patients {
		// Для каждого пациента вызываем метод GetPatientInfoByID для получения карты пациента
		patientInfoReq := &pb.PatientInfoRequest{Id: patient.Id}
		patientInfoRes, err := s.client.GetPatientInfoByID(ctx, patientInfoReq)
		if err != nil {
			log.Printf("Error calling GetPatientInfoByID: %v", err)
			continue
		}

		// Проверяем, прикреплен ли пациент к указанному медработнику
		if patientInfoRes.PatientCard.MedWorkerId == medWorkerID {
			patientCards = append(patientCards, &entity.PatientCardDTO{
				ID:              patientInfoRes.PatientCard.Id,
				AppointmentTime: patientInfoRes.PatientCard.AppointmentTime,
				Diagnosis:       patientInfoRes.PatientCard.Diagnosis,
				HasNodules:      patientInfoRes.PatientCard.HasNodules,
				Patient: entity.PatientDTO{
					ID:            patientInfoRes.Patient.Id,
					FirstName:     patientInfoRes.Patient.FirstName,
					LastName:      patientInfoRes.Patient.LastName,
					FatherName:    patientInfoRes.Patient.FatherName,
					MedicalPolicy: patientInfoRes.Patient.MedicalPolicy,
					Email:         patientInfoRes.Patient.Email,
					IsActive:      patientInfoRes.Patient.IsActive,
				},
			})
		}
	}

	return patientCards, nil
}
