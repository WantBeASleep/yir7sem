package uzi

import (
	pb "gateway/internal/generated/grpc/client/uzi"

	"google.golang.org/grpc"
)

type UziAdapter interface {
	pb.UziSrvClient
}

type adapter struct {
	pb.UziSrvClient
}

func New(
	conn *grpc.ClientConn,
) UziAdapter {
	client := pb.NewUziSrvClient(conn)

	return &adapter{
		UziSrvClient: client,
	}
}
