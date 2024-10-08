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

type PatientRepo struct {
	db *gorm.DB
}

func NewRepository(cfg *config.DB) (*PatientRepo, error) {
	db, err := gorm.Open(postgres.Open(utils.GetDSN(cfg)), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("create db gorm obj: %w", err)
	}

	db.AutoMigrate(
		&models.Patient{},
		&models.PatientCard{},
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

	PatientDB := mapper.PatientToModel(PatientInfo.Patient)
	if err := tx.
		Model(&models.Patient{}).
		Create(&PatientDB).
		Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create patient info: %w", err)
	}

	CardDB := mapper.PatientCardToModel(PatientInfo.Card)
	if err := tx.
		Model(&models.PatientCard{}).
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

	PatientDB := mapper.PatientToModel(PatientInfo.Patient)
	if err := tx.Model(&models.Patient{}).Where("id = ?", PatientDB.ID).Updates(&PatientDB).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update patient info: %w", err)
	}

	CardDB := mapper.PatientCardToModel(PatientInfo.Card)
	if err := tx.Model(&models.PatientCard{}).Where("id = ?", CardDB.ID).Updates(&CardDB).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update patient card: %w", err)
	}

	return tx.Commit().Error
}

func (p *PatientRepo) GetListPatient(ctx context.Context) ([]*entity.Patient, error) {
	var resp []*models.Patient
	query := p.db.WithContext(ctx).Model(&models.Patient{})

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
	var respPatient *models.Patient
	var respCard *models.PatientCard

	tx := p.db.WithContext(ctx)
	if err := tx.First(&respPatient, ID).Error; err != nil {
		return nil, fmt.Errorf("failed to get patient: %w", err)
	}

	if err := tx.Model(&models.PatientCard{}).Where("patient_id = ?", ID).First(&respCard).Error; err != nil {
		return nil, fmt.Errorf("failed to get patient card: %w", err)
	}

	PatientInfo := &entity.PatientInformation{
		Patient: mapper.PatientToEntity(respPatient),
		Card:    mapper.PatientCardToEntity(respCard),
	}

	return PatientInfo, nil
}
