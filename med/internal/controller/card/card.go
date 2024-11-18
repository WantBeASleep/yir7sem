package controller

import (
	"context"
	"errors"
	pb "yir/med/api"
	"yir/med/internal/entity"
	"yir/med/internal/usecase"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedMedCardServer
	cardUseCase usecase.Card
	logger      *zap.Logger
}

func NewServer(cardUseCase usecase.Card, logger *zap.Logger) *Server {
	return &Server{
		cardUseCase: cardUseCase,
		logger:      logger,
	}
}

func (s *Server) GetCards(ctx context.Context, request *pb.GetCardsRequest) (*pb.GetCardsResponse, error) {
	s.logger.Info("[Request] New request", zap.Any("data", request))
	limit := int(request.GetLimit())
	offset := int(request.GetOffset())

	cardList, err := s.cardUseCase.GetCards(ctx, limit, offset)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "Cards were not found")
		}
		s.logger.Error("Failed to fetch cards", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to fetch cards")
	}

	response := &pb.GetCardsResponse{
		Count:   uint64(cardList.Count),
		Results: []*pb.Card{},
	}
	for _, cardInfo := range cardList.Cards {
		cardResponse := &pb.Card{
			Id:              cardInfo.ID.String(),
			AppointmentTime: cardInfo.AppointmentTime,
			HasNodules:      cardInfo.HasNodules,
			Diagnosis:       cardInfo.Diagnosis,
			PatientId:       cardInfo.PatientID.String(),
			MedWorkerId:     cardInfo.MedWorkerID.String(),
		}
		response.Results = append(response.Results, cardResponse)
	}

	return response, nil
}

func (s *Server) PostCard(ctx context.Context, request *pb.PostCardRequest) (*pb.PostCardResponse, error) {
	uuid1, _ := uuid.Parse(request.PatientId)
	uuid2, _ := uuid.Parse(request.MedworkerId)
	card := &entity.PatientCard{
		HasNodules:  request.HasNodules,
		Diagnosis:   request.Diagnosis,
		PatientID:   uuid1,
		MedWorkerID: uuid2,
	}
	err := s.cardUseCase.PostCard(ctx, card)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return nil, nil
}

func (s *Server) GetCardByID(ctx context.Context, request *pb.GetCardByIDRequest) (*pb.GetCardByIDResponse, error) {
	CardInfo, err := s.cardUseCase.GetCardByID(ctx, request.Id)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrNotFound):
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	response := &pb.GetCardByIDResponse{
		Postcard: &pb.Card{
			Id:              CardInfo.ID.String(),
			AppointmentTime: CardInfo.AppointmentTime,
			HasNodules:      CardInfo.HasNodules,
			Diagnosis:       CardInfo.Diagnosis,
			PatientId:       CardInfo.PatientID.String(),
			MedWorkerId:     CardInfo.MedWorkerID.String(),
		},
	}
	return response, nil

}

func (s *Server) PutCard(ctx context.Context, request *pb.PutCardRequest) (*pb.PutCardResponse, error) {
	uuid1, _ := uuid.Parse(request.Id)
	uuid2, _ := uuid.Parse(request.PatientId)
	uuid3, _ := uuid.Parse(request.MedworkerId)
	Card := &entity.PatientCard{
		ID:          uuid1,
		HasNodules:  request.HasNodules,
		Diagnosis:   request.Diagnosis,
		PatientID:   uuid2,
		MedWorkerID: uuid3,
	}
	err := s.cardUseCase.PutCard(ctx, Card)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrNotFound):
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return nil, nil
}

func (s *Server) DeleteCard(ctx context.Context, request *pb.DeleteCardRequest) (*pb.DeleteCardResponse, error) {
	s.logger.Info("[Request] Delete card", zap.Any("request", request))

	err := s.cardUseCase.DeleteCard(ctx, request.Id)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "Card not found")
		}
		s.logger.Error("Failed to delete card", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to delete card")
	}

	return &pb.DeleteCardResponse{}, nil
}
