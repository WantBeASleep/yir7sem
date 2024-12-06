package card

import (
	"context"

	"med/internal/domain"
	pb "med/internal/generated/grpc/service"
	"med/internal/services/card"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CardHandler interface {
	CreateCard(ctx context.Context, in *pb.CreateCardIn) (*empty.Empty, error)
	GetCard(ctx context.Context, in *pb.GetCardIn) (*pb.GetCardOut, error)
	UpdateCard(ctx context.Context, in *pb.UpdateCardIn) (*pb.UpdateCardOut, error)
}

type handler struct {
	cardSrv card.Service
}

func New(
	cardSrv card.Service,
) CardHandler {
	return &handler{
		cardSrv: cardSrv,
	}
}

func (h *handler) CreateCard(ctx context.Context, in *pb.CreateCardIn) (*empty.Empty, error) {
	if err := h.cardSrv.CreateCard(ctx, domain.Card{
		DoctorID:  uuid.MustParse(in.Card.DoctorId),
		PatientID: uuid.MustParse(in.Card.PatientId),
		Diagnosis: in.Card.Diagnosis,
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &empty.Empty{}, nil
}

func (h *handler) GetCard(ctx context.Context, in *pb.GetCardIn) (*pb.GetCardOut, error) {
	card, err := h.cardSrv.GetCard(ctx, uuid.MustParse(in.DoctorId), uuid.MustParse(in.PatientId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.GetCardOut{Card: domainToPb(&card)}, nil
}

func (h *handler) UpdateCard(ctx context.Context, in *pb.UpdateCardIn) (*pb.UpdateCardOut, error) {
	card, err := h.cardSrv.UpdateCard(
		ctx,
		uuid.MustParse(in.Card.DoctorId),
		uuid.MustParse(in.Card.PatientId),
		card.UpdateCard{
			Diagnosis: in.Card.Diagnosis,
		},
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.UpdateCardOut{Card: domainToPb(&card)}, nil
}
