package card

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"yir/med/pkg/api/card"
)

type serverAPI struct {
	card.UnimplementedMedCardServer
}

func Register(server *grpc.Server) {
	card.RegisterMedCardServer(server, &serverAPI{})
}

func (s *serverAPI) GetCardList(ctx context.Context, request *card.GetCardListRequest) (*card.GetCardListResponse, error) {
	panic("not implemented")
}

func (s *serverAPI) CreateCard(ctx context.Context, request *card.CreateCardRequest) (*card.CreateCardResponse, error) {
	panic("not implemented")
}

func (s *serverAPI) GetCard(ctx context.Context, request *card.GetCardRequest) (*card.PatientCard, error) {
	panic("not implemented")
}

func (s *serverAPI) UpdateCard(ctx context.Context, request *card.UpdateCardRequest) (*empty.Empty, error) {
	panic("not implemented")
}

func (s *serverAPI) PatchCard(ctx context.Context, request *card.PatchCardRequest) (*card.PatientCard, error) {
	panic("not implemented")
}

func (s *serverAPI) DeleteCard(ctx context.Context, request *card.DeleteCardRequest) (*empty.Empty, error) {
	panic("not implemented")
}
