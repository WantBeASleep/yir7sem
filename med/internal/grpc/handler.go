package grpc

import (
	"yirv2/med/internal/generated/grpc/service"
	"yirv2/med/internal/grpc/card"
	"yirv2/med/internal/grpc/doctor"
	"yirv2/med/internal/grpc/patient"
)

type Handler struct {
	patient.PatientHandler
	doctor.DoctorHandler
	card.CardHandler

	service.UnsafeMedSrvServer
}

func New(
	patientHandler patient.PatientHandler,
	doctorHandler doctor.DoctorHandler,
	cardHandler card.CardHandler,
) *Handler {
	return &Handler{
		PatientHandler: patientHandler,
		DoctorHandler:  doctorHandler,
		CardHandler:    cardHandler,
	}
}
