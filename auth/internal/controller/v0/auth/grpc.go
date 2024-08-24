// порядок должены быть
// validation
// mapping
// use-case call
// errors handle
// mapping

package auth

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "yir/auth/api/v0/auth"
	"yir/auth/internal/controller/usecases"
	"yir/auth/internal/enity"
	"yir/auth/internal/controller/validation"
)

// Тут вроде можно просто server, но я не определился
type AuthServer struct {
	pb.UnimplementedAuthServer

	authUseCase usecases.Auth
}

func NewAuthServer(
	authUseCase usecases.Auth,
) *AuthServer {
	return &AuthServer{
		authUseCase: authUseCase,
	}
}

func (s *AuthServer) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	if err :=validation.ValidateLoginRequest(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

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
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &pb.LoginResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}, nil
}