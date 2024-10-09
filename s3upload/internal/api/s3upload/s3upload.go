package s3upload

import (
	"bytes"
	"fmt"
	"io"

	pb "yir/s3upload/api"
	"yir/s3upload/internal/api/usecases"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Controller struct {
	pb.UnimplementedS3UploadServer

	uziUseCase usecases.Uzi
}

func NewController(
	uziUseCase usecases.Uzi,
) *Controller {
	return &Controller{
		uziUseCase: uziUseCase,
	}
}

func (c *Controller) UploadAndSplitUziFile(req pb.S3Upload_UploadAndSplitUziFileServer) error {
	ctx := req.Context()
	buff := bytes.Buffer{}

	// потоково разбить на изображения .tiff не получится, так что будем просто загружать в оперативку
	for {
		req, err := req.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return status.Errorf(codes.Internal, fmt.Sprintf("stream read failed: %v", err))
		}

		_, err = buff.Write(req.File)
		if err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("add stream part to buff failed: %v", err))
		}
	}

	uziID, splittedIDs, err := c.uziUseCase.UploadAndSplitUziFile(ctx, buff.Bytes())
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("upload and splitting images: %v", err))
	}

	uziIDStr := uziID.String()
	splittedIDsStr := make([]string, 0, len(splittedIDs))
	for _, v := range splittedIDs {
		splittedIDsStr = append(splittedIDsStr, v.String())
	}

	err = req.SendAndClose(&pb.UploadUziFileResponse{
		UziId:     uziIDStr,
		ImagesIds: splittedIDsStr,
	})

	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("close gRPC stream: %v", err))
	}

	return nil
}
