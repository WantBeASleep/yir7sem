//TODO: Сделать магамеду

package med

import (
	"context"
	"log"
	pb "yir/gateway/rpc/med"

	empty "github.com/golang/protobuf/ptypes/empty"
)

type Ctrl struct {
	pb.UnimplementedMedCardServer
	pb.UnimplementedMedPatientServer
	pb.UnimplementedMedWorkersServer

	cardClient    pb.MedCardClient
	patientClient pb.MedPatientClient
	workerClient  pb.MedWorkersClient
}

func NewCtrl(
	cardClient pb.MedCardClient,
	patientClient pb.MedPatientClient,
	workerClient pb.MedWorkersClient,
) *Ctrl {
	return &Ctrl{
		cardClient:    cardClient,
		patientClient: patientClient,
		workerClient:  workerClient,
	}
}

// GetCards godoc
//
//	@Summary		Get cards
//	@Description	Получить список карт
//	@Tags			Cards
//	@Accept			json
//	@Produce		json
//	@Param			body	body		med.GetCardsRequest	true	"Запрос"
//	@Success		200		{object}	med.GetCardsResponse
//	@Failure		500		{string}	string	"Internal error"
//	@Router			/med/card/list [get]
func (c *Ctrl) GetCards(ctx context.Context, req *pb.GetCardsRequest) (*pb.GetCardsResponse, error) {
	log.Println("Called GetCards")
	return c.cardClient.GetCards(ctx, req)
}

// PostCard godoc
//
//	@Summary		Add card
//	@Description	Создать новую карту
//	@Tags			Cards
//	@Accept			json
//	@Produce		json
//	@Param			body	body		med.PostCardRequest	true	"Запрос"
//	@Success		200		{object}	med.PostCardResponse
//	@Failure		500		{string}	string	"Internal error"
//	@Router			/med/card/create [post]
func (c *Ctrl) PostCard(ctx context.Context, req *pb.PostCardRequest) (*pb.PostCardResponse, error) {
	log.Println("Called PostCard")
	return c.cardClient.PostCard(ctx, req)
}

// GetCardByID godoc
//
//	@Summary		Get card by ID
//	@Description	Получить карту по ID
//	@Tags			Cards
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID карты"
//	@Success		200	{object}	med.GetCardByIDResponse
//	@Failure		500	{string}	string	"Internal error"
//	@Router			/med/card/by/{id} [get]
func (c *Ctrl) GetCardByID(ctx context.Context, req *pb.GetCardByIDRequest) (*pb.GetCardByIDResponse, error) {
	log.Println("Called GetCardByID")
	return c.cardClient.GetCardByID(ctx, req)
}

// PutCard godoc
//
//	@Summary		Update card
//	@Description	Обновить карту
//	@Tags			Cards
//	@Accept			json
//	@Produce		json
//	@Param			body	body		med.PutCardRequest	true	"Запрос"
//	@Success		200		{object}	med.PutCardResponse
//	@Failure		500		{string}	string	"Internal error"
//	@Router			/med/card/update/{id} [put]
func (c *Ctrl) PutCard(ctx context.Context, req *pb.PutCardRequest) (*pb.PutCardResponse, error) {
	log.Println("Called PutCard")
	return c.cardClient.PutCard(ctx, req)
}

// Не работает (В сваггер не включен)
//
//	@Summary		Patch card
//	@Description	Частично обновить карту
//	@Tags			Cards
//	@Accept			json
//	@Produce		json
//	@Param			body	body		med.PatchCardRequest	true	"Запрос"
//	@Success		200		{object}	med.UpdateMedWorkerResponse
//	@Failure		500		{string}	string	"Internal error"
//	@Router			/med/card/patch/{id} [patch]
func (c *Ctrl) PatchCard(ctx context.Context, req *pb.PatchCardRequest) (*pb.PatchCardResponse, error) {
	log.Println("Called PatchCard")
	return c.cardClient.PatchCard(ctx, req)
}

// DeleteCard godoc
//
//	@Summary		Delete card
//	@Description	Удалить карту
//	@Tags			Cards
//	@Accept			json
//	@Produce		json
//	@Param			body	body	med.DeleteCardRequest	true	"Запрос"
//	@Success		200
//	@Failure		500	{string}	string	"Internal error"
//	@Router			/med/card/delete/{id} [delete]
func (c *Ctrl) DeleteCard(ctx context.Context, req *pb.DeleteCardRequest) (*empty.Empty, error) {
	log.Println("Called DeleteCard")
	return c.cardClient.DeleteCard(ctx, req)
}

// AddPatient godoc
//
//	@Summary		Add patient
//	@Description	Добавить пациента
//	@Tags			Patients
//	@Accept			json
//	@Produce		json
//	@Param			body	body	med.CreatePatientRequest	true	"Запрос"
//	@Success		200
//	@Failure		500	{string}	string	"Internal error"
//	@Router			/med/patient/create [post]
func (c *Ctrl) AddPatient(ctx context.Context, req *pb.CreatePatientRequest) (*empty.Empty, error) {
	log.Println("Called AddPatient")
	return c.patientClient.AddPatient(ctx, req)
}

// GetPatientList godoc
//
//	@Summary		Get patient list
//	@Description	Получить список пациентов
//	@Tags			Patients
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	med.PatientListResponse
//	@Failure		500	{string}	string	"Internal error"
//	@Router			/med/patient/list [get]
func (c *Ctrl) GetPatientList(ctx context.Context, req *empty.Empty) (*pb.PatientListResponse, error) {
	log.Println("Called GetPatientList")
	return c.patientClient.GetPatientList(ctx, req)
}

// GetPatientInfoByID godoc
//
//	@Summary		Get patient by ID
//	@Description	Получить информацию о пациенте по ID
//	@Tags			Patients
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID пациента"
//	@Success		200	{object}	med.PatientInfoResponse
//	@Failure		500	{string}	string	"Internal error"
//	@Router			/med/patient/info/{id} [get]
func (c *Ctrl) GetPatientInfoByID(ctx context.Context, req *pb.PatientInfoRequest) (*pb.PatientInfoResponse, error) {
	log.Println("Called GetPatientList")
	return c.patientClient.GetPatientInfoByID(ctx, req)
}

// UpdatePatient godoc
//
//	@Summary		Update patient
//	@Description	Обновить информацию о пациенте
//	@Tags			Patients
//	@Accept			json
//	@Produce		json
//	@Param			body	body	med.PatientUpdateRequest	true	"Запрос"
//	@Success		200
//	@Failure		500	{string}	string	"Internal error"
//	@Router			/med/patient/update/{patient.id} [put]
func (c *Ctrl) UpdatePatient(ctx context.Context, req *pb.PatientUpdateRequest) (*empty.Empty, error) {
	log.Println("Called UpdatePatient")
	return c.patientClient.UpdatePatient(ctx, req)
}

// GetMedWorkers godoc
//
//	@Summary		Get med workers
//	@Description	Получить список медицинских работников
//	@Tags			MedWorkers
//	@Accept			json
//	@Produce		json
//	@Param			body	body		med.GetMedworkerRequest	true	"Запрос"
//	@Success		200		{object}	med.GetMedworkerListResponse
//	@Failure		500		{string}	string	"Internal error"
//	@Router			/med/worker/list [get]
func (c *Ctrl) GetMedWorkers(ctx context.Context, req *pb.GetMedworkerRequest) (*pb.GetMedworkerListResponse, error) {
	log.Println("Called GetMedWorkers")
	return c.workerClient.GetMedWorkers(ctx, req)
}

// UpdateMedWorker godoc
//
//	@Summary		Update med worker
//	@Description	Обновить данные медицинского работника
//	@Tags			MedWorkers
//	@Accept			json
//	@Produce		json
//	@Param			body	body		med.UpdateMedWorkerRequest	true	"Запрос"
//	@Success		200		{object}	med.UpdateMedWorkerResponse
//	@Failure		500		{string}	string	"Internal error"
//	@Router			/med/worker/update/{id} [put]
func (c *Ctrl) UpdateMedWorker(ctx context.Context, req *pb.UpdateMedWorkerRequest) (*pb.UpdateMedWorkerResponse, error) {
	log.Println("Called UpdateMedWorker")
	return c.workerClient.UpdateMedWorker(ctx, req)
}

// GetMedWorkerByID godoc
//
//	@Summary		Get med worker by ID
//	@Description	Получить данные медицинского работника по ID
//	@Tags			MedWorkers
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID медицинского работника"
//	@Success		200	{object}	med.GetMedWorkerByIDResponse
//	@Failure		500	{string}	string	"Internal error"
//	@Router			/med/worker/by/{id} [get]
func (c *Ctrl) GetMedWorkerByID(ctx context.Context, req *pb.GetMedMedWorkerByIDRequest) (*pb.GetMedWorkerByIDResponse, error) {
	log.Println("Called GetMedWorkerByID")
	return c.workerClient.GetMedWorkerByID(ctx, req)
}

// PatchMedWorker godoc
//
//	@Summary		Patch med worker
//	@Description	Частично обновить данные медицинского работника
//	@Tags			MedWorkers
//	@Accept			json
//	@Produce		json
//	@Param			body	body		med.PatchMedWorkerRequest	true	"Запрос"
//	@Success		200		{object}	med.UpdateMedWorkerResponse
//	@Failure		500		{string}	string	"Internal error"
//	@Router			/med/worker/patch/{id} [patch]
func (c *Ctrl) PatchMedWorker(ctx context.Context, req *pb.PatchMedWorkerRequest) (*pb.UpdateMedWorkerResponse, error) {
	log.Println("Called PatchMedWorker")
	return c.workerClient.PatchMedWorker(ctx, req)
}

// AddMedWorker godoc
//
//	@Summary		Add med worker
//	@Description	Добавить нового медицинского работника
//	@Tags			MedWorkers
//	@Accept			json
//	@Produce		json
//	@Param			body	body		med.AddMedWorkerRequest	true	"Запрос"
//	@Success		200		{object}	med.AddMedWorkerResponse
//	@Failure		500		{string}	string	"Internal error"
//	@Router			/med/worker/create [post]
func (c *Ctrl) AddMedWorker(ctx context.Context, req *pb.AddMedWorkerRequest) (*pb.AddMedWorkerResponse, error) {
	log.Println("Called AddMedWorker")
	return c.workerClient.AddMedWorker(ctx, req)
}

// GetPatientsByMedWorker godoc
//
//	@Summary		Get patients by med worker
//	@Description	Получить список пациентов, закрепленных за медицинским работником
//	@Tags			MedWorkers
//	@Accept			json
//	@Produce		json
//	@Param			body	body		med.GetPatientsByMedWorkerRequest	true	"Запрос"
//	@Success		200		{object}	med.GetPatientsByMedWorkerResponse
//	@Failure		500		{string}	string	"Internal error"
//	@Router			/med/worker/patients/{worker_id} [get]
func (c *Ctrl) GetPatientsByMedWorker(ctx context.Context, req *pb.GetPatientsByMedWorkerRequest) (*pb.GetPatientsByMedWorkerResponse, error) {
	log.Println("Called GetPatientsByMedWorker")
	return c.workerClient.GetPatientsByMedWorker(ctx, req)
}
