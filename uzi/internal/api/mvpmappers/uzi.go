package mvpmappers

import (
	pb "yir/uzi/api/grpcapi"
	"yir/uzi/internal/entity"

	"github.com/google/uuid"
)

func UziToPBUziResp(uzi *entity.Uzi) *pb.UziReponse {
	return &pb.UziReponse{
		Id:         uzi.Id.String(),
		Url:        uzi.Url,
		Projection: uzi.Projection,
		PatientId:  uzi.PatientID.String(),
		DeviceId:   int64(uzi.DeviceID),
	}
}

func PBUziReqToUzi(uzi *pb.UziRequest) *entity.Uzi {
	return &entity.Uzi{
		Url:        uzi.Url,
		Projection: uzi.Projection,
		PatientID:  uuid.MustParse(uzi.PatientId),
		DeviceID:   int(uzi.DeviceId),
	}
}

func PBUpdateUziReqToUzi(uzi *pb.UziUpdateRequest) *entity.Uzi {
	return &entity.Uzi{
		Projection: uzi.Projection,
		PatientID:  uuid.MustParse(uzi.PatientId),
		DeviceID:   int(uzi.DeviceId),
	}
}
