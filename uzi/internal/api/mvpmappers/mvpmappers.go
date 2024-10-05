// MVP мапперы, пока не найдем пиздатый маппер
// или напишем сами, но я не могу тратить на это сейчас время
package mvpmappers

import (
	"yir/pkg/mappers"
	pb "yir/uzi/api"
	"yir/uzi/internal/entity"
	"yir/uzi/internal/entity/dto"

	"github.com/google/uuid"
)

func PBTiradsToTirads(tirads *pb.Tirads) *entity.Tirads {
	return &entity.Tirads{
		Tirads1: tirads.Tirads_1,
		Tirads2: tirads.Tirads_2,
		Tirads3: tirads.Tirads_3,
		Tirads4: tirads.Tirads_4,
		Tirads5: tirads.Tirads_5,
	}
}

func TiradsToPBTirads(tirads *entity.Tirads) *pb.Tirads {
	return &pb.Tirads{
		Tirads_1: tirads.Tirads1,
		Tirads_2: tirads.Tirads2,
		Tirads_3: tirads.Tirads3,
		Tirads_4: tirads.Tirads4,
		Tirads_5: tirads.Tirads5,
	}
}

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

func ImagesToPBImages(images []entity.Image) []*pb.Image {
	PBImages := make([]*pb.Image, 0, len(images))
	for _, v := range images {
		PBImages = append(PBImages, &pb.Image{
			Id:   v.Id.String(),
			Url:  v.Url,
			Page: int64(v.Page),
		})
	}

	return PBImages
}

func PBFormationsToDTOFormations(formations []*pb.Formation) []dto.Formation {
	entityFormation := make([]dto.Formation, 0, len(formations))
	for _, v := range formations {
		entityFormation = append(entityFormation, dto.Formation{
			Id:     uuid.MustParse(v.Id),
			Ai:     v.Ai,
			Tirads: PBTiradsToTirads(v.Tirads),
		})
	}

	return entityFormation
}

func DTOFormationsToPBFormations(formations []dto.Formation) []*pb.Formation {
	PBFormations := make([]*pb.Formation, 0, len(formations))
	for _, v := range formations {
		PBFormations = append(PBFormations, &pb.Formation{
			Id:     v.Id.String(),
			Ai:     v.Ai,
			Tirads: TiradsToPBTirads(v.Tirads),
		})
	}

	return PBFormations
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

func DTOSegmentsToPBSegments(segments []dto.Segment) []*pb.Segment {
	PBSegments := make([]*pb.Segment, 0, len(segments))
	for _, v := range segments {
		PBSegments = append(PBSegments, &pb.Segment{
			Id:          v.Id.String(),
			ImageId:     v.ImageID.String(),
			FormationId: v.FormationID.String(),
			ContorUrl:   v.ContorURL,
			Tirads:      TiradsToPBTirads(v.Tirads),
		})
	}

	return PBSegments
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

func UziToPBUziInfo(uziInfo *entity.Uzi) *pb.UziInfo {
	return &pb.UziInfo{
		Id:         uziInfo.Id.String(),
		Url:        uziInfo.Url,
		Projection: uziInfo.Projection,
		PatientId:  uziInfo.PatientID.String(),
		DeviceId:   int64(uziInfo.DeviceID),
	}
}

func PBUziToDTOUzi(req *pb.Uzi) *dto.Uzi {

	images := PBImagesToImages(req.Images, uuid.MustParse(req.UziInfo.Id))
	formations := PBFormationsToDTOFormations(req.Formations)
	segments := PBSegmentsToDTOSegments(req.Segments)

	return &dto.Uzi{
		UziInfo:    PBUziInfoToUzi(req.UziInfo),
		Images:     images,
		Formations: formations,
		Segments:   segments,
	}
}

func DTOUziToPBUzi(uzi *dto.Uzi) *pb.Uzi {
	return &pb.Uzi{
		UziInfo:    UziToPBUziInfo(uzi.UziInfo),
		Images:     ImagesToPBImages(uzi.Images),
		Formations: DTOFormationsToPBFormations(uzi.Formations),
		Segments:   DTOSegmentsToPBSegments(uzi.Segments),
	}
}

func DTOImageWithSegmentsToPBImageWithSegments(imageWithSegments *dto.ImageWithSegmentsFormations) *pb.ImageWithSegments {
	image := ImagesToPBImages([]entity.Image{*imageWithSegments.Image})[0]
	formations := DTOFormationsToPBFormations(imageWithSegments.Formations)
	segmetns := DTOSegmentsToPBSegments(imageWithSegments.Segments)

	return &pb.ImageWithSegments{
		Image:      image,
		Formations: formations,
		Segments:   segmetns,
	}
}

func PBFormationWithSegmentsToDTOFormationWithSegments(formationWithSegments *pb.FormationWithSegments) *dto.FormationWithSegments {
	formation := PBFormationsToDTOFormations([]*pb.Formation{formationWithSegments.Formation})[0]
	segments := PBSegmentsToDTOSegments(formationWithSegments.Segments)

	return &dto.FormationWithSegments{
		Formation: &formation,
		Segments:  segments,
	}
}

func DTOFormationWithSegmentsToPBFormationWithSegments(formationWithSegments *dto.FormationWithSegments) *pb.FormationWithSegments {
	formation := DTOFormationsToPBFormations([]dto.Formation{*formationWithSegments.Formation})[0]
	segments := DTOSegmentsToPBSegments(formationWithSegments.Segments)

	return &pb.FormationWithSegments{
		Formation: formation,
		Segments:  segments,
	}
}

func DevicesToPBDevices(devices []entity.Device) []*pb.Device {
	PBDevices := make([]*pb.Device, 0, len(devices))
	for _, v := range devices {
		PBDevices = append(PBDevices, &pb.Device{
			Id:   int64(v.Id),
			Name: v.Name,
		})
	}

	return PBDevices
}
