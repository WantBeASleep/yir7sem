package mvpmappers

import (
	kafka "yir/uzi/api/broker"
	pb "yir/uzi/api/grpcapi"
	"yir/uzi/internal/usecases/dto"

	"github.com/google/uuid"
)

func PBSegmentReqToDTOSegment(segment *pb.SegmentRequest) *dto.Segment {
	return &dto.Segment{
		ImageID:     uuid.MustParse(segment.ImageId),
		FormationID: uuid.MustParse(segment.FormationId),
		ContorURL:   segment.ContorUrl,
		Tirads:      PBTiradsToTirads(segment.Tirads),
	}
}

func KafkaSegmentReqToDTOSegment(segment *kafka.KafkaSegment) *dto.Segment {
	return &dto.Segment{
		Id:          uuid.MustParse(segment.Id),
		ImageID:     uuid.MustParse(segment.ImageId),
		FormationID: uuid.MustParse(segment.FormationId),
		ContorURL:   segment.ContorUrl,
		Tirads:      KafkaTiradsToTirads(segment.Tirads),
	}
}

func PBSegmentNestedReqToDTOSegment(segment *pb.SegmentNestedRequest) *dto.Segment {
	return &dto.Segment{
		ImageID:   uuid.MustParse(segment.ImageId),
		ContorURL: segment.ContorUrl,
		Tirads:    PBTiradsToTirads(segment.Tirads),
	}
}

func PBSegmentsReqToDTOSegments(segments []*pb.SegmentRequest) []dto.Segment {
	dtoSegments := make([]dto.Segment, 0, len(segments))
	for _, segment := range segments {
		dtoSegments = append(dtoSegments, *PBSegmentReqToDTOSegment(segment))
	}

	return dtoSegments
}

func KafkaSegmentsToDTOSegments(segments []*kafka.KafkaSegment) []dto.Segment {
	dtoSegments := make([]dto.Segment, 0, len(segments))
	for _, segment := range segments {
		dtoSegments = append(dtoSegments, *KafkaSegmentReqToDTOSegment(segment))
	}

	return dtoSegments
}

func PBSegmentsNestedReqToDTOSegments(segments []*pb.SegmentNestedRequest) []dto.Segment {
	dtoSegments := make([]dto.Segment, 0, len(segments))
	for _, segment := range segments {
		dtoSegments = append(dtoSegments, *PBSegmentNestedReqToDTOSegment(segment))
	}

	return dtoSegments
}

func DTOSegmentToPBSegmentResp(segment dto.Segment) *pb.SegmentResponse {
	return &pb.SegmentResponse{
		Id:          segment.Id.String(),
		FormationId: segment.FormationID.String(),
		ImageId:     segment.ImageID.String(),

		ContorUrl: segment.ContorURL,
		Tirads:    TiradsToPBTirads(segment.Tirads),
	}
}

func DTOSegmentsToPBSegmentsResp(segments []dto.Segment) []*pb.SegmentResponse {
	PBSegments := make([]*pb.SegmentResponse, 0, len(segments))
	for _, segment := range segments {
		PBSegments = append(PBSegments, DTOSegmentToPBSegmentResp(segment))
	}

	return PBSegments
}
