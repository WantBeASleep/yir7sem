package controller

import (
	"context"
	"errors"
	"yir/med/internal/entity"
	pb "yir/med/internal/pb/patient"
	"yir/med/internal/usecase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.UnimplementedMedPatientServer
	patientUseCase usecase.Patient
}

func NewServer(patientUseCase usecase.Patient) *Server {
	return &Server{
		patientUseCase: patientUseCase,
	}
}

func (s *Server) AddPatient(ctx context.Context, request *pb.CreatePatientRequest) (*emptypb.Empty, error) {
	PatientInfo := &entity.PatientInformation{
		Patient: &entity.Patient{
			ID:            uint64(request.Patient.Id),
			FirstName:     request.Patient.FirstName,
			LastName:      request.Patient.LastName,
			FatherName:    request.Patient.FatherName,
			MedicalPolicy: request.Patient.MedicalPolicy,
			Email:         request.Patient.Email,
			IsActive:      request.Patient.IsActive,
		},
		Card: &entity.PatientCard{
			ID:              uint64(request.PatientCard.Id),
			AppointmentTime: request.PatientCard.AppointmentTime,
			HasNodules:      request.PatientCard.HasNodules,
			Diagnosis:       request.PatientCard.Diagnosis,
			MedWorkerID:     request.PatientCard.MedWorkerId,
			PatientID:       request.PatientCard.PatientId,
		},
	}

	err := s.patientUseCase.AddPatient(ctx, PatientInfo)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return nil, nil
}

func (s *Server) UpdatePatient(ctx context.Context, request *pb.PatientUpdateRequest) (*emptypb.Empty, error) {
	PatientInfo := &entity.PatientInformation{
		Patient: &entity.Patient{
			ID:            uint64(request.Patient.Id),
			FirstName:     request.Patient.FirstName,
			LastName:      request.Patient.LastName,
			FatherName:    request.Patient.FatherName,
			MedicalPolicy: request.Patient.MedicalPolicy,
			Email:         request.Patient.Email,
			IsActive:      request.Patient.IsActive,
		},
		Card: &entity.PatientCard{
			ID:              uint64(request.PatientCard.Id),
			AppointmentTime: request.PatientCard.AppointmentTime,
			HasNodules:      request.PatientCard.HasNodules,
			Diagnosis:       request.PatientCard.Diagnosis,
			MedWorkerID:     request.PatientCard.MedWorkerId,
			PatientID:       request.PatientCard.PatientId,
		},
	}

	err := s.patientUseCase.UpdatePatient(ctx, PatientInfo)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrNotFound):
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return nil, nil
}

func (s *Server) GetPatientList(ctx context.Context, empty *emptypb.Empty) (*pb.PatientListResponse, error) {
	PatientList, err := s.patientUseCase.GetPatientList(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	patients := make([]*pb.Patient, len(PatientList))
	for i := range PatientList {
		patients[i] = &pb.Patient{
			Id:            PatientList[i].ID,
			FirstName:     PatientList[i].FirstName,
			LastName:      PatientList[i].LastName,
			FatherName:    PatientList[i].FatherName,
			MedicalPolicy: PatientList[i].MedicalPolicy,
			Email:         PatientList[i].Email,
			IsActive:      PatientList[i].IsActive,
		}
	}
	resp := &pb.PatientListResponse{
		Patients: patients,
	}
	return resp, nil
}

func (s *Server) GetPatientInfoByID(ctx context.Context, request *pb.PatientInfoRequest) (*pb.PatientInfoResponse, error) {
	PatientInfo, err := s.patientUseCase.GetPatientInfoByID(ctx, uint64(request.Id))
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrNotFound):
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return &pb.PatientInfoResponse{
		Patient: &pb.Patient{
			Id:            PatientInfo.Patient.ID,
			FirstName:     PatientInfo.Patient.FirstName,
			LastName:      PatientInfo.Patient.LastName,
			FatherName:    PatientInfo.Patient.FatherName,
			MedicalPolicy: PatientInfo.Patient.MedicalPolicy,
			Email:         PatientInfo.Patient.Email,
			IsActive:      PatientInfo.Patient.IsActive,
		},
		PatientCard: &pb.PatientCard{
			Id:              PatientInfo.Card.ID,
			AppointmentTime: PatientInfo.Card.AppointmentTime,
			HasNodules:      PatientInfo.Card.HasNodules,
			Diagnosis:       PatientInfo.Card.Diagnosis,
			MedWorkerId:     PatientInfo.Card.MedWorkerID,
			PatientId:       PatientInfo.Card.PatientID,
		},
	}, nil
}
