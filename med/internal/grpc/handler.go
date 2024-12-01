package grpc

import (
	"yir/med/internal/generated/grpc/service"
	"yir/med/internal/grpc/card"
	"yir/med/internal/grpc/doctor"
	"yir/med/internal/grpc/patient"
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
