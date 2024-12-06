package doctor

import (
	"context"

	"med/internal/domain"
	pb "med/internal/generated/grpc/service"
	"med/internal/services/doctor"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DoctorHandler interface {
	RegisterDoctor(ctx context.Context, in *pb.RegisterDoctorIn) (*empty.Empty, error)
	GetDoctor(ctx context.Context, in *pb.GetDoctorIn) (*pb.GetDoctorOut, error)
	UpdateDoctor(ctx context.Context, in *pb.UpdateDoctorIn) (*pb.UpdateDoctorOut, error)
}

type handler struct {
	doctorSrv doctor.Service
}

func New(
	doctorSrv doctor.Service,
) DoctorHandler {
	return &handler{
		doctorSrv: doctorSrv,
	}
}

func (h *handler) RegisterDoctor(ctx context.Context, in *pb.RegisterDoctorIn) (*empty.Empty, error) {
	if err := h.doctorSrv.RegisterDoctor(ctx, domain.Doctor{
		Id:       uuid.MustParse(in.Doctor.Id),
		FullName: in.Doctor.Fullname,
		Org:      in.Doctor.Org,
		Job:      in.Doctor.Job,
		Desc:     in.Doctor.Desc,
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &empty.Empty{}, nil
}

func (h *handler) GetDoctor(ctx context.Context, in *pb.GetDoctorIn) (*pb.GetDoctorOut, error) {
	doctor, err := h.doctorSrv.GetDoctor(ctx, uuid.MustParse(in.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.GetDoctorOut{Doctor: domainToPb(&doctor)}, nil
}

func (h *handler) UpdateDoctor(ctx context.Context, in *pb.UpdateDoctorIn) (*pb.UpdateDoctorOut, error) {
	doctor, err := h.doctorSrv.UpdateDoctor(ctx, uuid.MustParse(in.Id), doctor.UpdateDoctor{
		Org:  in.Org,
		Job:  in.Job,
		Desc: in.Desc,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.UpdateDoctorOut{Doctor: domainToPb(&doctor)}, nil
}
