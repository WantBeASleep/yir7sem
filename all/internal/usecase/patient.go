package usecase

import (
	"context"
	"fmt"
	"service/all/internal/entity"
	"service/all/internal/repository"

	"go.uber.org/zap"
)

type PatientUseCase struct {
	PatientRepo repository.PatientRepo

	logger *zap.Logger
}

func NewPatientUseCase(PatientRepo repository.PatientRepo, logger *zap.Logger) *PatientUseCase {
	return &PatientUseCase{
		PatientRepo: PatientRepo,
		logger:      logger,
	}
}

func (p *PatientUseCase) AddPatient(ctx context.Context, PatientInfo *entity.PatientInformation) error {
	p.logger.Debug("Starting AddPatient usecase", zap.Any("PatientInformation", PatientInfo))

	p.logger.Info("Adding new patient")
	err := p.PatientRepo.CreatePatient(ctx, PatientInfo)
	if err != nil {
		p.logger.Error("Failed to add patient to database", zap.Error(err))
		return fmt.Errorf("add patient in DB: %w", err)
	}

	p.logger.Info("Successfully added new patient", zap.Any("PatientInformation", PatientInfo))
	p.logger.Debug("AddPatient usecase complete", zap.Any("PatientInformation", PatientInfo))
	return nil
}

func (p *PatientUseCase) UpdatePatient(ctx context.Context, PatientInfo *entity.PatientInformation) error {
	p.logger.Debug("Starting UpdatePatient usecase", zap.Any("PatientInformation", PatientInfo))

	p.logger.Info("Updating patient information", zap.String("patient_id", fmt.Sprintf("%d", PatientInfo.Patient.ID.String())))
	err := p.PatientRepo.UpdatePatient(ctx, PatientInfo)
	if err != nil {
		p.logger.Error("Failed to update patient information", zap.Error(err))
		return fmt.Errorf("update patient: %w", err)
	}

	p.logger.Info("Successfully updated patient information", zap.String("patient_id", fmt.Sprintf("%d", PatientInfo.Patient.ID.String())))
	p.logger.Debug("UpdatePatient usecase complete", zap.Any("PatientInformation", PatientInfo))
	return nil
}

func (p *PatientUseCase) GetPatientList(ctx context.Context) ([]*entity.Patient, error) {
	p.logger.Debug("Starting GetPatientList usecase")

	p.logger.Info("Fetching patient list")
	patients, err := p.PatientRepo.GetListPatient(ctx)
	if err != nil {
		p.logger.Error("Failed to fetch patient list", zap.Error(err))
		return nil, fmt.Errorf("get patient list: %w", err)
	}

	p.logger.Info("Successfully fetched patient list", zap.Int("number_of_patients", len(patients)))
	p.logger.Debug("GetPatientList usecase complete")
	return patients, nil
}

func (p *PatientUseCase) GetPatientInfoByID(ctx context.Context, ID string) (*entity.PatientInformation, error) {
	p.logger.Debug("Starting GetPatientInfoByID usecase", zap.Any("patient_id", ID))

	p.logger.Info("Fetching patient information", zap.Any("patient_id", ID))
	patient, err := p.PatientRepo.GetPatientInfoByID(ctx, ID)
	if err != nil {
		p.logger.Error("Failed to fetch patient information", zap.Error(err), zap.Any("patient_id", ID))
		return nil, fmt.Errorf("get patient info by id: %w", err)
	}

	p.logger.Info("Successfully fetched patient information", zap.Any("patient_id", ID))
	p.logger.Debug("GetPatientInfoByID usecase complete", zap.Any("patient_id", ID))
	return patient, nil
}
