package s3

import (
	"context"
	"fmt"
	"io"
	"yir/s3upload/internal/config"
	"yir/s3upload/internal/entity"

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

func entityLoadOptsToMinioPutOpts(opts []entity.LoadOption) minio.PutObjectOptions {
	loadOpts := entity.LoadOpts{}

	for _, opt := range opts {
		opt(&loadOpts)
	}

	minioOpts := minio.PutObjectOptions{
		ContentType: loadOpts.ContentType,
	}

	return minioOpts
}

func entityGetOptsToMinioGetOpts(opts []entity.GetOption) minio.GetObjectOptions {
	getOpts := entity.GetOpts{}

	for _, opt := range opts {
		opt(&getOpts)
	}

	minioOpts := minio.GetObjectOptions{}

	return minioOpts
}

func (r *Repo) Upload(ctx context.Context, path string, data io.Reader, opts ...entity.LoadOption) error {
	// поправить с -1, на нужный размер, пока что так

	minioOpts := entityLoadOptsToMinioPutOpts(opts)

	_, err := r.client.PutObject(ctx, r.bucket, path, data, -1, minioOpts)
	if err != nil {
		return fmt.Errorf("upload to S3: %w", err)
	}

	return err
}

// stream файла, поэтому io.ReadCloser
func (r *Repo) Get(ctx context.Context, path string, opts ...entity.GetOption) (io.Reader, error) {

	minioOpts := entityGetOptsToMinioGetOpts(opts)

	obj, err := r.client.GetObject(ctx, r.bucket, path, minioOpts)
	if err != nil {
		return nil, fmt.Errorf("get obj from S3: %w", err)
	}

	return obj, nil
}
