package medservice

import (
	"context"
	"yir/gateway/internal/custom"
	"yir/gateway/internal/entity/cardmodel"
	"yir/gateway/internal/pb/medpb"

	"go.uber.org/zap"
)

func (medsrv *MedService) GetCards(ctx context.Context, limit, offset uint64) (*cardmodel.PatientCardList, error) {
	req := &medpb.GetCardsRequest{
		Limit:  limit,
		Offset: offset,
	}
	data, err := medsrv.CardClient.GetCards(ctx, req)
	if err != nil {
		custom.Logger.Error(
			"failed get cards",
			zap.Error(err),
		)
		return nil, err
	}
	resp := &cardmodel.PatientCardList{
		Cards: make([]cardmodel.PatientCard, 0, 1),
		Count: data.Count,
	}
	for _, v := range data.Results {
		resp.Cards = append(resp.Cards, cardmodel.PatientCard{
			ID:              v.Id,
			AppointmentTime: v.AppointmentTime,
			HasNodules:      v.HasNodules,
			Diagnosis:       v.Diagnosis,
			MedWorkerID:     v.MedWorkerId,
			PatientID:       v.Patient.Id,
		})
	}
	return resp, nil
}

// Что передавать как аргумент??? Ниже тоже issue!
func (medsrv *MedService) PostCard(ctx context.Context, card *cardmodel.PatientInformation) error {
	req := &medpb.PostCardRequest{
		HasNodules: card.Card.HasNodules,
		Diagnosis:  card.Card.Diagnosis,
		Patient: &medpb.Patient{
			Id:            card.Patient.ID,
			FirstName:     card.Patient.FirstName,
			LastName:      card.Patient.LastName,
			FatherName:    card.Patient.FatherName,
			MedicalPolicy: card.Patient.MedicalPolicy,
			Email:         card.Patient.Email,
			IsActive:      card.Patient.IsActive,
		},
		MedworkerId: card.Card.MedWorkerID,
	}
	_, err := medsrv.CardClient.PostCard(ctx, req) // нахуя тут что то возвращается????????
	if err != nil {
		custom.Logger.Error(
			"failed to post card",
			zap.Error(err),
		)
		return err
	}
	return nil
}
func (medsrv *MedService) GetCardByID(ctx context.Context, id uint64) (*cardmodel.PatientCard, error) {
	req := &medpb.GetCardByIDRequest{
		Id: id,
	}
	data, err := medsrv.CardClient.GetCardByID(ctx, req)
	if err != nil {
		custom.Logger.Error(
			"failed to get card by id",
			zap.Error(err),
		)
		return nil, err
	}
	resp := &cardmodel.PatientCard{
		ID:              data.Postcard.Id,
		AppointmentTime: data.Postcard.AppointmentTime,
		Diagnosis:       data.Postcard.Diagnosis,
		HasNodules:      data.Postcard.HasNodules,
		MedWorkerID:     data.Postcard.MedWorkerId,
		PatientID:       data.Postcard.Patient.Id,
	}
	return resp, nil
}
func (medsrv *MedService) PutCard(ctx context.Context, card *cardmodel.PatientCard) (*cardmodel.PatientCard, error) {
	req := &medpb.PutCardRequest{
		Id:          card.ID,
		HasNodules:  card.HasNodules,
		Diagnosis:   card.Diagnosis,
		PatientId:   card.PatientID,
		MedworkerId: card.MedWorkerID,
	}
	data, err := medsrv.CardClient.PutCard(ctx, req)
	if err != nil {
		custom.Logger.Error(
			"failed to put card",
			zap.Error(err),
		)
		return nil, err
	}
	resp := &cardmodel.PatientCard{
		ID:              data.Postcard.Id,
		AppointmentTime: data.Postcard.AppointmentTime,
		HasNodules:      data.Postcard.HasNodules,
		Diagnosis:       data.Postcard.Diagnosis,
		MedWorkerID:     data.Postcard.MedWorkerId,
		PatientID:       data.Postcard.Patient.Id,
	}
	return resp, nil
}
func (medsrv *MedService) DeleteCard(ctx context.Context, id uint64) error {
	req := &medpb.DeleteCardRequest{
		Id: id,
	}
	_, err := medsrv.CardClient.DeleteCard(ctx, req) // опять чет возвращается :/
	if err != nil {
		custom.Logger.Error(
			"failed to delete card",
			zap.Error(err),
		)
		return err
	}
	return nil
}
