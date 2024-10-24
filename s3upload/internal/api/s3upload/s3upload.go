package s3upload

import (
	"fmt"
	"io"

	pb "yir/s3upload/api"
	"yir/s3upload/internal/api/usecases"
	"yir/s3upload/internal/entity"
	"yir/s3upload/internal/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Controller struct {
	pb.UnimplementedS3UploadServer

	uziUseCase usecases.Uzi
	logger     *zap.Logger
}

func NewController(
	uziUseCase usecases.Uzi,
	logger *zap.Logger,
) *Controller {
	return &Controller{
		uziUseCase: uziUseCase,
		logger:     logger,
	}
}

func (c *Controller) Upload(req pb.S3Upload_UploadServer) error {
	ctx := req.Context()
	reader := utils.NewUploadGRPCReader(req)

	path, err := reader.GetPath()
	if err != nil {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("get path failed: %v", err))
	}

	file := &entity.File{
		Path: path,
		Data: reader,
	}

	if err := c.uziUseCase.UploadFile(ctx, file); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("upload file: %v", err))
	}

	if err := req.SendAndClose(&emptypb.Empty{}); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("close gRPC stream: %v", err))
	}

	return nil
}

func (c *Controller) Get(req *pb.GetRequest, stream pb.S3Upload_GetServer) error {
	ctx := stream.Context()

	path := req.GetPath()

	s3Stream, err := c.uziUseCase.GetFile(ctx, path)
	if err != nil {
		return fmt.Errorf("get stream from S3: %w", err)
	}

	buff := make([]byte, 5*1024*1024) // 5MB

	for {
		n, err := s3Stream.Read(buff)
		if n != 0 {
			err := stream.Send(&pb.GetFile{
				FileContent: buff[:n],
			})
			if err != nil {
				return fmt.Errorf("stream image to client: %w", err)
			}
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("read from s3: %w", err)
		}
	}

	return nil
}
