package grpc

import (
	"med/internal/generated/grpc/service"
	"med/internal/grpc/card"
	"med/internal/grpc/doctor"
	"med/internal/grpc/patient"
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
