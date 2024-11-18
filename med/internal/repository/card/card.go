package repository

import (
	"context"
	"fmt"
	"yir/med/internal/entity"
	"yir/med/internal/repository/config"
	"yir/med/internal/repository/mapper"
	"yir/med/internal/repository/models"
	"yir/med/internal/repository/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CardRepo struct {
	db *gorm.DB
}

func NewRepository(cfg *config.DB) (*CardRepo, error) {
	fmt.Println("Connecting with DSN:", utils.GetDSN(cfg))
	db, err := gorm.Open(postgres.Open(utils.GetDSN(cfg)), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("create db gorm obj: %w", err)
	}
	db.AutoMigrate(
		&models.PatientInfo{},
		&models.PatientCardInfo{},
		&models.MedWorkerInfo{},
	)
	return &CardRepo{
		db: db,
	}, nil
}

// func (c *CardRepo) ListCards(ctx context.Context, limit, offset int) ([]*entity.PatientCard, int, error) {
// 	var (
// 		resp  []models.PatientCardInfo
// 		total int64
// 	)
// 	query := c.db.WithContext(ctx).
// 		Model(&models.PatientCardInfo{}).
// 		Count(&total)
// 	if err := query.Error; err != nil {
// 		return nil, 0, err
// 	}
// 	if err := query.Limit(limit).Offset(offset).Find(&resp).Error; err != nil {
// 		return nil, 0, err
// 	}

// 	entities := make([]*entity.PatientCard, len(resp))
// 	for i, card := range resp {
// 		medcard, err := mapper.PatientCardToEntity(&card)
// 		if err != nil {
// 			return nil, 0, err
// 		}
// 		entities[i] = medcard
// 	}

//		return entities, int(total), nil
//	}
func (c *CardRepo) ListCards(ctx context.Context, limit, offset int) ([]*entity.PatientCard, int, error) {
	var (
		cardsModel []models.PatientCardInfo
		total      int64
	)

	query := c.db.WithContext(ctx).
		Model(&models.PatientCardInfo{}).
		Count(&total)

	if err := query.Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(limit).Offset(offset).Find(&cardsModel).Error; err != nil {
		return nil, 0, err
	}

	var Cards []*entity.PatientCard
	for _, cardModel := range cardsModel {
		Cards = append(Cards, mapper.PatientCardToEntity(&cardModel))
	}

	return Cards, int(total), nil
}

func (c *CardRepo) CreateCard(ctx context.Context, Card *entity.PatientCard) (*entity.PatientCard, error) {
	tx := c.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to create patient: %w", tx.Error)
	}
	CardDB := mapper.PatientCardToModels(Card)
	if err := tx.
		Model(&models.PatientCardInfo{}).
		Create(&CardDB).
		Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create card info: %w", err)
	}
	resp := mapper.PatientCardToEntity(CardDB)

	return resp, tx.Commit().Error

}

func (c *CardRepo) CardByID(ctx context.Context, ID string) (*entity.PatientCard, error) {
	var cardModel models.PatientCardInfo
	if err := c.db.WithContext(ctx).Where("id = ?", ID).First(&cardModel).Error; err != nil {
		return nil, fmt.Errorf("failed to get patient card by id: %w", err)
	}

	cardEntity := mapper.PatientCardToEntity(&cardModel)

	return cardEntity, nil
}

func (c *CardRepo) UpdateCardInfo(ctx context.Context, Card *entity.PatientCard) error {
	query := c.db.WithContext(ctx).Begin()
	if query.Error != nil {
		return fmt.Errorf("failed to update card info: %w", query.Error)
	}
	CardDB := mapper.PatientCardToModels(Card)
	if err := query.Model(&models.PatientCardInfo{}).Where("id = ?", CardDB.ID).Updates(&CardDB).Error; err != nil {
		query.Rollback()
		return fmt.Errorf("failed to update patient card: %w", err)
	}
	return query.Commit().Error
}

func (c *CardRepo) PatchCardInfo(ctx context.Context, Card *entity.PatientCard) error {
	card := mapper.PatientCardToModels(Card)
	query := c.db.WithContext(ctx).
		Model(&models.PatientCardInfo{}).
		Where("id = ?", Card.ID).
		Updates(card)
	if err := query.Error; err != nil {
		return err
	}

	if query.RowsAffected == 0 {
		return entity.ErrNotFound
	}

	return nil
}

func (c *CardRepo) DeleteCardInfo(ctx context.Context, ID string) error {
	query := c.db.WithContext(ctx).
		Model(&models.PatientCardInfo{}).
		Where("id = ?", ID).
		Delete(&models.PatientCardInfo{})
	if err := query.Error; err != nil {
		return err
	}

	if query.RowsAffected == 0 {
		return entity.ErrNotFound
	}

	return nil
}
