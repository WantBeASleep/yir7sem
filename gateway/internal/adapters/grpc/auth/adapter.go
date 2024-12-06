package auth

import (
	pb "gateway/internal/generated/grpc/client/auth"

	"google.golang.org/grpc"
)

type AuthAdapter interface {
	pb.AuthSrvClient
}

type adapter struct {
	pb.AuthSrvClient
}

func New(
	conn *grpc.ClientConn,
) AuthAdapter {
	client := pb.NewAuthSrvClient(conn)

	return &adapter{
		AuthSrvClient: client,
	}
}
