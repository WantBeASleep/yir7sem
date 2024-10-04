// MVP мапперы, пока не найдем пиздатый маппер
package mvpmappers

import (
	"yir/pkg/mappers"
	pb "yir/uzi/api"
	"yir/uzi/internal/entity"
	"yir/uzi/internal/entity/dto"

	"github.com/google/uuid"
)

func PBImagesToImages(images []*pb.Image, uziID uuid.UUID) []entity.Image {
	entityImages := make([]entity.Image, 0, len(images))
	for _, v := range images {
		entityImages = append(entityImages, entity.Image{
			Id:    uuid.MustParse(v.Id),
			Url:   v.Url,
			UziID: uziID,
			Page:  int(v.Page),
		})
	}

	return entityImages
}

func PBFormationsToDTOFormations(formations []*pb.Formation) []dto.Formation {
	entityFormation := make([]dto.Formation, 0, len(formations))
	for _, v := range formations {
		entityFormation = append(entityFormation, dto.Formation{
			Id:     uuid.MustParse(v.Id),
			Ai:     v.Ai,
			Tirads: mappers.MustTransformObj[pb.Tirads, entity.Tirads](v.Tirads),
		})
	}

	return entityFormation
}

func PBSegmentsToDTOSegments(segments []*pb.Segment) []dto.Segment {
	entitySegments := make([]dto.Segment, 0, len(segments))
	for _, v := range segments {
		entitySegments = append(entitySegments, dto.Segment{
			Id:          uuid.MustParse(v.Id),
			ImageID:     uuid.MustParse(v.ImageId),
			FormationID: uuid.MustParse(v.FormationId),
			ContorURL:   v.ContorUrl,
			Tirads:      mappers.MustTransformObj[pb.Tirads, entity.Tirads](v.Tirads),
		})
	}

	return entitySegments
}

func PBUziInfoToUzi(uziInfo *pb.UziInfo) *entity.Uzi {
	return &entity.Uzi{
		Id:         uuid.MustParse(uziInfo.Id),
		Url:        uziInfo.Url,
		Projection: uziInfo.Projection,
		PatientID:  uuid.MustParse(uziInfo.PatientId),
		DeviceID:   int(uziInfo.DeviceId),
	}
}

func UziToDTOUzi(req *pb.Uzi) *dto.Uzi {

	images := PBImagesToImages(req.Images, uuid.MustParse(req.UziInfo.Id))
	formations := PBFormationsToDTOFormations(req.Formations)
	segments := PBSegmentsToDTOSegments(req.Segments)

	return &dto.Uzi{
		UziInfo: PBUziInfoToUzi(req.UziInfo),
		Images:     images,
		Formations: formations,
		Segments:   segments,
	}
}

