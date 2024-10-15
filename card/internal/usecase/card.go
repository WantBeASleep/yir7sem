package usecase

import (
	"context"
	"fmt"
	"service/internal/entity"
	"service/internal/repository"

	"go.uber.org/zap"
)

type CardUseCase struct {
	CardRepo repository.Card
	logger   *zap.Logger
}

func NewCardUseCase(CardRepo repository.Card, logger *zap.Logger) *CardUseCase {
	return &CardUseCase{
		CardRepo: CardRepo,
		logger:   logger,
	}
}

func (c *CardUseCase) PostCard(ctx context.Context, Card *entity.PatientCard) error {
	c.logger.Debug("Starting PostCard usecase", zap.Any("CardInformation", Card))
	c.logger.Info("Adding new card")
	err := c.CardRepo.CreateCard(ctx, Card)
	if err != nil {
		c.logger.Error("Failed to add card to database", zap.Error(err))
		return fmt.Errorf("add card to DB: %w", err)
	}
	c.logger.Info("Successfully added new patient", zap.Any("CardInformation", Card))
	c.logger.Debug("PostCard usecase complete", zap.Any("CardInformation", Card))
	return nil
}

func (c *CardUseCase) GetCards(ctx context.Context, limit, offset int) (*entity.PatientCardList, error) {
	c.logger.Debug("Starting GetCardList usecase")
	c.logger.Info("Fetching cards list", zap.Int("limit", limit), zap.Int("offset", offset))
	cards, count, err := c.CardRepo.ListCards(ctx, limit, offset)
	if err != nil {
		c.logger.Error("Failed to fecth cards list", zap.Error(err))
		return nil, fmt.Errorf("get cards list: %w", err)
	}
	cardList := &entity.PatientCardList{
		Cards: make([]entity.PatientCard, len(cards)),
		Count: count,
	}
	for i, card := range cards {
		cardList.Cards[i] = *card
	}
	c.logger.Info("Successfully fetched cards list", zap.Int("number_of_cards", len(cards)))
	c.logger.Debug("GetCardList usecase complete")
	return cardList, nil
}

func (c *CardUseCase) GetCardByID(ctx context.Context, ID uint64) (*entity.PatientCard, error) {
	c.logger.Debug("Starting GetCardByID usecase", zap.Uint64("card_id", ID))
	c.logger.Info("fetching card by id", zap.Uint64("card_id", ID))
	cards, err := c.CardRepo.CardByID(ctx, ID)
	if err != nil {
		c.logger.Error("failed to fetch card by id", zap.Error(err), zap.Uint64("card_id", ID))
		return nil, fmt.Errorf("get card by id: %w", err)
	}
	c.logger.Info("Successfully fetched card information", zap.Uint64("card_id", ID))
	c.logger.Debug("GetPatientInfoByID usecase complete", zap.Uint64("card_id", ID))
	return cards, nil
}

func (c *CardUseCase) PutCard(ctx context.Context, Card *entity.PatientCard) error {
	c.logger.Debug("Starting PutCard usecase", zap.Any("Patient Card", Card))
	c.logger.Info("Updating card information", zap.String("card_id", fmt.Sprintf("%d", Card.ID)))
	err := c.CardRepo.UpdateCardInfo(ctx, Card)
	if err != nil {
		c.logger.Error("Failed to update card information", zap.Error(err))
		return fmt.Errorf("update card: %w", err)
	}

	c.logger.Info("Successfully updated card information", zap.String("card_id", fmt.Sprintf("%d", Card.ID)))
	c.logger.Debug("PutCard usecase complete", zap.Any("CardInformation", Card))
	return nil
}

func (c *CardUseCase) DeleteCard(ctx context.Context, ID uint64) error {
	c.logger.Debug("Starting DeleteCard usecase", zap.Uint64("card_id", ID))
	c.logger.Info("Deleting card by id", zap.Uint64("card_id", ID))

	err := c.CardRepo.DeleteCardInfo(ctx, int(ID))
	if err != nil {
		c.logger.Error("Failed to delete card", zap.Error(err), zap.Uint64("card_id", ID))
		return fmt.Errorf("delete card: %w", err)
	}

	c.logger.Info("Successfully deleted card", zap.Uint64("card_id", ID))
	c.logger.Debug("DeleteCard usecase complete", zap.Uint64("card_id", ID))
	return nil
}
