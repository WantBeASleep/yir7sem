package models

import (
	pb "yir/gateway/rpc/uzi"

	"github.com/google/uuid"
)

type FormationResp struct {
	Id     uuid.UUID `json:"id"`
	Tirads Tirads    `json:"tirads_id"`
	Ai     bool      `json:"ai"`
}

func PBFormationRespToFormationResp(in *pb.FormationResponse) FormationResp {
	return FormationResp{
		Id:     uuid.MustParse(in.Id),
		Tirads: PBTiradsToTirads(in.Tirads),
		Ai:     in.Ai,
	}
}

// TODO: посмотреть, возможно не нужно(id будет в path)
type FormationReq struct {
	Tirads Tirads `json:"tirads_id"`
	Ai     bool   `json:"ai"`
}

func FormationReqToPB(in *FormationReq) *pb.FormationRequest {
	return &pb.FormationRequest{
		Tirads: TiradsToPB(&in.Tirads),
		Ai: in.Ai,
	}
}

type FormationWithSegments struct {
	Formation FormationResp `json:"formation"`
	Segments  []SegmentResp `json:"segments"`
}

type FormationWithNestedSegmentsReq struct {
	Segmets []SegmentNestedReq `json:"segments"`
	Tirads  Tirads             `json:"tirads_id"`
	Ai      bool               `json:"ai"`
}

func FormationWithNestedSegmentsReqToPB(in *FormationWithNestedSegmentsReq) *pb.FormationWithNestedSegmentsRequest {
	return &pb.FormationWithNestedSegmentsRequest{
		Segments: SegmentsNestedToPB(in.Segmets),
		Tirads:   TiradsToPB(&in.Tirads),
		Ai:       in.Ai,
	}
}

type FormationWithSegmentsIDs struct {
	Formation uuid.UUID  `json:"formation"`
	Segments  uuid.UUIDs `json:"segments"`
}

func CreateFormationWithSegmentsRespToFormationWithSegmentsIDs(in *pb.CreateFormationWithSegmentsResponse) FormationWithSegmentsIDs {
	segments := make(uuid.UUIDs, 0, len(in.SegmentsIds))
	for _, s := range in.SegmentsIds {
		segments = append(segments, uuid.MustParse(s))
	}
	return FormationWithSegmentsIDs{
		Formation: uuid.MustParse(in.FormationId),
		Segments:  segments,
	}
}
