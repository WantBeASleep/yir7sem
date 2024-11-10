package medservice

import (
	"context"
	"yir/gateway/internal/custom"
	"yir/gateway/internal/entity/patientmodel"
	"yir/gateway/internal/pb/medpb"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (medsrv *MedService) AddPatient(ctx context.Context, Info *patientmodel.PatientInformation) error {
	patient := &medpb.Patient{
		Id:            Info.Patient.ID,
		FirstName:     Info.Patient.FirstName,
		LastName:      Info.Patient.LastName,
		FatherName:    Info.Patient.FatherName,
		MedicalPolicy: Info.Patient.MedicalPolicy,
		Email:         Info.Patient.Email,
		IsActive:      Info.Patient.IsActive,
	}
	// Надо подумать над рефакторингом
	// Нужен ли он или оставить так как есть
	patientCard := &medpb.Card{ // это другой тип, не  PostCard! из пациентов
		HasNodules:  Info.Card.HasNodules,
		Diagnosis:   Info.Card.Diagnosis,
		Patient:     patient,
		MedWorkerId: Info.Card.MedWorkerID,
	}
	req := &medpb.CreatePatientRequest{
		Patient:     patient,
		PatientCard: patientCard,
	}
	_, err := medsrv.PatientClient.AddPatient(ctx, req)
	if err != nil {
		custom.Logger.Error(
			"failed adding patient",
			zap.Error(err),
		)
		return err
	}
	return nil
}
func (medsrv *MedService) UpdatePatient(ctx context.Context, Info *patientmodel.PatientInformation) error {
	patient := &medpb.Patient{
		Id:            Info.Patient.ID,
		FirstName:     Info.Patient.FirstName,
		LastName:      Info.Patient.LastName,
		FatherName:    Info.Patient.FatherName,
		MedicalPolicy: Info.Patient.MedicalPolicy,
		Email:         Info.Patient.Email,
		IsActive:      Info.Patient.IsActive,
	}
	// Надо подумать над рефакторингом
	// Нужен ли он или оставить так как есть
	patientCard := &medpb.Card{ // это другой тип, не  PostCard! из пациентов
		HasNodules:  Info.Card.HasNodules,
		Diagnosis:   Info.Card.Diagnosis,
		Patient:     patient,
		MedWorkerId: Info.Card.MedWorkerID,
	}
	req := &medpb.PatientUpdateRequest{
		Patient:     patient,
		PatientCard: patientCard,
	}
	_, err := medsrv.PatientClient.UpdatePatient(ctx, req)
	if err != nil {
		custom.Logger.Error(
			"failed updating patient",
			zap.Error(err),
		)
		return err
	}
	return nil
}
func (medsrv *MedService) GetPatientList(ctx context.Context) ([]patientmodel.Patient, error) {
	data, err := medsrv.PatientClient.GetPatientList(ctx, &emptypb.Empty{})
	if err != nil {
		custom.Logger.Error(
			"faild get patient list",
			zap.Error(err),
		)
		return nil, err
	}
	resp := make([]patientmodel.Patient, 0, 1)
	for _, v := range data.Patients {
		resp = append(resp,
			patientmodel.Patient{
				ID:            v.Id,
				FirstName:     v.FirstName,
				LastName:      v.LastName,
				FatherName:    v.FatherName,
				MedicalPolicy: v.MedicalPolicy,
				Email:         v.Email,
				IsActive:      v.IsActive,
			},
		)
	}
	return resp, nil
}
func (medsrv *MedService) GetPatientInfoByID(ctx context.Context, ID uint64) (*patientmodel.PatientInformation, error) {
	req := &medpb.PatientInfoRequest{
		Id: ID,
	}
	data, err := medsrv.PatientClient.GetPatientInfoByID(ctx, req)
	if err != nil {
		custom.Logger.Error(
			"faild get patient by id",
			zap.Error(err),
		)
	}
	patient := &patientmodel.Patient{
		ID:            data.Patient.Id,
		FirstName:     data.Patient.FirstName,
		LastName:      data.Patient.LastName,
		FatherName:    data.Patient.FatherName,
		MedicalPolicy: data.Patient.MedicalPolicy,
		Email:         data.Patient.Email,
		IsActive:      data.Patient.IsActive,
	}
	patientCard := &patientmodel.PatientCard{
		ID:              data.PatientCard.Id,
		HasNodules:      data.PatientCard.HasNodules,
		Diagnosis:       data.PatientCard.Diagnosis,
		AppointmentTime: data.PatientCard.AppointmentTime,
		MedWorkerID:     data.PatientCard.MedWorkerId,
		PatientID:       data.PatientCard.Patient.Id,
	}
	resp := &patientmodel.PatientInformation{
		Patient: patient,
		Card:    patientCard,
	}
	return resp, nil
}
