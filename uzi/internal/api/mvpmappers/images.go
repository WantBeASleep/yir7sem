package mvpmappers

import (
	pb "yir/uzi/api/grpcapi"
	"yir/uzi/internal/entity"
	"yir/uzi/internal/usecases/dto"
)

func ImageToPBImageResp(image *entity.Image) *pb.ImageResponse {
	return &pb.ImageResponse{
		Id:   image.Id.String(),
		Url:  image.Url,
		Page: int64(image.Page),
	}
}

func ImagesToPBImagesResp(images []entity.Image) []*pb.ImageResponse {
	PBImages := make([]*pb.ImageResponse, 0, len(images))
	for _, image := range images {
		PBImages = append(PBImages, ImageToPBImageResp(&image))
	}

	return PBImages
}

func DTOImageWithFormationsSegmentsToPBImageWithFormationsSegments(imageWithFormationsSegments *dto.ImageWithFormationsSegments) *pb.ImageWithFormationsSegments {
	image := ImageToPBImageResp(imageWithFormationsSegments.Image)
	formations := DTOFormationsToPBFormationsResp(imageWithFormationsSegments.Formations)
	segmetns := DTOSegmentsToPBSegmentsResp(imageWithFormationsSegments.Segments)

	return &pb.ImageWithFormationsSegments{
		Image:      image,
		Formations: formations,
		Segments:   segmetns,
	}
}
