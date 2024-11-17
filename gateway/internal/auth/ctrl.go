package auth

import (
	pb "yir/gateway/rpc/auth"
	"context"
)

type Ctrl struct {
	pb.UnimplementedAuthServer

	client pb.AuthClient
}

func NewCtrl(client pb.AuthClient) *Ctrl {
	return &Ctrl{client: client}
}

func (c *Ctrl) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return c.client.Login(ctx, req)
}
func (c *Ctrl) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return c.client.Register(ctx, req)
}
func (c *Ctrl) TokenRefresh(ctx context.Context, req *pb.TokenRefreshRequest) (*pb.TokenRefreshResponse, error) {
	return c.client.TokenRefresh(ctx, req)
}