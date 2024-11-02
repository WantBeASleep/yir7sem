package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	pb "yir/s3upload/api"
)

type FileMeta struct {
	Path        string
	ContentType string
}

type File struct {
	FileMeta FileMeta
	FileBin  []byte
}

func fileMetaToPB(meta *FileMeta) *pb.FileMeta {
	return &pb.FileMeta{
		Path:        meta.Path,
		ContentType: meta.ContentType,
	}
}

func pBFileMetaToEntity(meta *pb.FileMeta) *FileMeta {
	return &FileMeta{
		Path:        meta.GetPath(),
		ContentType: meta.GetContentType(),
	}
}

type Config struct {
	UploaderBatchSize int
}

type Option func(*Config)

var (
	WithUploaderBatchSize = func(size int) Option {
		return Option(func(c *Config) {
			c.UploaderBatchSize = size
		})
	}
)

type S3Client struct {
	client pb.S3Client

	cfg Config
}

func NewS3Client(
	client pb.S3Client,
	opts ...Option,
) *S3Client {
	// default
	cfg := Config{
		UploaderBatchSize: 1024 * 1024,
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	return &S3Client{
		client: client,
		cfg:    cfg,
	}
}

func (c *S3Client) Upload(ctx context.Context, meta *FileMeta, fileBin io.Reader) error {
	stream, err := c.client.Upload(ctx)
	if err != nil {
		return fmt.Errorf("open stream to upload: %w", err)
	}

	buf := make([]byte, c.cfg.UploaderBatchSize)
	for {
		n, err := fileBin.Read(buf)
		if n != 0 {
			err := stream.Send(&pb.File{
				FileMeta: fileMetaToPB(meta),
				FileBin:  buf,
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

// используется стрим под капотом
func (c *S3Client) GetFullFileByStream(ctx context.Context, path string) (*File, error) {
	stream, err := c.client.Get(ctx, &pb.GetRequest{Path: path})
	if err != nil {
		return nil, fmt.Errorf("open stream to download: %w", err)
	}

	res := &File{}
	buf := bytes.Buffer{}
	i := 0
	for {
		file, err := stream.Recv()
		i++
		if file != nil {
			if i == 1 {
				res.FileMeta = *pBFileMetaToEntity(file.FileMeta)
			}
			buf.Write(file.FileBin)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("get receive from s3 service: %w", err)
		}
	}

	res.FileBin = buf.Bytes()
	return res, nil
}
