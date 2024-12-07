package repository

import (
	"context"

	"github.com/WantBeASleep/goooool/daolib"

	"github.com/jmoiron/sqlx"
	minio "github.com/minio/minio-go/v7"
)

type DAO interface {
	daolib.DAO
	NewFileRepo() FileRepo
	NewDeviceQuery(ctx context.Context) DeviceQuery
	NewUziQuery(ctx context.Context) UziQuery
	NewImageQuery(ctx context.Context) ImageQuery
	NewSegmentQuery(ctx context.Context) SegmentQuery
	NewNodeQuery(ctx context.Context) NodeQuery
	NewEchographicQuery(ctx context.Context) EchographicQuery
}

type dao struct {
	daolib.DAO

	s3       *minio.Client
	s3bucket string
}

func NewRepository(psql *sqlx.DB, s3 *minio.Client, s3bucket string) DAO {
	return &dao{
		DAO:      daolib.NewDao(psql),
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

// POSTNIGRES
func (d *dao) NewDeviceQuery(ctx context.Context) DeviceQuery {
	deviceQuery := &deviceQuery{}
	d.NewRepo(ctx, deviceQuery)

	return deviceQuery
}

func (d *dao) NewUziQuery(ctx context.Context) UziQuery {
	uziQuery := &uziQuery{}
	d.NewRepo(ctx, uziQuery)

	return uziQuery
}

func (d *dao) NewImageQuery(ctx context.Context) ImageQuery {
	imageQuery := &imageQuery{}
	d.NewRepo(ctx, imageQuery)

	return imageQuery
}

func (d *dao) NewSegmentQuery(ctx context.Context) SegmentQuery {
	segment := &segmentQuery{}
	d.NewRepo(ctx, segment)

	return segment
}

func (d *dao) NewNodeQuery(ctx context.Context) NodeQuery {
	node := &nodeQuery{}
	d.NewRepo(ctx, node)

	return node
}

func (d *dao) NewEchographicQuery(ctx context.Context) EchographicQuery {
	echographic := &echographicQuery{}
	d.NewRepo(ctx, echographic)

	return echographic
}
