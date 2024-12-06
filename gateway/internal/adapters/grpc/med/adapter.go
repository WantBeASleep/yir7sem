package med

import (
	pb "gateway/internal/generated/grpc/client/med"

	"google.golang.org/grpc"
)

type MedAdapter interface {
	pb.MedSrvClient
}

type adapter struct {
	pb.MedSrvClient
}

func New(
	conn *grpc.ClientConn,
) MedAdapter {
	client := pb.NewMedSrvClient(conn)

	return &adapter{
		MedSrvClient: client,
	}
}
