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

	segmentsIDS, err := u.CreateDTOSegments(ctx, req.Segments)
	if err != nil {
		return uuid.Nil, nil, fmt.Errorf("insert dto segments: %w", err)
	}

	return formationID, segmentsIDS, nil
}

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

	dtoFormation, err := u.GetDTOFormation(ctx, formation)
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

	dtoSegments, err := u.GetDTOSegments(ctx, segments)
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

	dtoUpdateFormation, err := u.GetDTOFormation(ctx, updateFormation)
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
	formationID, err := u.uziRepo.CreateFormation(ctx, entFormation)
	if err != nil {
		u.logger.Error("Insert formations", zap.Error(err))
		return uuid.Nil, fmt.Errorf("insert formations: %w", err)
	}
	u.logger.Debug("[Response] Created dto formation")

	return formationID, nil
}

func (u *UziUseCase) GetDTOFormation(ctx context.Context, formation *entity.Formation) (*dto.Formation, error) {
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

func (u *UziUseCase) GetDTOFormations(ctx context.Context, formations []entity.Formation) ([]dto.Formation, error) {
	dtoFormations := make([]dto.Formation, 0, len(formations))
	for _, formation := range formations {
		u.logger.Debug("[Request] Get formation tirads")
		tirads, err := u.uziRepo.GetTiradsByID(ctx, formation.TiradsID)
		if err != nil {
			u.logger.Error("Get formation tirads")
			return nil, fmt.Errorf("get tirads [id %q]: %w", formation.TiradsID, err)
		}
		u.logger.Debug("[Response] Get formation tirads")

		dtoFormation := mappers.MustTransformObj[entity.Formation, dto.Formation](&formation)
		dtoFormation.Tirads = tirads

		dtoFormations = append(dtoFormations, *dtoFormation)
	}

	return dtoFormations, nil
}
