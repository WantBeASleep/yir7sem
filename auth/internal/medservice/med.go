package services

import (
	"context"
	"fmt"
	pb "yir/all/api"
	"yir/auth/internal/entity"

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

func (s *Service) AddMed(ctx context.Context, createData *entity.RequestRegister) (int, error) {
	req := &pb.AddMedWorkerRequest{
		FirstName:       createData.FirstName,
		LastName:        createData.LastName,
		MiddleName:      createData.FathersName,
		MedOrganization: createData.MedOrg,
	}
	resp, err := s.medClient.AddMedWorker(ctx, req)
	if err != nil {
		return 0, fmt.Errorf("call med service: %w", err)
	}

	return int(resp.Worker.Id), nil
}
