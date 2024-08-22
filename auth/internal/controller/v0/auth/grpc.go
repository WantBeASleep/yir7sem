package auth

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "yir/auth/api/v0/auth"
	"yir/auth/internal/controller/usecases"
	"yir/auth/internal/enity"

	"go.uber.org/zap"
)

// Тут вроде можно просто server, но я не определился
type AuthServer struct {
	pb.UnimplementedAuthServer

	authUseCase usecases.Auth
	logger      *zap.Logger
}

func NewAuthServer(
	authUseCase usecases.Auth,
	logger *zap.Logger,
) *AuthServer {
	return &AuthServer{
		authUseCase: authUseCase,
		logger:      logger,
	}
}

func (s *AuthServer) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	domainRequest := enity.RequestLogin{
		Email:    request.Email,
		Password: request.Password,
	}

	tokenPair, err := s.authUseCase.Login(ctx, &domainRequest)
	if err != nil {
		switch {
		case errors.Is(err, enity.ErrNotFound):
			return nil, status.Error(codes.NotFound, err.Error())
		case errors.Is(err, enity.ErrPassHashNotEqual):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	return &pb.LoginResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}, nil
}

func (s *AuthServer) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return nil, nil
}

func (s *AuthServer) TokenRefresh(ctx context.Context, request *pb.TokenRefreshRequest) (*pb.TokenRefreshResponse, error) {
	return nil, nil
}
