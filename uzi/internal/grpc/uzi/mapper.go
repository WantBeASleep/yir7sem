package uzi

import (
	"time"

	"uzi/internal/domain"
	pb "uzi/internal/generated/grpc/service"
)

func domainUziToPbUzi(d *domain.Uzi) *pb.Uzi {
	if d == nil {
		return nil
	}

	return &pb.Uzi{
		Id:         d.Id.String(),
		Projection: d.Projection,
		Checked:    d.Checked,
		PatientId:  d.PatientID.String(),
		DeviceId:   int64(d.DeviceID),
		CreateAt:   d.CreateAt.Format(time.RFC3339),
	}
}

func domainEchographicToPb(d *domain.Echographic) *pb.Echographic {
	if d == nil {
		return nil
	}

	return &pb.Echographic{
		Id:              d.Id.String(),
		Contors:         d.Contors,
		LeftLobeLength:  d.LeftLobeLength,
		LeftLobeWidth:   d.LeftLobeWidth,
		LeftLobeThick:   d.LeftLobeThick,
		LeftLobeVolum:   d.LeftLobeVolum,
		RightLobeLength: d.RightLobeLength,
		RightLobeWidth:  d.RightLobeWidth,
		RightLobeThick:  d.RightLobeThick,
		RightLobeVolum:  d.RightLobeVolum,
		GlandVolum:      d.GlandVolum,
		Isthmus:         d.Isthmus,
		Struct:          d.Struct,
		Echogenicity:    d.Echogenicity,
		RegionalLymph:   d.RegionalLymph,
		Vascularization: d.Vascularization,
		Location:        d.Location,
		Additional:      d.Additional,
		Conclusion:      d.Conclusion,
	}
}
