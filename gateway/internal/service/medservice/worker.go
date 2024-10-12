package medservice

import (
	"context"
	"yir/gateway/internal/custom"
	"yir/gateway/internal/entity/patientmodel"
	"yir/gateway/internal/entity/workermodel"
	"yir/gateway/internal/pb/medpb"

	"go.uber.org/zap"
)

func (medsrv *MedService) GetMedWorkers(ctx context.Context, limit, offset uint64) (*workermodel.MedicalWorkerList, error) {
	req := &medpb.GetMedworkerRequest{
		Limit:  limit,
		Offset: offset,
	}
	data, err := medsrv.WorkerClient.GetMedWorkers(ctx, req)
	if err != nil {
		custom.Logger.Error(
			"failed to get workers",
			zap.Error(err),
		)
		return nil, err
	}
	resp := &workermodel.MedicalWorkerList{
		Workers: make([]workermodel.MedicalWorker, 0, 1),
		Count:   data.Count,
	}
	for _, v := range data.Results {
		resp.Workers = append(resp.Workers,
			workermodel.MedicalWorker{
				ID:              v.Id,
				FirstName:       v.FirstName,
				MiddleName:      v.MiddleName,
				LastName:        v.LastName,
				MedOrganization: v.MedOrganization,
				Job:             v.Job,
				IsRemoteWorker:  v.IsRemoteWorker,
				ExpertDetails:   v.ExpertDetails,
			},
		)
	}
	return resp, nil
}
func (medsrv *MedService) GetMedWorkerByID(ctx context.Context, id uint64) (*workermodel.MedicalWorker, error) {
	req := &medpb.GetMedMedWorkerByIDRequest{
		Id: id,
	}
	data, err := medsrv.WorkerClient.GetMedWorkerByID(ctx, req)
	if err != nil {
		custom.Logger.Error(
			"failed to get worker by id",
			zap.Error(err),
		)
		return nil, err
	}
	resp := &workermodel.MedicalWorker{
		ID:              data.Worker.Id,
		FirstName:       data.Worker.FirstName,
		MiddleName:      data.Worker.MiddleName,
		LastName:        data.Worker.LastName,
		MedOrganization: data.Worker.MedOrganization,
		Job:             data.Worker.Job,
		IsRemoteWorker:  data.Worker.IsRemoteWorker,
		ExpertDetails:   data.Worker.ExpertDetails,
	}
	return resp, nil
}
func (medsrv *MedService) UpdateMedWorker(ctx context.Context, id uint64, in *workermodel.MedicalWorkerUpdateRequest) (*workermodel.MedicalWorker, error) {
	req := &medpb.UpdateMedWorkerRequest{
		FirstName:      in.FirstName,
		LastName:       in.LastName,
		MiddleName:     in.MiddleName,
		Job:            in.Job,
		IsRemoteWorker: in.IsRemoteWorker,
		ExpertDetails:  in.ExpertDetails,
	}
	data, err := medsrv.WorkerClient.UpdateMedWorker(ctx, req)
	if err != nil {
		custom.Logger.Error(
			"failed to update worker",
			zap.Error(err),
		)
		return nil, err
	}
	resp := &workermodel.MedicalWorker{
		ID:             data.Worker.Id,
		FirstName:      data.Worker.FirstName,
		LastName:       data.Worker.LastName,
		MiddleName:     data.Worker.MiddleName,
		Job:            data.Worker.Job,
		IsRemoteWorker: data.Worker.IsRemoteWorker,
		ExpertDetails:  data.Worker.ExpertDetails,
	}
	return resp, nil
}
func (medsrv *MedService) AddMedWorker(ctx context.Context, in *workermodel.AddMedicalWorkerRequest) (*workermodel.MedicalWorker, error) {
	req := &medpb.AddMedWorkerRequest{
		FirstName:      in.FirstName,
		LastName:       in.LastName,
		MiddleName:     in.MiddleName,
		Job:            in.Job,
		IsRemoteWorker: in.IsRemoteWorker,
		ExpertDetails:  in.ExpertDetails,
	}
	data, err := medsrv.WorkerClient.AddMedWorker(ctx, req)
	if err != nil {
		custom.Logger.Error(
			"failed to add worker",
			zap.Error(err),
		)
		return nil, err
	}
	resp := &workermodel.MedicalWorker{
		ID:             data.Worker.Id,
		FirstName:      data.Worker.FirstName,
		LastName:       data.Worker.LastName,
		MiddleName:     data.Worker.MiddleName,
		Job:            data.Worker.Job,
		IsRemoteWorker: data.Worker.IsRemoteWorker,
		ExpertDetails:  data.Worker.ExpertDetails,
	}
	return resp, nil
}
func (medsrv *MedService) GetPatientsByMedWorker(ctx context.Context, id uint64) (*workermodel.MedicalWorkerWithPatients, error) {
	req := &medpb.GetPatientsByMedWorkerRequest{
		MedWorkerId: id,
	}
	data, err := medsrv.WorkerClient.GetPatientsByMedWorker(ctx, req)
	if err != nil {
		custom.Logger.Error(
			"failed to get patients of worker",
			zap.Error(err),
		)
		return nil, err
	}
	resp := &workermodel.MedicalWorkerWithPatients{
		// MedWorker: data.MedWorkerId ??????
		Patients: make([]patientmodel.PatientCard, 0, 1),
	}
	for _, v := range data.Cards {
		resp.Patients = append(resp.Patients,
			patientmodel.PatientCard{
				ID:              v.Id,
				AppointmentTime: v.AppointmentTime,
				HasNodules:      v.HasNodules,
				Diagnosis:       v.Diagnosis,
				MedWorkerID:     v.MedWorkerId,
				PatientID:       v.Patient.Id,
			},
		)
	}
	return resp, nil
}
