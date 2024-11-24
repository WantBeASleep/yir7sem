package mvpmappers

import (
	"encoding/json"
	"yir/uzi/api/broker"
	kafka "yir/uzi/api/broker"
	pb "yir/uzi/api/grpcapi"
	"yir/uzi/internal/usecases/dto"

	"github.com/google/uuid"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func PBPointsToString(contor []*pb.Point) string {
	res := make([]Point, 0, len(contor))
	for _, cn := range contor {
		res = append(res, Point{
			X: int(cn.X),
			Y: int(cn.Y),
		})
	}

	JSON, _ := json.Marshal(res)
	return string(JSON)
}

func KafkaPointsToString(contor []*broker.Point) string {
	res := make([]Point, 0, len(contor))
	for _, cn := range contor {
		res = append(res, Point{
			X: int(cn.X),
			Y: int(cn.Y),
		})
	}

	JSON, _ := json.Marshal(res)
	return string(JSON)
}

func StringToContor(contor string) []*pb.Point {
	var contorJSON []Point
	json.Unmarshal([]byte(contor), &contorJSON)

	res := make([]*pb.Point, 0, len(contorJSON))
	for _, cn := range contorJSON {
		res = append(res, &pb.Point{
			X: int64(cn.X),
			Y: int64(cn.Y),
		})
	}

	return res
}

func PBSegmentReqToDTOSegment(segment *pb.SegmentRequest) *dto.Segment {
	return &dto.Segment{
		ImageID:     uuid.MustParse(segment.ImageId),
		FormationID: uuid.MustParse(segment.FormationId),
		Contor:      PBPointsToString(segment.Contor),
		Tirads:      PBTiradsToTirads(segment.Tirads),
	}
}

func KafkaSegmentReqToDTOSegment(segment *kafka.KafkaSegment) *dto.Segment {
	return &dto.Segment{
		Id:          uuid.MustParse(segment.Id),
		ImageID:     uuid.MustParse(segment.ImageId),
		FormationID: uuid.MustParse(segment.FormationId),
		Contor:      KafkaPointsToString(segment.Contor),
		Tirads:      KafkaTiradsToTirads(segment.Tirads),
	}
}

func PBSegmentNestedReqToDTOSegment(segment *pb.SegmentNestedRequest) *dto.Segment {
	return &dto.Segment{
		ImageID: uuid.MustParse(segment.ImageId),
		Contor:  PBPointsToString(segment.Contor),
		Tirads:  PBTiradsToTirads(segment.Tirads),
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

		Contor: StringToContor(segment.Contor),
		Tirads: TiradsToPBTirads(segment.Tirads),
	}
}

func DTOSegmentsToPBSegmentsResp(segments []dto.Segment) []*pb.SegmentResponse {
	PBSegments := make([]*pb.SegmentResponse, 0, len(segments))
	for _, segment := range segments {
		PBSegments = append(PBSegments, DTOSegmentToPBSegmentResp(segment))
	}

	return PBSegments
}
