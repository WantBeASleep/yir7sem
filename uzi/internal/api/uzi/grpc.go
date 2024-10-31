package uzi

// пока 500-тим оставляем на рефакторинг

import (
	"context"
	"fmt"
	pb "yir/uzi/api/grpcapi"
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
			&pb.SegmentRequest{},
			&pb.FormationRequest{},
			&pb.ImageRequest{},
			&pb.UziRequest{},
			&pb.Id{},

			&pb.CreateUziRequest{},
			&pb.UpdateUziRequest{},
			&pb.CreateFormationWithSegmentsRequest{},
			&pb.UpdateFormationRequest{},
			&pb.UpdateSegmentRequest{},
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

func (s *Server) CreateUzi(ctx context.Context, req *pb.CreateUziRequest) (*pb.Id, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("validation failed: %v", err))
	}

	if err := s.uziUseCase.CreateUziInfo(ctx, mvpmappers.PBCreateUziInfoReqToUzi(req)); err != nil {
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

	return mvpmappers.DTOUziToPBUzi(uzi), nil
}

func (s *Server) GetUziInfo(ctx context.Context, req *pb.UziIdRequest) (*pb.UziInfo, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("validation failed: %s", err.Error()))
	}

	uziID := uuid.MustParse(req.UziId)
	uziInfo, err := s.uziUseCase.GetUziInfo(ctx, uziID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Что то пошло не так: %s", err.Error()))
	}

	return mvpmappers.UziToPBUziInfo(uziInfo), nil
}

func (s *Server) UpdateUziInfo(ctx context.Context, req *pb.UpdateUziRequest) (*empty.Empty, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("validation failed: %s", err.Error()))
	}

	uziID := uuid.MustParse(req.UziId)
	uziInfo := mvpmappers.PBUziInfoToUzi(req.UziInfo)
	if err := s.uziUseCase.UpdateUziInfo(ctx, uziID, uziInfo); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Что то пошло не так: %s", err.Error()))
	}

	return &empty.Empty{}, nil
}

func (s *Server) GetImageWithSegments(ctx context.Context, req *pb.ImageIdRequest) (*pb.ImageWithSegments, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("validation failed: %s", err.Error()))
	}

	imageID := uuid.MustParse(req.ImageId)
	imageWithSegments, err := s.uziUseCase.GetImageWithSegmentsFormations(ctx, imageID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Что то пошло не так: %s", err.Error()))
	}

	return mvpmappers.DTOImageWithSegmentsToPBImageWithSegments(imageWithSegments), nil
}

func (s *Server) InsertFormationWithSegments(ctx context.Context, req *pb.InsertFormationWithSegmentsRequest) (*empty.Empty, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("validation failed: %s", err.Error()))
	}

	// Не отследил при проектировании API
	// Добавляем узлы с сегментами, сегменты строго привязанные к картинкам
	// Картинки строго привязанны в uzi --> не нужен uziID здесь
	// uziID := uuid.MustParse(req.UziId)

	formationWithSegment := mvpmappers.PBFormationWithSegmentsToDTOFormationWithSegments(req.FormationWithSegments)
	if err := s.uziUseCase.InsertFormationWithSegments(ctx, formationWithSegment); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Что то пошло не так: %s", err.Error()))
	}

	return &empty.Empty{}, nil
}

func (s *Server) GetFormationWithSegments(ctx context.Context, req *pb.FormationIdRequest) (*pb.FormationWithSegments, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("validation failed: %s", err.Error()))
	}

	formationID := uuid.MustParse(req.FormationId)
	formationsWithSegments, err := s.uziUseCase.GetFormationWithSegments(ctx, formationID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Что то пошло не так: %s", err.Error()))
	}

	return mvpmappers.DTOFormationWithSegmentsToPBFormationWithSegments(formationsWithSegments), nil
}

func (s *Server) UpdateFormation(ctx context.Context, req *pb.UpdateFormationRequest) (*empty.Empty, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("validation failed: %s", err.Error()))
	}

	formationID := uuid.MustParse(req.FormationId)
	formation := mvpmappers.PBFormationsToDTOFormations([]*pb.Formation{req.Formation})[0]

	if err := s.uziUseCase.UpdateFormation(ctx, formationID, &formation); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Что то пошло не так: %s", err.Error()))
	}

	return &empty.Empty{}, nil
}

func (s *Server) GetDeviceList(ctx context.Context, _ *empty.Empty) (*pb.GetDeviceListResponse, error) {
	devices, err := s.uziUseCase.DeviceList(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Что то пошло не так: %s", err.Error()))
	}

	return &pb.GetDeviceListResponse{
		Devices: mvpmappers.DevicesToPBDevices(devices),
	}, nil
}
