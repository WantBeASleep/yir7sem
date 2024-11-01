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

func CreateUziReqToUzi(uzi *pb.CreateUziRequest) *entity.Uzi {
	return &entity.Uzi{
		Url:        uzi.Uzi.Url,
		Projection: uzi.Uzi.Projection,
		PatientID:  uuid.MustParse(uzi.Uzi.PatientId),
		DeviceID:   int(uzi.Uzi.DeviceId),
	}
}