package repository

import (
	"context"
	"io"
	"yir/gateway/internal/entity"
	"yir/s3upload/pkg/client"
)

type S3Repo struct {
	S3 *client.S3Client
}

func (c *S3Repo) Upload(ctx context.Context, meta *entity.FileMeta, fileBin io.Reader) error {
	metaData := &client.FileMeta{
		Path:        meta.Path,
		ContentType: meta.ContentType,
	}
	if err := c.S3.Upload(ctx, metaData, fileBin); err != nil {
		return err
	}
	return nil
}

func (c *S3Repo) Get(ctx context.Context, path string) (*entity.File, error) {
	data, err := c.S3.GetFullFileByStream(ctx, path)
	if err != nil {
		return nil, err
	}
	resp := &entity.File{
		FileMeta: &entity.FileMeta{
			Path:        data.FileMeta.Path,
			ContentType: data.FileMeta.ContentType,
		},
		FileBin: data.FileBin,
	}
	return resp, nil
}
