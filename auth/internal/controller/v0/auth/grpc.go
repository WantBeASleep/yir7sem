package auth

import (
	"context"
	"go.uber.org/zap"
	"yir/auth/internal/controller/usecases"
	pb "yir/auth/api/v0/auth"
)

// Тут вроде можно просто server, но я не определился
type AuthServer struct {
	pb.UnimplementedAuthServer

	authUseCase usecases.Auth
	logger *zap.Logger
}

func NewAuthServer(
	authUseCase usecases.Auth,
	logger *zap.Logger,
) *AuthServer {
	return &AuthServer{
		authUseCase: authUseCase,
		logger: logger,
	}
}

func (s *AuthServer) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, nil
}

func (s *AuthServer) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return nil, nil
}

func (s *AuthServer) TokenRefresh(ctx context.Context, request *pb.TokenRefreshRequest) (*pb.TokenRefreshResponse, error) {
	return nil, nil
}