// MVP мапперы, пока не найдем пиздатый маппер
package mvpmappers

import (
	"yir/pkg/mappers"
	pb "yir/uzi/api"
	"yir/uzi/internal/entity"
	"yir/uzi/internal/entity/dto"

	"github.com/google/uuid"
)

func UziToDTOUzi(req *pb.Uzi) *dto.Uzi {

	images := make([]entity.Image, 0, len(req.Images))
	for _, v := range req.Images {
		images = append(images, entity.Image{
			Id:    uuid.MustParse(v.Id),
			Url:   v.Url,
			UziID: uuid.MustParse(req.UziInfo.Id),
			Page:  int(v.Page),
		})
	}

	formations := make([]dto.Formation, 0, len(req.Formations))
	for _, v := range req.Formations {
		formations = append(formations, dto.Formation{
			Id:     uuid.MustParse(v.Id),
			Ai:     v.Ai,
			Tirads: mappers.MustTransformObj[pb.Tirads, entity.Tirads](v.Tirads),
		})
	}

	segments := make([]dto.Segment, 0, len(req.Segments))
	for _, v := range req.Segments {
		segments = append(segments, dto.Segment{
			Id:          uuid.MustParse(v.Id),
			ImageID:     uuid.MustParse(v.ImageId),
			FormationID: uuid.MustParse(v.FormationId),
			ContorURL:   v.ContorUrl,
			Tirads:      mappers.MustTransformObj[pb.Tirads, entity.Tirads](v.Tirads),
		})
	}

	return &dto.Uzi{
		UziInfo: &entity.Uzi{
			Id:         uuid.MustParse(req.UziInfo.Id),
			Url:        req.UziInfo.Url,
			Projection: req.UziInfo.Projection,
			PatientID:  uuid.MustParse(req.UziInfo.PatientId),
			DeviceID:   int(req.UziInfo.DeviceId),
		},
		Images:     images,
		Formations: formations,
		Segments:   segments,
	}
}

func DTOUziToUzi(uzi *dto.Uzi) *pb.Uzi {

}
