package s3

import (
	"bytes"
	"context"
	"fmt"
	"path/filepath"
	"yir/s3upload/internal/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Repo struct {
	client *minio.Client

	// не уверен что так нужно делать, но по идее разные бакеты будут просто разные репы
	bucket string
}

// Dependency Injection тут ебать как по пизде идет, в тех долг
// бакет хуякет, надо подумать как тут лучше будет, пока что MVP
func NewRepo(cfg *config.S3, bucket string) (*Repo, error) {
	client, err := minio.New(cfg.Endpoint, &minio.Options{
		// MVP
		Secure: false,
		Creds:  credentials.NewStaticV4(cfg.AccessToken, cfg.SecretToken, ""),
	})
	if err != nil {
		return nil, fmt.Errorf("init S3 client: %w", err)
	}

	return &Repo{
		client: client,
		bucket: bucket,
	}, nil
}

func (r *Repo) Upload(ctx context.Context, path string, filename string, data []byte) error {
	// поправить с -1, на нужный размер, пока что так
	// надо в entity сделать структурку для метаданных, но опять же, пока что плевать
	_, err := r.client.PutObject(ctx, r.bucket, filepath.Join(path, filename), bytes.NewBuffer(data), -1, minio.PutObjectOptions{})
	if err != nil {
		return fmt.Errorf("upload to S3: %w", err)
	}

	return err
}
