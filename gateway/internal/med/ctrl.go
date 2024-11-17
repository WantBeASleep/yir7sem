//TODO: Сделать магамеду

package med

import (
	"context"
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

func (c *Ctrl) GetCards(ctx context.Context, req *pb.GetCardsRequest) (*pb.GetCardsResponse, error) {
	return c.cardClient.GetCards(ctx, req)
}

func (c *Ctrl) PostCard(ctx context.Context, req *pb.PostCardRequest) (*pb.PostCardResponse, error) {
	return c.cardClient.PostCard(ctx, req)
}

func (c *Ctrl) GetCardByID(ctx context.Context, req *pb.GetCardByIDRequest) (*pb.GetCardByIDResponse, error) {
	return c.cardClient.GetCardByID(ctx, req)
}

func (c *Ctrl) PutCard(ctx context.Context, req *pb.PutCardRequest) (*pb.PutCardResponse, error) {
	return c.cardClient.PutCard(ctx, req)
}

func (c *Ctrl) PatchCard(ctx context.Context, req *pb.PatchCardRequest) (*pb.PatchCardResponse, error) {
	return c.cardClient.PatchCard(ctx, req)
}

func (c *Ctrl) DeleteCard(ctx context.Context, req *pb.DeleteCardRequest) (*pb.DeleteCardResponse, error) {
	return c.cardClient.DeleteCard(ctx, req)
}

func (c *Ctrl) AddPatient(ctx context.Context, req *pb.CreatePatientRequest) (*empty.Empty, error) {
	return c.patientClient.AddPatient(ctx, req)
}
func (c *Ctrl) GetPatientList(ctx context.Context, req *empty.Empty) (*pb.PatientListResponse, error) {
	return c.patientClient.GetPatientList(ctx, req)
}
func (c *Ctrl) GetPatientInfoByID(ctx context.Context, req *pb.PatientInfoRequest) (*pb.PatientInfoResponse, error) {
	return c.patientClient.GetPatientInfoByID(ctx, req)
}
func (c *Ctrl) UpdatePatient(ctx context.Context, req *pb.PatientUpdateRequest) (*empty.Empty, error) {
	return c.patientClient.UpdatePatient(ctx, req)
}

func (c *Ctrl) GetMedWorkers(ctx context.Context, req *pb.GetMedworkerRequest) (*pb.GetMedworkerListResponse, error) {
	return c.workerClient.GetMedWorkers(ctx, req)
}
func (c *Ctrl) UpdateMedWorker(ctx context.Context, req *pb.UpdateMedWorkerRequest) (*pb.UpdateMedWorkerResponse, error) {
	return c.workerClient.UpdateMedWorker(ctx, req)
}
func (c *Ctrl) GetMedWorkerByID(ctx context.Context, req *pb.GetMedMedWorkerByIDRequest) (*pb.GetMedWorkerByIDResponse, error) {
	return c.workerClient.GetMedWorkerByID(ctx, req)
}
func (c *Ctrl) PatchMedWorker(ctx context.Context, req *pb.PatchMedWorkerRequest) (*pb.UpdateMedWorkerResponse, error) {
	return c.workerClient.PatchMedWorker(ctx, req)
}
func (c *Ctrl) AddMedWorker(ctx context.Context, req *pb.AddMedWorkerRequest) (*pb.AddMedWorkerResponse, error) {
	return c.workerClient.AddMedWorker(ctx, req)
}
func (c *Ctrl) GetPatientsByMedWorker(ctx context.Context, req *pb.GetPatientsByMedWorkerRequest) (*pb.GetPatientsByMedWorkerResponse, error) {
	return c.workerClient.GetPatientsByMedWorker(ctx, req)
}
