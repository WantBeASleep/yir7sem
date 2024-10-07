package s3upload

import (
	"bytes"
	"io"
	"fmt"

	pb "yir/s3upload/api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Controller struct{
	pb.UnimplementedS3UploadServer
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

	

	return nil
}

