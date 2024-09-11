package medworkers

import (
	"context"
	"errors"
	pb "yir/medworkers/api/medworkers"
	"yir/medworkers/internal/controller/usecases"
	"yir/medworkers/internal/entity"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedMedWorkersServer

	MedWorkerUseCase usecases.MedicalWorker
	logger           *zap.Logger
}

func NewServer(MedWorkerUseCase usecases.MedicalWorker, logger *zap.Logger) *Server {
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
			Id:              uint64(worker.ID),
			FirstName:       worker.FirstName,
			MiddleName:      worker.MiddleName,
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
	s.logger.Info("Получен запрос GetMedWorkerByID", zap.Uint64("ID", request.Id))
	worker, err := s.MedWorkerUseCase.GetMedWorkerByID(ctx, int(request.Id))
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			s.logger.Warn("Медицинский работник не найден", zap.Uint64("ID", request.Id))
			return nil, status.Error(codes.NotFound, "Медицинский работник не найден")
		}
		s.logger.Error("Не получилось достать информацию о медицинском работнике", zap.Error(err))
		return nil, status.Error(codes.Internal, "Не получилось достать информацию о медицинском работнике")
	}

	response := &pb.GetMedWorkerByIDResponse{
		Worker: &pb.MedWorker{
			Id:              uint64(worker.ID),
			FirstName:       worker.FirstName,
			MiddleName:      worker.MiddleName,
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
	s.logger.Info("Получен запрос UpdateMedWorker", zap.Uint64("ID", request.Id))

	updateData := &entity.MedicalWorkerUpdateRequest{
		FirstName:       request.FirstName,
		MiddleName:      request.MiddleName,
		LastName:        request.LastName,
		MedOrganization: request.MedOrganization,
		Job:             request.Job,
		IsRemoteWorker:  request.IsRemoteWorker,
		ExpertDetails:   request.ExpertDetails,
	}

	updatedWorker, err := s.MedWorkerUseCase.UpdateMedWorker(ctx, int(request.Id), updateData)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			s.logger.Warn("Медицинский работник не найден для обновления", zap.Uint64("ID", request.Id))
			return nil, status.Error(codes.NotFound, "медицинский работник не найден")
		}
		s.logger.Error("Не получилось обновить информацию о медицинском работнике", zap.Error(err))
		return nil, status.Error(codes.Internal, "не получилось обновить информацию о медицинском работнике")
	}

	response := &pb.UpdateMedWorkerResponse{
		Worker: &pb.MedWorker{
			Id:              uint64(updatedWorker.ID),
			FirstName:       updatedWorker.FirstName,
			MiddleName:      updatedWorker.MiddleName,
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
		MiddleName:      request.GetMiddleName(),
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
			Id:              uint64(medworker.ID),
			FirstName:       medworker.FirstName,
			MiddleName:      medworker.MiddleName,
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
