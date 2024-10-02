package usecases

import (
	"context"
	"fmt"
	"service/internal/entity"
	"service/internal/usecases/repositories"

	"go.uber.org/zap"
)

type MedWorkerUseCase struct {
	MedWorkerRepo  repositories.MedWorker
	PatientService repositories.PatientService
	logger         *zap.Logger
}

func NewMedWorkerUseCase(MedWorkerRepo repositories.MedWorker, patientService repositories.PatientService, logger *zap.Logger) *MedWorkerUseCase {
	return &MedWorkerUseCase{MedWorkerRepo: MedWorkerRepo,
		PatientService: patientService,
		logger:         logger,
	}
}

func (m *MedWorkerUseCase) GetMedWorkers(ctx context.Context, limit, offset int) (*entity.MedicalWorkerList, error) {
	m.logger.Info("Fetching medical workers", zap.Int("limit", limit), zap.Int("offset", offset))
	// Получаем список медработников из репозитория
	workers, count, err := m.MedWorkerRepo.ListMedicalWorkers(ctx, limit, offset)
	if err != nil {
		m.logger.Error("Failed to fetch medical workers", zap.Error(err))
		return nil, fmt.Errorf("failed to fetch medical workers: %w", err)
	}
	medicalWorkerList := &entity.MedicalWorkerList{
		Workers: make([]entity.MedicalWorker, len(workers)),
		Count:   count,
	}

	// Копируем данные из модели репозитория в сущности(im not sure if needed)
	for i, worker := range workers {
		medicalWorkerList.Workers[i] = *worker
	}

	return medicalWorkerList, nil
}

func (m *MedWorkerUseCase) GetMedWorkerByID(ctx context.Context, ID int) (*entity.MedicalWorker, error) {
	m.logger.Info("Fetching medical worker by ID", zap.Int("ID", ID))

	// Используем репозиторий для получения медработника по ID
	worker, err := m.MedWorkerRepo.GetMedicalWorkerByID(ctx, ID)
	if err != nil {
		if err == entity.ErrNotFound {
			m.logger.Warn("Медицинский работник не найден", zap.Int("ID", ID))
			return nil, fmt.Errorf("Медицинский работник с %d не был найден: %w", ID, err)
		}
		m.logger.Error("Failed to fetch medical worker by ID", zap.Error(err))
		return nil, fmt.Errorf("failed to fetch medical worker by ID %d: %w", ID, err)
	}

	return worker, nil
}

func (m *MedWorkerUseCase) UpdateMedWorker(ctx context.Context, ID int, updateData *entity.MedicalWorkerUpdateRequest) (*entity.MedicalWorker, error) {
	m.logger.Info("Updating medical worker", zap.Int("ID", ID))
	// Сначала находим медработника по ID
	worker, err := m.MedWorkerRepo.GetMedicalWorkerByID(ctx, ID)
	if err != nil {
		if err == entity.ErrNotFound {
			m.logger.Warn("Работник не был найден для изменения данных", zap.Int("ID", ID))
			return nil, fmt.Errorf("Медицинский работник с ID %d не был найден: %w", ID, err)
		}
		m.logger.Error("Failed to fetch medical worker for update", zap.Error(err))
		return nil, fmt.Errorf("failed to fetch medical worker by ID %d: %w", ID, err)
	}

	// Обновляем поля медработника на основе данных из updateData
	worker.FirstName = updateData.FirstName
	worker.LastName = updateData.LastName
	worker.MiddleName = updateData.MiddleName
	worker.MedOrganization = updateData.MedOrganization
	worker.Job = updateData.Job
	worker.IsRemoteWorker = updateData.IsRemoteWorker
	worker.ExpertDetails = updateData.ExpertDetails
	// Сохраняем обновленные данные в базе данных
	if err := m.MedWorkerRepo.UpdateMedicalWorker(ctx, worker); err != nil {
		m.logger.Error("Ошибка при обновлении данных медицинского работника", zap.Error(err))
		return nil, fmt.Errorf("Ошибка при обновлении данных медицинского работника с ID %d: %w", ID, err)
	}

	return worker, nil
}
func (m *MedWorkerUseCase) AddMedWorker(ctx context.Context, createData *entity.AddMedicalWorkerRequest) (*entity.MedicalWorker, error) {
	m.logger.Info("Adding new medical worker", zap.String("FirstName", createData.FirstName))

	// Создаем объект MedicalWorker на основе данных из запроса
	medworker := &entity.MedicalWorker{
		FirstName:       createData.FirstName,
		MiddleName:      createData.MiddleName,
		LastName:        createData.LastName,
		MedOrganization: createData.MedOrganization,
		Job:             createData.Job,
		IsRemoteWorker:  createData.IsRemoteWorker,
		ExpertDetails:   createData.ExpertDetails,
	}

	// Добавляем медработника через репозиторий и получаем ID
	id, err := m.MedWorkerRepo.AddMedicalWorker(ctx, medworker)
	if err != nil {
		m.logger.Error("Failed to add new medical worker", zap.Error(err))
		return nil, fmt.Errorf("failed to add medical worker: %w", err)
	}

	// Присваиваем ID новому медработнику
	medworker.ID = id
	m.logger.Info("Successfully added new medical worker", zap.Int("ID", id))

	// Возвращаем объект MedicalWorker и ошибку
	return medworker, nil
}

func (m *MedWorkerUseCase) GetPatientsByMedWorker(ctx context.Context, medWorkerId uint64) (*entity.MedicalWorkerWithPatients, error) {
	m.logger.Info("Fetching patients for medical worker", zap.Any("med_worker_id", medWorkerId))

	// Получение информации о медработнике из базы данных
	medWorker, err := m.MedWorkerRepo.GetMedicalWorkerByID(ctx, int(medWorkerId))
	if err != nil {
		return nil, err
	}

	// Вызов gRPC метода из PatientService для получения списка пациентов
	patientPtrs, err := m.PatientService.GetPatientsByMedWorker(ctx, medWorkerId)
	if err != nil {
		return nil, err
	}

	// Преобразуем срез указателей в срез значений
	patients := make([]entity.PatientCardDTO, len(patientPtrs))
	for i, patientPtr := range patientPtrs {
		patients[i] = *patientPtr
	}

	// Возвращаем данные с врачом и его пациентами
	return &entity.MedicalWorkerWithPatients{
		MedWorker: *medWorker,
		Patients:  patients,
	}, nil
}
