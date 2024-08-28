package repositories

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"yir/medworkers/internal/config"
	"yir/medworkers/internal/entity"
	"yir/medworkers/internal/repositories/mappers"
	"yir/medworkers/internal/repositories/models"
	"yir/medworkers/internal/repositories/utils"
)

type MedicalWorkerRepo struct {
	db *gorm.DB
}

func NewRepository(cfg *config.DB) (*MedicalWorkerRepo, error) {
	fmt.Println("Connecting with DSN:", utils.GetDSN(cfg))
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: utils.GetDSN(cfg), //Непонятно, нужен ли указатель
	}), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("create database gorm object: %w", err)
	}
	db.AutoMigrate(&models.MedWorkerInfo{})

	return &MedicalWorkerRepo{
		db: db,
	}, nil
}

func (r *MedicalWorkerRepo) GetMedicalWorkerByID(ctx context.Context, ID int) (*entity.MedicalWorker, error) {
	var worker models.MedWorkerInfo

	query := r.db.WithContext(ctx).
		Model(&models.MedWorkerInfo{}).
		Where("id = ?", ID)
	if err := query.Take(&worker).Error; err != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {

			return nil, entity.ErrNotFound
		}
		return nil, err
	}
	medworker, err := mappers.ToMedWorkerEntity(&worker)
	if err != nil {
		return nil, err
	}
	return medworker, nil
}

/*
	func (r *MedicalWorkerRepo) CreateMedicalWorker(ctx context.Context, medworker *entity.MedicalWorker) (int, error) {
		worker, err := mappers.ToMedWorkerModel(medworker)
		if err != nil {
			return 0, err
		}

		if err := r.db.WithContext(ctx).
			Model(&models.MedWorkerInfo{}).
			Create(&worker).
			Error; err != nil {
			return 0, err
		}
		return int(medworker.ID), nil
	}
*/
func (r *MedicalWorkerRepo) UpdateMedicalWorker(ctx context.Context, medworker *entity.MedicalWorker) error {
	worker, err := mappers.ToMedWorkerModel(medworker)
	if err != nil {
		return err
	}

	query := r.db.WithContext(ctx).
		Model(&models.MedWorkerInfo{}).
		Where("id = ?", medworker.ID).
		Updates(worker)
	if err := query.Error; err != nil {
		return err
	}

	if query.RowsAffected == 0 {
		return entity.ErrNotFound
	}

	return nil
}

func (r *MedicalWorkerRepo) PatchMedicalWorker(ctx context.Context, medworker *entity.MedicalWorker) error {
	worker, err := mappers.ToMedWorkerModel(medworker)
	if err != nil {
		return err
	}

	query := r.db.WithContext(ctx).
		Model(&models.MedWorkerInfo{}).
		Where("id = ?", medworker.ID).
		Updates(worker)
	if err := query.Error; err != nil {
		return err
	}

	if query.RowsAffected == 0 {
		return entity.ErrNotFound
	}

	return nil
}

func (r *MedicalWorkerRepo) ListMedicalWorkers(ctx context.Context, limit, offset int) ([]*entity.MedicalWorker, int, error) {
	var workers []models.MedWorkerInfo
	var total int64

	query := r.db.WithContext(ctx).
		Model(&models.MedWorkerInfo{}).
		Count(&total)
	if err := query.Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(limit).Offset(offset).Find(&workers).Error; err != nil {
		return nil, 0, err
	}

	entities := make([]*entity.MedicalWorker, len(workers))
	for i, worker := range workers {
		medworker, err := mappers.ToMedWorkerEntity(&worker)
		if err != nil {
			return nil, 0, err
		}
		entities[i] = medworker
	}

	return entities, int(total), nil
}

func (r *MedicalWorkerRepo) DeleteMedicalWorker(ctx context.Context, ID int) error {
	query := r.db.WithContext(ctx).
		Model(&models.MedWorkerInfo{}).
		Where("id = ?", ID).
		Delete(&models.MedWorkerInfo{})
	if err := query.Error; err != nil {
		return err
	}

	if query.RowsAffected == 0 {
		return entity.ErrNotFound
	}

	return nil
}
