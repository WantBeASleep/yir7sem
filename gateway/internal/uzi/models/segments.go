package models

import (
	pb "yir/gateway/rpc/uzi"

	"github.com/google/uuid"
)

type SegmentResp struct {
	ID          uuid.UUID `json:"id"`
	ImageID     uuid.UUID `json:"image_id"`
	FormationID uuid.UUID `json:"formation_id"`

	Contor []Point `json:"contor"`
	Tirads Tirads  `json:"tirads"`
}

func PBSegmentToSegmentResp(in *pb.SegmentResponse, contor []Point) SegmentResp {
	return SegmentResp{
		ID:          uuid.MustParse(in.Id),
		ImageID:     uuid.MustParse(in.ImageId),
		FormationID: uuid.MustParse(in.FormationId),

		Contor: contor,
		Tirads: PBTiradsToTirads(in.Tirads),
	}
}

// TODO: починить это в .proto на имя SegmentNestedToFormation
type SegmentNestedReq struct {
	ImageID uuid.UUID `json:"image_id"`

	Contor []Point `json:"contor"`
	Tirads Tirads  `json:"tirads"`
}

func SegmentsNestedToPB(in []SegmentNestedReq) []*pb.SegmentNestedRequest {
	res := make([]*pb.SegmentNestedRequest, 0, len(in))
	for _, seg := range in {
		res = append(res, &pb.SegmentNestedRequest{
			ImageId: seg.ImageID.String(),
			Contor: ContorToPB(seg.Contor),
			Tirads: TiradsToPB(&seg.Tirads),
		})
	}

	return res
}
