package uzi

import (
	"context"

	"uzi/internal/domain"
	pb "uzi/internal/generated/grpc/service"
	"uzi/internal/services/uzi"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UziHandler interface {
	CreateUzi(ctx context.Context, req *pb.CreateUziIn) (*pb.CreateUziOut, error)
	GetUzi(ctx context.Context, in *pb.GetUziIn) (*pb.GetUziOut, error)
	GetPatientUzis(ctx context.Context, in *pb.GetPatientUzisIn) (*pb.GetPatientUzisOut, error)
	GetEchographic(ctx context.Context, in *pb.GetEchographicIn) (*pb.GetEchographicOut, error)
	UpdateUzi(ctx context.Context, req *pb.UpdateUziIn) (*pb.UpdateUziOut, error)
	UpdateEchographic(ctx context.Context, in *pb.UpdateEchographicIn) (*pb.UpdateEchographicOut, error)
}

type handler struct {
	uziSrv uzi.Service
}

func New(
	uziSrv uzi.Service,
) UziHandler {
	return &handler{
		uziSrv: uziSrv,
	}
}

func (h *handler) CreateUzi(ctx context.Context, in *pb.CreateUziIn) (*pb.CreateUziOut, error) {
	uuid, err := h.uziSrv.CreateUzi(ctx, domain.Uzi{
		Projection: in.Projection,
		PatientID:  uuid.MustParse(in.PatientId),
		DeviceID:   int(in.DeviceId),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.CreateUziOut{Id: uuid.String()}, nil
}

func (h *handler) GetEchographic(ctx context.Context, in *pb.GetEchographicIn) (*pb.GetEchographicOut, error) {
	echographic, err := h.uziSrv.GetUziEchographicsByID(ctx, uuid.MustParse(in.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	pbEchographic := domainEchographicToPb(&echographic)

	return &pb.GetEchographicOut{
		Echographic: pbEchographic,
	}, nil
}

func (h *handler) GetUzi(ctx context.Context, in *pb.GetUziIn) (*pb.GetUziOut, error) {
	uzi, err := h.uziSrv.GetUziByID(ctx, uuid.MustParse(in.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	pbUzi := domainUziToPbUzi(&uzi)

	return &pb.GetUziOut{
		Uzi: pbUzi,
	}, nil
}

func (h *handler) GetPatientUzis(ctx context.Context, in *pb.GetPatientUzisIn) (*pb.GetPatientUzisOut, error) {
	uzis, err := h.uziSrv.GetUzisByPatientID(ctx, uuid.MustParse(in.PatientId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	resp := make([]*pb.Uzi, 0, len(uzis))
	for _, v := range uzis {
		resp = append(resp, domainUziToPbUzi(&v))
	}

	return &pb.GetPatientUzisOut{Uzis: resp}, nil
}

func (h *handler) UpdateUzi(ctx context.Context, in *pb.UpdateUziIn) (*pb.UpdateUziOut, error) {
	uzi, err := h.uziSrv.UpdateUzi(ctx,
		uuid.MustParse(in.Id),
		uzi.UpdateUzi{
			Projection: in.Projection,
			Checked:    in.Checked,
		},
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.UpdateUziOut{
		Uzi: domainUziToPbUzi(&uzi),
	}, nil
}

func (h *handler) UpdateEchographic(ctx context.Context, in *pb.UpdateEchographicIn) (*pb.UpdateEchographicOut, error) {
	echographic, err := h.uziSrv.UpdateEchographic(
		ctx,
		uuid.MustParse(in.Echographic.Id),
		uzi.UpdateEchographic{
			Contors:         in.Echographic.Contors,
			LeftLobeLength:  in.Echographic.LeftLobeLength,
			LeftLobeWidth:   in.Echographic.LeftLobeWidth,
			LeftLobeThick:   in.Echographic.LeftLobeThick,
			LeftLobeVolum:   in.Echographic.LeftLobeVolum,
			RightLobeLength: in.Echographic.RightLobeLength,
			RightLobeWidth:  in.Echographic.RightLobeWidth,
			RightLobeThick:  in.Echographic.RightLobeThick,
			RightLobeVolum:  in.Echographic.RightLobeVolum,
			GlandVolum:      in.Echographic.GlandVolum,
			Isthmus:         in.Echographic.Isthmus,
			Struct:          in.Echographic.Struct,
			Echogenicity:    in.Echographic.Echogenicity,
			RegionalLymph:   in.Echographic.RegionalLymph,
			Vascularization: in.Echographic.Vascularization,
			Location:        in.Echographic.Location,
			Additional:      in.Echographic.Additional,
			Conclusion:      in.Echographic.Conclusion,
		})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.UpdateEchographicOut{
		Echographic: domainEchographicToPb(&echographic),
	}, nil
}
