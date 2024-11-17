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
			&pb.UziRequest{},
			&pb.Id{},
			&pb.FormationWithNestedSegmentsRequest{},
			&pb.SegmentNestedRequest{},

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

func (s *Server) CreateUzi(ctx context.Context, req *pb.UziRequest) (*pb.Id, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("validation failed: %v", err))
	}

	id, err := s.uziUseCase.CreateUzi(ctx, mvpmappers.PBUziReqToUzi(req))
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Что то пошло не так: %s", err.Error()))
	}

	return &pb.Id{
		Id: id.String(),
	}, nil
}

func (s *Server) GetUzi(ctx context.Context, req *pb.Id) (*pb.UziReponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("validation failed: %s", err.Error()))
	}

	uziID := uuid.MustParse(req.Id)
	uzi, err := s.uziUseCase.GetUzi(ctx, uziID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Что то пошло не так: %s", err.Error()))
	}

	return mvpmappers.UziToPBUziResp(uzi), nil
}

func (s *Server) GetReport(ctx context.Context, req *pb.Id) (*pb.Report, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("validation failed: %s", err.Error()))
	}

	reportID := uuid.MustParse(req.Id)
	report, err := s.uziUseCase.GetReport(ctx, reportID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Что то пошло не так: %s", err.Error()))
	}

	return mvpmappers.DTOReportToPBReport(report), nil
}

func (s *Server) UpdateUzi(ctx context.Context, req *pb.UpdateUziRequest) (*pb.UziReponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("validation failed: %s", err.Error()))
	}

	uziID := uuid.MustParse(req.Id)
	uzi := mvpmappers.PBUpdateUziReqToUzi(req.Uzi)
	updatedUzi, err := s.uziUseCase.UpdateUzi(ctx, uziID, uzi)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Что то пошло не так: %s", err.Error()))
	}

	return mvpmappers.UziToPBUziResp(updatedUzi), nil
}

func (s *Server) GetImageWithFormationsSegments(ctx context.Context, req *pb.Id) (*pb.ImageWithFormationsSegments, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("validation failed: %s", err.Error()))
	}

	imageID := uuid.MustParse(req.Id)
	imageWithSegments, err := s.uziUseCase.GetImageWithFormationsSegments(ctx, imageID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Что то пошло не так: %s", err.Error()))
	}

	return mvpmappers.DTOImageWithFormationsSegmentsToPBImageWithFormationsSegments(imageWithSegments), nil
}

func (s *Server) CreateFormationWithSegments(ctx context.Context, req *pb.CreateFormationWithSegmentsRequest) (*pb.CreateFormationWithSegmentsResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("validation failed: %s", err.Error()))
	}

	formationWithSegment := mvpmappers.PBCreateFormationWithSegmentsReqToDTOFormationWithSegments(req)
	formationID, segmentsIDS, err := s.uziUseCase.CreateFormationWithSegments(ctx, formationWithSegment)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Что то пошло не так: %s", err.Error()))
	}

	return &pb.CreateFormationWithSegmentsResponse{
		FormationId: formationID.String(),
		SegmentsIds: segmentsIDS.Strings(),
	}, nil
}

func (s *Server) GetFormationWithSegments(ctx context.Context, req *pb.Id) (*pb.FormationWithSegments, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("validation failed: %s", err.Error()))
	}

	formationID := uuid.MustParse(req.Id)
	formationsWithSegments, err := s.uziUseCase.GetFormationWithSegments(ctx, formationID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Что то пошло не так: %s", err.Error()))
	}

	return mvpmappers.DTOFormationWithSegmentsToPBFormationWithSegments(formationsWithSegments), nil
}

func (s *Server) UpdateFormation(ctx context.Context, req *pb.UpdateFormationRequest) (*pb.FormationResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("validation failed: %s", err.Error()))
	}

	formationID := uuid.MustParse(req.Id)
	formation := mvpmappers.PBFormationReqToDTOFormation(req.Formation)

	updatedFormation, err := s.uziUseCase.UpdateFormation(ctx, formationID, formation)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Что то пошло не так: %s", err.Error()))
	}

	return mvpmappers.DTOFormationToPBFormationResp(updatedFormation), nil
}

func (s *Server) UpdateSegment(ctx context.Context, req *pb.UpdateSegmentRequest) (*pb.SegmentResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("validation failed: %s", err.Error()))
	}

	segmentID := uuid.MustParse(req.Id)
	segment := mvpmappers.PBSegmentReqToDTOSegment(&pb.SegmentRequest{Tirads: req.Tirads})

	updatedSegment, err := s.uziUseCase.UpdateSegment(ctx, segmentID, segment)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Что то пошло не так: %s", err.Error()))
	}

	return mvpmappers.DTOSegmentToPBSegmentResp(*updatedSegment), nil
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
