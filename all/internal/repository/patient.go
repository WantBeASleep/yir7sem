package repository

import (
	"context"
	"fmt"
	"service/all/internal/entity"
	"service/all/internal/repository/config"
	"service/all/internal/repository/mapper"
	"service/all/internal/repository/models"
	"service/all/internal/repository/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PatientRepo struct {
	db *gorm.DB
}

func NewRepository(cfg *config.DB) (*PatientRepo, error) {
	db, err := gorm.Open(postgres.Open(utils.GetDSN(cfg)), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("create db gorm obj: %w", err)
	}

	db.AutoMigrate(
		&models.PatientInfo{},
		&models.PatientCardInfo{},
	)

	return &PatientRepo{
		db: db,
	}, nil
}

func (p *PatientRepo) CreatePatient(ctx context.Context, PatientInfo *entity.PatientInformation) error {
	tx := p.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to update patient: %w", tx.Error)
	}

	PatientDB := mapper.PatientToModels(PatientInfo.Patient)
	if err := tx.
		Model(&models.PatientInfo{}).
		Create(&PatientDB).
		Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create patient info: %w", err)
	}

	CardDB := mapper.PatientCardToModels(PatientInfo.Card)
	if err := tx.
		Model(&models.PatientCardInfo{}).
		Create(&CardDB).
		Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create patient card: %w", err)
	}

	return tx.Commit().Error

}

func (p *PatientRepo) UpdatePatient(ctx context.Context, PatientInfo *entity.PatientInformation) error {
	tx := p.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to update patient: %w", tx.Error)
	}

	PatientDB := mapper.PatientToModels(PatientInfo.Patient)
	if err := tx.Model(&models.PatientInfo{}).Where("id = ?", PatientDB.ID).Updates(&PatientDB).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update patient info: %w", err)
	}

	CardDB := mapper.PatientCardToModels(PatientInfo.Card)
	if err := tx.Model(&models.PatientCardInfo{}).Where("id = ?", CardDB.ID).Updates(&CardDB).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update patient card: %w", err)
	}

	return tx.Commit().Error
}

func (p *PatientRepo) GetListPatient(ctx context.Context) ([]*entity.Patient, error) {
	var resp []*models.PatientInfo
	query := p.db.WithContext(ctx).Model(&models.PatientInfo{})

	if err := query.Find(&resp).Error; err != nil {
		return nil, fmt.Errorf("failed to get list patient: %w", err)
	}

	patients := make([]*entity.Patient, len(resp))
	for i := range resp {
		patients[i] = mapper.PatientToEntity(resp[i])
	}

	return patients, nil
}

func (p *PatientRepo) GetPatientInfoByID(ctx context.Context, ID uint64) (*entity.PatientInformation, error) {
	var respPatient *models.PatientInfo
	var respCard *models.PatientCardInfo

	tx := p.db.WithContext(ctx)
	if err := tx.First(&respPatient, ID).Error; err != nil {
		return nil, fmt.Errorf("failed to get patient: %w", err)
	}

	if err := tx.Model(&models.PatientCardInfo{}).Where("patient_id = ?", ID).First(&respCard).Error; err != nil {
		return nil, fmt.Errorf("failed to get patient card: %w", err)
	}
	cardEntity, err := mapper.PatientCardToEntity(respCard)
	if err != nil {
		return nil, fmt.Errorf("failed to map patient card: %w", err)
	}
	PatientInfo := &entity.PatientInformation{
		Patient: mapper.PatientToEntity(respPatient),
		Card:    cardEntity,
	}

	return PatientInfo, nil
}
