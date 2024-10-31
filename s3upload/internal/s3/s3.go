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

func s3MetaToEntityMeta(meta minio.ObjectInfo) *entity.FileMeta {
	res := entity.FileMeta{}

	res.Path = meta.Key
	res.ContentType = meta.ContentType

	return &res
}

func entityMetaToMinioPutOpts(meta *entity.FileMeta) minio.PutObjectOptions {
	res := minio.PutObjectOptions{}

	if meta.ContentType != "" {
		res.ContentType = meta.ContentType
	}

	return res
}

func (r *Repo) Upload(ctx context.Context, file *entity.File) error {
	// поправить с -1, на нужный размер, пока что так
	minioOpts := entityMetaToMinioPutOpts(file.Meta)

	_, err := r.client.PutObject(ctx, r.bucket, file.Meta.Path, file.Data, -1, minioOpts)
	if err != nil {
		return fmt.Errorf("upload to S3: %w", err)
	}

	return err
}

// stream файла, поэтому io.ReadCloser
func (r *Repo) Get(ctx context.Context, path string) (*entity.FileMeta, io.Reader, error) {

	metaInfo, err := r.client.StatObject(ctx, r.bucket, path, minio.StatObjectOptions{})
	if err != nil {
		return nil, nil, fmt.Errorf("get meta info from S3: %w", err)
	}

	obj, err := r.client.GetObject(ctx, r.bucket, path, minio.GetObjectOptions{})
	if err != nil {
		return nil, nil, fmt.Errorf("get obj from S3: %w", err)
	}

	return s3MetaToEntityMeta(metaInfo), obj, nil
}
