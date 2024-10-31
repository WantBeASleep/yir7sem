package s3upload

import (
	"bytes"
	"context"
	"fmt"
	"io"

	pb "yir/s3upload/api"
	"yir/s3upload/internal/api/mvpmappers"
	"yir/s3upload/internal/api/usecases"
	"yir/s3upload/internal/entity"
	"yir/s3upload/internal/utils"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Controller struct {
	pb.UnimplementedS3Server

	uziUseCase usecases.Uzi
}

func NewController(
	uziUseCase usecases.Uzi,
) *Controller {
	return &Controller{
		uziUseCase: uziUseCase,
	}
}

func (c *Controller) Upload(req pb.S3_UploadServer) error {
	ctx := req.Context()
	reader := utils.NewUploadGRPCReader(req)

	meta, err := reader.GetMeta()
	if err != nil || meta.GetPath() == "" {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("get meta failed: %v", err))
	}

	file := &entity.File{
		Meta: mvpmappers.PBFileMetaToEntity(meta),
		Data: reader,
	}

	if err := c.uziUseCase.UploadFile(ctx, file); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("upload file: %v", err))
	}

	if err := req.SendAndClose(&empty.Empty{}); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("close gRPC stream: %v", err))
	}

	return nil
}

func (c *Controller) UploadFull(ctx context.Context, req *pb.File) (*empty.Empty, error) {
	if req.GetFileMeta().GetPath() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "path is required")
	}

	buf := bytes.NewBuffer(req.FileBin)
	file := &entity.File{
		Meta: mvpmappers.PBFileMetaToEntity(req.FileMeta),
		Data: buf,
	}

	if err := c.uziUseCase.UploadFile(ctx, file); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("upload file: %v", err))
	}

	return &empty.Empty{}, nil
}

func (c *Controller) Get(req *pb.GetRequest, stream pb.S3_GetServer) error {
	ctx := stream.Context()

	path := req.GetPath()

	meta, s3Stream, err := c.uziUseCase.GetFile(ctx, path)
	if err != nil {
		return fmt.Errorf("get stream from S3: %w", err)
	}

	buff := make([]byte, 1024*1024) // 1MB

	for {
		n, err := s3Stream.Read(buff)
		if n != 0 {
			err := stream.Send(&pb.File{
				FileMeta: mvpmappers.EntityFileMetaToPB(meta),
				FileBin:  buff[:n],
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

func (c *Controller) GetFull(ctx context.Context, req *pb.GetRequest) (*pb.File, error) {
	path := req.GetPath()
	meta, s3Stream, err := c.uziUseCase.GetFile(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("get stream from S3: %w", err)
	}

	fileBin, err := io.ReadAll(s3Stream)
	if err != nil {
		return nil, fmt.Errorf("read full stream from S3: %w", err)
	}

	return &pb.File{
		FileMeta: mvpmappers.EntityFileMetaToPB(meta),
		FileBin:  fileBin,
	}, nil
}
