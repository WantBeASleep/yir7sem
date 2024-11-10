package s3

import (
	"bytes"
	"context"
	"fmt"
	"yir/s3upload/pkg/client"
	"yir/uzi/internal/entity/imagesplitter"
)

type Client struct {
	client *client.S3Client
}

func NewClient(
	client *client.S3Client,
) *Client {
	return &Client{
		client: client,
	}
}

func (c *Client) Upload(ctx context.Context, path string, file *imagesplitter.File) error {
	reader := bytes.NewBuffer(file.FileBin)

	if err := c.client.Upload(ctx, &client.FileMeta{
		Path:        path,
		ContentType: file.FileMeta.ContentType,
	}, reader); err != nil {
		return fmt.Errorf("load file to S3: %w", err)
	}

	return nil
}

func (c *Client) GetFile(ctx context.Context, path string) (*imagesplitter.File, error) {
	file, err := c.client.GetFullFileByStream(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("get file from s3: %w", err)
	}

	return &imagesplitter.File{
		FileMeta: imagesplitter.FileMeta{
			ContentType: file.FileMeta.ContentType,
		},
		FileBin: file.FileBin,
	}, nil
}
