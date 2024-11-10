package medworkers

import (
	"context"
	"errors"
	pb "yir/med/api"
	"yir/med/internal/entity"
	"yir/med/internal/usecase"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedMedWorkersServer

	MedWorkerUseCase usecase.MedicalWorker
	logger           *zap.Logger
}

func NewServer(MedWorkerUseCase usecase.MedicalWorker, logger *zap.Logger) *Server {
	return &Server{
		MedWorkerUseCase: MedWorkerUseCase,
		logger:           logger,
	}
}

func (s *Server) GetMedWorkers(ctx context.Context, request *pb.GetMedworkerRequest) (*pb.GetMedworkerListResponse, error) {
	s.logger.Info("[Request] New request", zap.Any("data", request))
	limit := int(request.GetLimit())
	offset := int(request.GetOffset())
	medWorkerList, err := s.MedWorkerUseCase.GetMedWorkers(ctx, limit, offset)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "Медицинские работники не найдены")
		}
		s.logger.Error("Не получилось достать медицинских работников", zap.Error(err))
		return nil, status.Error(codes.Internal, "Не получилось достать медицинских работников")
	}
	response := &pb.GetMedworkerListResponse{
		Count:   uint64(medWorkerList.Count),
		Results: []*pb.MedWorker{},
	}
	for _, worker := range medWorkerList.Workers {
		workerResponse := &pb.MedWorker{
			Id:              worker.ID.String(),
			FirstName:       worker.FirstName,
			FathersName:     worker.FathersName,
			LastName:        worker.LastName,
			MedOrganization: worker.MedOrganization,
			Job:             worker.Job,
			IsRemoteWorker:  worker.IsRemoteWorker,
			ExpertDetails:   worker.ExpertDetails,
		}
		response.Results = append(response.Results, workerResponse)
	}

	return response, nil
}

func (s *Server) GetMedWorkerByID(ctx context.Context, request *pb.GetMedMedWorkerByIDRequest) (*pb.GetMedWorkerByIDResponse, error) {
	s.logger.Info("Получен запрос GetMedWorkerByID", zap.Any("ID", request.Id))
	worker, err := s.MedWorkerUseCase.GetMedWorkerByID(ctx, request.Id)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			s.logger.Warn("Медицинский работник не найден", zap.Any("ID", request.Id))
			return nil, status.Error(codes.NotFound, "Медицинский работник не найден")
		}
		s.logger.Error("Не получилось достать информацию о медицинском работнике", zap.Error(err))
		return nil, status.Error(codes.Internal, "Не получилось достать информацию о медицинском работнике")
	}

	response := &pb.GetMedWorkerByIDResponse{
		Worker: &pb.MedWorker{
			Id:              worker.ID.String(),
			FirstName:       worker.FirstName,
			FathersName:     worker.FathersName,
			LastName:        worker.LastName,
			MedOrganization: worker.MedOrganization,
			Job:             worker.Job,
			IsRemoteWorker:  worker.IsRemoteWorker,
			ExpertDetails:   worker.ExpertDetails,
		},
	}

	return response, nil
}

func (s *Server) UpdateMedWorker(ctx context.Context, request *pb.UpdateMedWorkerRequest) (*pb.UpdateMedWorkerResponse, error) {
	s.logger.Info("Получен запрос UpdateMedWorker", zap.Any("ID", request.Id))

	updateData := &entity.MedicalWorkerUpdateRequest{
		FirstName:       request.FirstName,
		FathersName:     request.FathersName,
		LastName:        request.LastName,
		MedOrganization: request.MedOrganization,
		Job:             request.Job,
		IsRemoteWorker:  request.IsRemoteWorker,
		ExpertDetails:   request.ExpertDetails,
	}

	updatedWorker, err := s.MedWorkerUseCase.UpdateMedWorker(ctx, request.Id, updateData)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			s.logger.Warn("Медицинский работник не найден для обновления", zap.Any("ID", request.Id))
			return nil, status.Error(codes.NotFound, "медицинский работник не найден")
		}
		s.logger.Error("Не получилось обновить информацию о медицинском работнике", zap.Error(err))
		return nil, status.Error(codes.Internal, "не получилось обновить информацию о медицинском работнике")
	}

	response := &pb.UpdateMedWorkerResponse{
		Worker: &pb.MedWorker{
			Id:              updatedWorker.ID.String(),
			FirstName:       updatedWorker.FirstName,
			FathersName:     updatedWorker.FathersName,
			LastName:        updatedWorker.LastName,
			MedOrganization: updatedWorker.MedOrganization,
			Job:             updatedWorker.Job,
			IsRemoteWorker:  updatedWorker.IsRemoteWorker,
			ExpertDetails:   updatedWorker.ExpertDetails,
		},
	}

	return response, nil
}

func (s *Server) AddMedWorker(ctx context.Context, request *pb.AddMedWorkerRequest) (*pb.AddMedWorkerResponse, error) {
	s.logger.Info("Received AddMedWorker request", zap.Any("request", request))

	medworkerReq := &entity.AddMedicalWorkerRequest{
		FirstName:       request.GetFirstName(),
		FathersName:     request.GetFathersName(),
		LastName:        request.GetLastName(),
		MedOrganization: request.GetMedOrganization(),
		Job:             request.GetJob(),
		IsRemoteWorker:  request.GetIsRemoteWorker(),
		ExpertDetails:   request.GetExpertDetails(),
	}

	medworker, err := s.MedWorkerUseCase.AddMedWorker(ctx, medworkerReq)
	if err != nil {
		s.logger.Error("Failed to add medworker", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to add medworker: %v", err)
	}

	response := &pb.AddMedWorkerResponse{
		Worker: &pb.MedWorker{
			Id:              medworker.ID.String(),
			FirstName:       medworker.FirstName,
			FathersName:     medworker.FathersName,
			LastName:        medworker.LastName,
			MedOrganization: medworker.MedOrganization,
			Job:             medworker.Job,
			IsRemoteWorker:  medworker.IsRemoteWorker,
			ExpertDetails:   medworker.ExpertDetails,
		},
	}

	s.logger.Info("Medworker successfully added", zap.Any("worker", response.Worker))
	return response, nil
}

func (s *Server) GetPatientsByMedWorker(ctx context.Context, request *pb.GetPatientsByMedWorkerRequest) (*pb.GetPatientsByMedWorkerResponse, error) {
	s.logger.Info("[Request] Get patients by med worker", zap.Any("data", request))

	cards, medWorkerID, err := s.MedWorkerUseCase.GetPatientsByMedWorker(ctx, request.MedWorkerId)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "No patients found for the specified med worker")
		}
		s.logger.Error("Failed to fetch patients by med worker", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to fetch patients by med worker")
	}

	response := &pb.GetPatientsByMedWorkerResponse{
		MedWorkerId: string(medWorkerID),
		Cards:       []*pb.Card{},
	}

	for _, card := range cards {
		cardResponse := &pb.Card{
			Id:              card.ID.String(),
			AppointmentTime: card.AppointmentTime,
			HasNodules:      card.HasNodules,
			Diagnosis:       card.Diagnosis,
			MedWorkerId:     card.MedWorkerID.String(),
			Patient: &pb.Patient{
				Id: card.PatientID.String(),
				// FirstName:     card.Patient.FirstName,  //эт подумать - оставлять полностью или только айдишник
				// LastName:      card.Patient.LastName,
				// FatherName:    card.Patient.FatherName,
				// MedicalPolicy: card.Patient.MedicalPolicy,
				// Email:         card.Patient.Email,
				// IsActive:      card.Patient.IsActive,
			},
		}
		response.Cards = append(response.Cards, cardResponse)
	}

	return response, nil
}

// func (s *Server) GetPatientsByMedWorker(ctx context.Context, req *pb.GetPatientsByMedWorkerRequest) (*pb.GetPatientsByMedWorkerResponse, error) {
// 	s.logger.Info("Received request for patients", zap.Any("medWorkerId", req.GetMedWorkerId()))

// 	// Получаем медработника
// 	medWorker, err := s.MedWorkerUseCase.GetMedWorkerByID(ctx, int(req.GetMedWorkerId()))
// 	if err != nil {
// 		s.logger.Error("Failed to get medworker", zap.Error(err))
// 		return nil, status.Errorf(codes.Internal, "failed to get medworker: %v", err)
// 	}

// 	// Вызов сервиса пациентов через gRPC/HTTP для получения пациентов
// 	patients, err := s.PatientClient.GetPatientsByMedWorkerID(ctx, req.GetMedWorkerId())
// 	if err != nil {
// 		s.logger.Error("Failed to get patients from patient service", zap.Error(err))
// 		return nil, status.Errorf(codes.Internal, "failed to get patients: %v", err)
// 	}

// 	// Формируем ответ
// 	response := &pb.GetPatientsByMedWorkerResponse{
// 		MedWorker: &pb.MedWorker{
// 			Id:              uint64(medWorker.ID),
// 			LastName:        medWorker.LastName,
// 			FirstName:       medWorker.FirstName,
// 			FathersName:      medWorker.FathersName,
// 			MedOrganization: medWorker.MedOrganization,
// 			Job:             medWorker.Job,
// 			IsRemoteWorker:  medWorker.IsRemoteWorker,
// 			ExpertDetails:   medWorker.ExpertDetails,
// 		},
// 		Cards: []*pb.PatientCard{},
// 	}

// 	for _, patientCard := range patients {
// 		response.Cards = append(response.Cards, &pb.PatientCard{
// 			Id:              uint64(patientCard.ID),
// 			AppointmentTime: patientCard.AppointmentTime.String(),
// 			HasNodules:      patientCard.HasNodules,
// 			Diagnosis:       patientCard.Diagnosis,
// 			Patient: &pb.Patient{
// 				Id:            uint64(patientCard.Patient.ID),
// 				FirstName:     patientCard.Patient.FirstName,
// 				LastName:      patientCard.Patient.LastName,
// 				FatherName:    patientCard.Patient.FatherName,
// 				MedicalPolicy: patientCard.Patient.MedicalPolicy,
// 				Email:         patientCard.Patient.Email,
// 				IsActive:      patientCard.Patient.IsActive,
// 			},
// 		})
// 	}

// 	return response, nil
// }

// func (s *Server) GetPatientsByMedWorker(ctx context.Context, req *pb.GetPatientsByMedWorkerRequest) (*pb.GetPatientsByMedWorkerResponse, error) {
// 	s.logger.Info("Received request for patients", zap.Uint64("med_worker_id", req.MedWorkerId))

// 	// Вызов usecase для получения пациентов по id врача
// 	result, err := s.MedWorkerUseCase.GetPatientsByMedWorker(ctx, req.MedWorkerId)
// 	if err != nil {
// 		s.logger.Error("Failed to get patients", zap.Error(err))
// 		return nil, status.Error(codes.Internal, "failed to get patients")
// 	}

// 	// Формирование ответа
// 	response := &pb.GetPatientsByMedWorkerResponse{
// 		MedWorker: &pb.MedWorker{
// 			Id:              uint64(result.MedWorker.ID),
// 			FirstName:       result.MedWorker.FirstName,
// 			FathersName:      result.MedWorker.FathersName,
// 			LastName:        result.MedWorker.LastName,
// 			MedOrganization: result.MedWorker.MedOrganization,
// 			Job:             result.MedWorker.Job,
// 			IsRemoteWorker:  result.MedWorker.IsRemoteWorker,
// 			ExpertDetails:   result.MedWorker.ExpertDetails,
// 		},
// 	}

// 	for _, patientCard := range result.Patients {
// 		response.Cards = append(response.Cards, &pb.PatientCard{
// 			Id:              patientCard.ID,
// 			AppointmentTime: patientCard.AppointmentTime,
// 			HasNodules:      patientCard.HasNodules,
// 			Diagnosis:       patientCard.Diagnosis,
// 			Patient: &pb.Patient{
// 				Id:            patientCard.Patient.ID,
// 				FirstName:     patientCard.Patient.FirstName,
// 				LastName:      patientCard.Patient.LastName,
// 				FatherName:    patientCard.Patient.FatherName,
// 				MedicalPolicy: patientCard.Patient.MedicalPolicy,
// 				Email:         patientCard.Patient.Email,
// 				IsActive:      patientCard.Patient.IsActive,
// 			},
// 		})
// 	}

// 	return response, nil
// }
