package mvpmappers

import (
	"yir/s3upload/api"
	"yir/s3upload/internal/entity"
)

func PBFileMetaToEntity(pb *api.FileMeta) *entity.FileMeta {
	return &entity.FileMeta{
		Path:        pb.GetPath(),
		ContentType: pb.GetContentType(),
	}
}

func EntityFileMetaToPB(ent *entity.FileMeta) *api.FileMeta {
	return &api.FileMeta{
		Path:        ent.Path,
		ContentType: ent.ContentType,
	}
}
