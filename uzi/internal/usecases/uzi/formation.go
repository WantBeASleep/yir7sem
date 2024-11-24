package uzi

import (
	"context"
	"fmt"
	"yir/pkg/mappers"
	"yir/uzi/internal/entity"
	"yir/uzi/internal/usecases/dto"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (u *UziUseCase) CreateFormationWithSegments(ctx context.Context, req *dto.FormationWithSegments) (uuid.UUID, uuid.UUIDs, error) {
	formationID, err := u.CreateDTOFormation(ctx, req.Formation)
	if err != nil {
		return uuid.Nil, nil, fmt.Errorf("insert dto formation: %w", err)
	}

	for i := range req.Segments {
		req.Segments[i].FormationID = formationID
	}

	segmentsIDS, err := u.CreateDTOSegments(ctx, req.Segments)
	if err != nil {
		return uuid.Nil, nil, fmt.Errorf("insert dto segments: %w", err)
	}

	return formationID, segmentsIDS, nil
}

func (u *UziUseCase) InsertFormationsAndSegemetsSeparately(ctx context.Context, formations []dto.Formation, segments []dto.Segment) error {
	for _, formation := range formations {
		err := u.insertDTOFormation(ctx, &formation)
		if err != nil {
			return fmt.Errorf("insert formation: %w", err)
		}
	}

	err := u.InsertDTOSegments(ctx, segments)
	if err != nil {
		return fmt.Errorf("insert segments: %w", err)
	}

	return nil
}

// TODO: Не понятно нужно ли это
func (u *UziUseCase) CreateFormationsWithSegments(ctx context.Context, req []dto.FormationWithSegments) ([]dto.FormationWithSegmentsIDs, error) {
	resp := make([]dto.FormationWithSegmentsIDs, 0, len(req))
	for _, formation := range req {
		formationID, segmentsIDS, err := u.CreateFormationWithSegments(ctx, &formation)
		if err != nil {
			return nil, fmt.Errorf("create formation with segments: %w", err)
		}

		resp = append(resp, dto.FormationWithSegmentsIDs{
			FormationID: formationID,
			SegmentsIDs: segmentsIDS,
		})
	}

	return resp, nil
}

func (u *UziUseCase) GetFormationWithSegments(ctx context.Context, id uuid.UUID) (*dto.FormationWithSegments, error) {
	u.logger.Debug("[Request] Get formation by ID", zap.String("formation id", id.String()))
	formation, err := u.uziRepo.GetFormationByID(ctx, id)
	if err != nil {
		u.logger.Error("Get formation by ID", zap.Error(err))
		return nil, fmt.Errorf("get formation by ID: %w", err)
	}
	u.logger.Debug("[Response] Got formation by ID", zap.String("formation id", id.String()))

	dtoFormation, err := u.GetDTOFormationFromEntity(ctx, formation)
	if err != nil {
		return nil, fmt.Errorf("get dto formation: %w", err)
	}

	u.logger.Debug("[Request] Get formation segments", zap.String("formation id", id.String()))
	segments, err := u.uziRepo.GetSegmentsByFormationID(ctx, id)
	if err != nil {
		u.logger.Error("Get formation segments", zap.Error(err))
		return nil, fmt.Errorf("get formation segments: %w", err)
	}
	u.logger.Debug("[Response] Get formation segments", zap.Int("count segments", len(segments)))

	dtoSegments, err := u.GetDTOSegmentsFromEntity(ctx, segments)
	if err != nil {
		return nil, fmt.Errorf("get dto segments: %w", err)
	}

	return &dto.FormationWithSegments{
		Formation: dtoFormation,
		Segments:  dtoSegments,
	}, nil
}

func (u *UziUseCase) UpdateFormation(ctx context.Context, id uuid.UUID, req *dto.Formation) (*dto.Formation, error) {
	u.logger.Debug("[Request] Create tirads")
	tiradsID, err := u.uziRepo.CreateTirads(ctx, req.Tirads)
	if err != nil {
		u.logger.Error("Create tirads", zap.Error(err))
		return nil, fmt.Errorf("create tirads: %w", err)
	}
	u.logger.Debug("[Response] Created tirads", zap.Int("id", tiradsID))

	formation := mappers.MustTransformObj[dto.Formation, entity.Formation](req)
	formation.TiradsID = tiradsID

	u.logger.Debug("[Request] Update Formation", zap.Any("id", id), zap.Any("data", req))
	updateFormation, err := u.uziRepo.UpdateFormation(ctx, id, formation)
	if err != nil {
		u.logger.Error("Update Formation", zap.Error(err))
		return nil, fmt.Errorf("update formation: %w", err)
	}
	u.logger.Debug("[Response] Updated Formation")

	dtoUpdateFormation, err := u.GetDTOFormationFromEntity(ctx, updateFormation)
	if err != nil {
		return nil, fmt.Errorf("get updated formation: %w", err)
	}

	return dtoUpdateFormation, nil
}

func (u *UziUseCase) CreateDTOFormation(ctx context.Context, formation *dto.Formation) (uuid.UUID, error) {
	u.logger.Debug("[Request] Create tirads")
	tiradsID, err := u.uziRepo.CreateTirads(ctx, formation.Tirads)
	if err != nil {
		u.logger.Error("Create tirads", zap.Error(err))
		return uuid.Nil, fmt.Errorf("create tirads: %w", err)
	}
	u.logger.Debug("[Response] Created tirads")

	entFormation := mappers.MustTransformObj[dto.Formation, entity.Formation](formation)
	entFormation.TiradsID = tiradsID

	u.logger.Debug("[Request] Create dto formation")
	formationID := uuid.New()
	entFormation.Id = formationID
	if err := u.uziRepo.InsertFormation(ctx, entFormation); err != nil {
		u.logger.Error("Insert formations", zap.Error(err))
		return uuid.Nil, fmt.Errorf("insert formations: %w", err)
	}
	u.logger.Debug("[Response] Created dto formation")

	return formationID, nil
}

// TODO: Изначально было очень плохим решением выделять это dto с tirads, потому что это плодит невероятное количество неудобств
// единственное где это нужно - на стороне controller'а, правильно было бы вынести все вложения/выделения туда,
// а в entity оставить неприкосаемые сущности formation и entity
func (u *UziUseCase) insertDTOFormation(ctx context.Context, formation *dto.Formation) error {
	u.logger.Debug("[Request] Create tirads")
	tiradsID, err := u.uziRepo.CreateTirads(ctx, formation.Tirads)
	if err != nil {
		u.logger.Error("Create tirads", zap.Error(err))
		return fmt.Errorf("create tirads: %w", err)
	}
	u.logger.Debug("[Response] Created tirads")

	entFormation := mappers.MustTransformObj[dto.Formation, entity.Formation](formation)
	entFormation.TiradsID = tiradsID

	u.logger.Debug("[Request] Create dto formation")
	if err := u.uziRepo.InsertFormation(ctx, entFormation); err != nil {
		u.logger.Error("Insert formations", zap.Error(err))
		return fmt.Errorf("insert formations: %w", err)
	}
	u.logger.Debug("[Response] Created dto formation")

	return nil
}

func (u *UziUseCase) GetDTOFormationFromEntity(ctx context.Context, formation *entity.Formation) (*dto.Formation, error) {
	u.logger.Debug("[Request] Get formation tirads")
	tirads, err := u.uziRepo.GetTiradsByID(ctx, formation.TiradsID)
	if err != nil {
		u.logger.Error("Get formation tirads")
		return nil, fmt.Errorf("get tirads [id %q]: %w", formation.TiradsID, err)
	}
	u.logger.Debug("[Response] Get formation tirads")

	dtoFormation := mappers.MustTransformObj[entity.Formation, dto.Formation](formation)
	dtoFormation.Tirads = tirads

	return dtoFormation, nil
}

func (u *UziUseCase) GetDTOFormationsFromEntity(ctx context.Context, formations []entity.Formation) ([]dto.Formation, error) {
	u.logger.Debug("[Request] Get formation tirads", zap.Int("count formations", len(formations)))
	dtoFormations := make([]dto.Formation, 0, len(formations))
	for _, formation := range formations {
		tirads, err := u.uziRepo.GetTiradsByID(ctx, formation.TiradsID)
		if err != nil {
			u.logger.Error("Get formation tirads")
			return nil, fmt.Errorf("get tirads [id %q]: %w", formation.TiradsID, err)
		}

		dtoFormation := mappers.MustTransformObj[entity.Formation, dto.Formation](&formation)
		dtoFormation.Tirads = tirads

		dtoFormations = append(dtoFormations, *dtoFormation)
	}
	u.logger.Debug("[Response] Get formation tirads")

	return dtoFormations, nil
}
