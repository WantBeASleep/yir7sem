package auth

import (
	"context"
	"errors"
	"fmt"
	pb "yir/auth/api/auth"
	"yir/auth/internal/api/usecases"
	"yir/auth/internal/entity"

	protovalidate "github.com/bufbuild/protovalidate-go"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedAuthServer

	authUseCase usecases.Auth
	validator   *protovalidate.Validator
}

func NewServer(
	authUseCase usecases.Auth,
) (*Server, error) {
	validator, err := protovalidate.New(
		protovalidate.WithDisableLazy(true),
		protovalidate.WithMessages(
			&pb.LoginRequest{},
			&pb.RegisterRequest{},
			&pb.TokenRefreshRequest{},
		),
	)
	if err != nil {
		return nil, fmt.Errorf("init validator: %w", err)
	}

	return &Server{
		authUseCase: authUseCase,
		validator:   validator,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("validation failed: %v", err))
	}

	tokensPair, err := s.authUseCase.Login(ctx, req.Mail, req.Password)
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

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("validation failed: %v", err))
	}

	domainRegister := entity.RequestRegister{
		Mail:        req.Mail,
		LastName:    req.LastName,
		FirstName:   req.FirstName,
		FathersName: req.FathersName,
		MedOrg:      req.MedOrganization,
		Password:    req.Password,
	}

	userID, err := s.authUseCase.Register(ctx, &domainRegister)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.RegisterResponse{
		Id: userID.String(),
	}, nil
}

func (s *Server) TokenRefresh(ctx context.Context, req *pb.TokenRefreshRequest) (*pb.TokenRefreshResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("validation failed: %v", err))
	}

	tokensPair, err := s.authUseCase.TokenRefresh(ctx, req.RefreshToken)
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
