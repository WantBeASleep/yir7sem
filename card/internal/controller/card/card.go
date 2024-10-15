package controller

import (
	"context"
	"errors"
	pb "service/api/cards"
	"service/internal/entity"
	"service/internal/usecase"

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
	for _, card := range cardList.Cards {
		cardResponse := &pb.Card{
			Id:                 uint64(card.ID),
			AcceptanceDatetime: card.AppointmentTime,
			HasNodules:         card.HasNodules,
			Diagnosis:          card.Diagnosis,
			PatientId:          card.PatientID,
			MedworkerId:        card.MedWorkerID,
		}
		response.Results = append(response.Results, cardResponse)
	}
	return response, nil
}

func (s *Server) PostCard(ctx context.Context, request *pb.PostCardRequest) (*pb.PostCardResponse, error) {
	Card := &entity.PatientCard{
		HasNodules:  request.HasNodules,
		Diagnosis:   request.Diagnosis,
		PatientID:   request.PatientId,
		MedWorkerID: request.MedworkerId,
	}
	err := s.cardUseCase.PostCard(ctx, Card)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return nil, nil
}

func (s *Server) GetCardByID(ctx context.Context, request *pb.GetCardByIDRequest) (*pb.GetCardByIDResponse, error) {
	CardInfo, err := s.cardUseCase.GetCardByID(ctx, uint64(request.Id))
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
			Id:                 uint64(CardInfo.ID),
			AcceptanceDatetime: CardInfo.AppointmentTime,
			HasNodules:         CardInfo.HasNodules,
			Diagnosis:          CardInfo.Diagnosis,
			PatientId:          CardInfo.PatientID,
			MedworkerId:        CardInfo.MedWorkerID,
		},
	}
	return response, nil

}

func (s *Server) PutCard(ctx context.Context, request *pb.PutCardRequest) (*pb.PutCardResponse, error) {
	Card := &entity.PatientCard{
		ID:          int(request.Id),
		HasNodules:  request.HasNodules,
		Diagnosis:   request.Diagnosis,
		PatientID:   request.PatientId,
		MedWorkerID: request.MedworkerId,
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
