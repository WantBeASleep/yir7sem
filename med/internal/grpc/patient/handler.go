package patient

import (
	"context"

	"github.com/WantBeASleep/goooool/gtclib"

	"med/internal/domain"
	pb "med/internal/generated/grpc/service"
	"med/internal/services/patient"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PatientHandler interface {
	CreatePatient(ctx context.Context, in *pb.CreatePatientIn) (*pb.CreatePatientOut, error)
	GetPatient(ctx context.Context, in *pb.GetPatientIn) (*pb.GetPatientOut, error)
	GetDoctorPatients(ctx context.Context, in *pb.GetDoctorPatientsIn) (*pb.GetDoctorPatientsOut, error)
	UpdatePatient(ctx context.Context, in *pb.UpdatePatientIn) (*pb.UpdatePatientOut, error)
}

type handler struct {
	patientSrv patient.Service
}

func New(
	patientSrv patient.Service,
) PatientHandler {
	return &handler{
		patientSrv: patientSrv,
	}
}

func (h *handler) CreatePatient(ctx context.Context, in *pb.CreatePatientIn) (*pb.CreatePatientOut, error) {
	id, err := h.patientSrv.CreatePatient(ctx, domain.Patient{
		FullName:   in.Fullname,
		Email:      in.Email,
		Policy:     in.Policy,
		Active:     in.Active,
		Malignancy: in.Malignancy,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.CreatePatientOut{Id: id.String()}, nil
}

func (h *handler) GetPatient(ctx context.Context, in *pb.GetPatientIn) (*pb.GetPatientOut, error) {
	patient, err := h.patientSrv.GetPatient(ctx, uuid.MustParse(in.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.GetPatientOut{Patient: domainToPb(&patient)}, nil
}

func (h *handler) GetDoctorPatients(ctx context.Context, in *pb.GetDoctorPatientsIn) (*pb.GetDoctorPatientsOut, error) {
	patients, err := h.patientSrv.GetDoctorPatients(ctx, uuid.MustParse(in.DoctorId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	res := make([]*pb.Patient, 0, len(patients))
	for _, v := range patients {
		res = append(res, domainToPb(&v))
	}

	return &pb.GetDoctorPatientsOut{Patients: res}, nil
}

func (h *handler) UpdatePatient(ctx context.Context, in *pb.UpdatePatientIn) (*pb.UpdatePatientOut, error) {
	patient, err := h.patientSrv.UpdatePatient(
		ctx,
		uuid.MustParse(in.DoctorId),
		uuid.MustParse(in.Id),
		patient.UpdatePatient{
			Active:      in.Active,
			Malignancy:  in.Malignancy,
			LastUziDate: gtclib.Timestamp.ToTimePointer(in.LastUziDate),
		},
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.UpdatePatientOut{Patient: domainToPb(&patient)}, nil
}
