// TODO: возможно все это полная хуета, но я настолько устал, что я рот ебал что то сейчас делать
package repository

import (
	minio "github.com/minio/minio-go/v7"
)

type DAO interface {
	NewFileRepo() FileRepo
}

type dao struct {
	s3       *minio.Client
	s3bucket string
}

func NewRepository(s3 *minio.Client, s3bucket string) DAO {
	return &dao{
		s3:       s3,
		s3bucket: s3bucket,
	}
}

// SS3
func (d *dao) NewFileRepo() FileRepo {
	return &fileRepo{
		s3:     d.s3,
		bucket: d.s3bucket,
	}
}
