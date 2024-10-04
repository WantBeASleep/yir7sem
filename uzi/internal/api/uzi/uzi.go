package uzi

// пока 500-тим оставляем на рефакторинг

import (
	"context"
	"fmt"
	pb "yir/uzi/api"
	"yir/uzi/internal/api/mvpmappers"
	"yir/uzi/internal/api/usecases"

	protovalidate "github.com/bufbuild/protovalidate-go"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedUziAPIServer

	uziUseCase usecases.Uzi
	validator  *protovalidate.Validator
}

func NewServer(
	uziUseCase usecases.Uzi,
) (*Server, error) {
	// мб стоит придумать тут что то для парса
	validator, err := protovalidate.New(
		protovalidate.WithDisableLazy(true),
		protovalidate.WithMessages(
			&pb.Tirads{},
			&pb.Device{},
			&pb.Segment{},
			&pb.Formation{},
			&pb.Image{},
			&pb.UziInfo{},
			&pb.Uzi{},

			&pb.UziIdRequest{},
			&pb.UziIdRequest{},
			&pb.ImageIdRequest{},
			&pb.FormationIdRequest{},

			&pb.UpdateUziRequest{},
			&pb.InsertFormationWithSegmentsRequest{},
			&pb.UpdateFormationRequest{},
		),
	)
	if err != nil {
		return nil, fmt.Errorf("init validator: %w", err)
	}

	return &Server{
		uziUseCase: uziUseCase,
		validator:  validator,
	}, nil
}

func (s *Server) InsertUzi(ctx context.Context, req *pb.Uzi) (*empty.Empty, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("validation failed: %s", err.Error()))
	}

	// MVP mapper moment
	// Нужен кастомный маппер с возможностью задачи конфига маппинга между структурами
	// Здесь нужен именно конфиг для string --> uuid
	if err := s.uziUseCase.InsertUzi(ctx, mvpmappers.UziToDTOUzi(req)); err != nil {
		// пока 500-тим оставляем на рефакторинг
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Что то пошло не так: %s", err.Error()))
	}

	return &empty.Empty{}, nil
}

func (s *Server) GetUzi(ctx context.Context, req *pb.UziIdRequest) (*pb.Uzi, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("validation failed: %s", err.Error()))
	}

	uziID := uuid.MustParse(req.UziId)
	uzi, err := s.uziUseCase.GetUzi(ctx, uziID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Что то пошло не так: %s", err.Error()))
	}

}
