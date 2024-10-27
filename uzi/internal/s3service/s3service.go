package s3service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	pb "yir/s3upload/api"
)

type S3 struct {
	client pb.S3UploadClient
}

func NewS3Service(
	client pb.S3UploadClient,
) *S3 {
	return &S3{
		client: client,
	}
}

// полностью загрузить в память
func (s *S3) FullGetByPath(ctx context.Context, path string) ([]byte, error) {
	stream, err := s.client.Get(ctx, &pb.GetRequest{Path: path})
	if err != nil {
		return nil, fmt.Errorf("open stream to download: %w", err)
	}

	buf := bytes.Buffer{}
	for {
		file, err := stream.Recv()
		if file != nil {
			buf.Write(file.FileContent)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("get receive from s3 service: %w", err)
		}
	}

	return buf.Bytes(), nil
}

func (s *S3) Upload(ctx context.Context, path string, data []byte) error {
	stream, err := s.client.Upload(ctx)
	if err != nil {
		return fmt.Errorf("open stream to upload: %w", err)
	}

	reader := bytes.NewReader(data)
	buf := [1024 * 1024]byte{}
	for {
		n, err := reader.Read(buf[:])
		if n != 0 {
			err := stream.Send(&pb.UploadFile{
				Path: path,
				File: buf[:],
			})
			if err != nil {
				return fmt.Errorf("send data to upload stream: %w", err)
			}
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("read from buffered img data: %w", err)
		}
	}

	_, err = stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("close upload stream: %w", err)
	}

	return nil
}
