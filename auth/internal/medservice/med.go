package services

import (
	"context"
	"fmt"
	"yir/auth/internal/entity"
	pb "yir/med/api"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type Service struct {
	medClient pb.MedWorkersClient
}

func NewService(medServiceAddress string) (*Service, error) {
	conn, err := grpc.Dial(medServiceAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to med service: %w", err)
	}
	medClient := pb.NewMedWorkersClient(conn)
	return &Service{medClient: medClient}, nil
}

func (s *Service) AddMed(ctx context.Context, createData *entity.RequestRegister) (uuid.UUID, error) {
	req := &pb.AddMedWorkerRequest{
		FirstName:       createData.FirstName,
		LastName:        createData.LastName,
		MiddleName:      createData.FathersName,
		MedOrganization: createData.MedOrg,
	}
	resp, err := s.medClient.AddMedWorker(ctx, req)
	if err != nil {
		return uuid.Nil, fmt.Errorf("call med service: %w", err)
	}

	return uuid.MustParse(resp.Worker.Id), nil
}
