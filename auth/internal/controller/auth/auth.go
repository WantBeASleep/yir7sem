package auth

import (
	"context"
	"errors"
	pb "service/auth/api/auth"
	"service/auth/internal/controller/usecases"
	"service/auth/internal/controller/validation"
	"service/auth/internal/entity"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedAuthServer

	authUseCase usecases.Auth
}

func NewServer(
	authUseCase usecases.Auth,
) *Server {
	return &Server{
		authUseCase: authUseCase,
	}
}

func (s *Server) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	if err := validation.ValidateLoginRequest(request); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation login request: %v", err.Error())
	}

	domainRequest := entity.RequestLogin{
		Email:    request.Email,
		Password: request.Password,
	}

	tokensPair, err := s.authUseCase.Login(ctx, &domainRequest)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrNotFound):
			return nil, status.Error(codes.NotFound, err.Error())
		case errors.Is(err, entity.ErrPassHashNotEqual):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &pb.LoginResponse{
		AccessToken:  tokensPair.AccessToken,
		RefreshToken: tokensPair.RefreshToken,
	}, nil
}

func (s *Server) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if err := validation.ValidateRegisterRequest(request); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation register request: %v", err.Error())
	}

	domainRegister := entity.RequestRegister{
		Email:       request.Email,
		LastName:    request.LastName,
		FirstName:   request.FirstName,
		FathersName: request.FathersName,
		MedOrg:      request.MedOrganization,
		Password:    request.Password,
	}

	userID, err := s.authUseCase.Register(ctx, &domainRegister)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.RegisterResponse{
		UserId: userID,
	}, nil
}

func (s *Server) TokenRefresh(ctx context.Context, request *pb.TokenRefreshRequest) (*pb.TokenRefreshResponse, error) {
	if err := validation.ValidateTokenRefreshRequest(request); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation token refresh request: %v", err.Error())
	}

	tokensPair, err := s.authUseCase.TokenRefresh(ctx, request.RefreshToken)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrExpiredSession):
			return nil, status.Error(codes.Unauthenticated, err.Error())
		case errors.Is(err, entity.ErrInvalidToken):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, entity.ErrNotFound):
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &pb.TokenRefreshResponse{
		AccessToken:  tokensPair.AccessToken,
		RefreshToken: tokensPair.RefreshToken,
	}, nil
}
