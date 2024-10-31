package uzi

import (
	"context"
	"fmt"
	"yir/pkg/mappers"
	"yir/uzi/internal/entity"
	"yir/uzi/internal/entity/dto"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (u *UziUseCase) InsertFormationWithSegments(ctx context.Context, req *dto.FormationWithSegments) error {
	if err := u.InsertDTOFormations(ctx, []dto.Formation{*req.Formation}); err != nil {
		return fmt.Errorf("insert dto formation: %w", err)
	}

	if err := u.InsertDTOSegments(ctx, req.Segments); err != nil {
		return fmt.Errorf("insert dto segments: %w", err)
	}

	return nil
}

// generates ID
func (u *UziUseCase) CreateFormationsWithSegments(ctx context.Context, req []dto.FormationWithSegments) ([]dto.FormationWithSegmentsIDs, error) {
	resp := make([]dto.FormationWithSegmentsIDs, 0, len(req))
	for i, formation := range req {
		formationID, _ := uuid.NewRandom()
		req[i].Formation.Id = formationID

		segmentsIDs := make(uuid.UUIDs, 0, len(formation.Segments))
		for j := range formation.Segments {
			segmentID, _ := uuid.NewRandom()
			segmentsIDs = append(segmentsIDs, segmentID)
			req[i].Segments[j].Id = segmentID
		}

		if err := u.InsertFormationWithSegments(ctx, &req[i]); err != nil {
			return nil, fmt.Errorf("insert formation [number %d] with segments: %w", i, err)
		}

		resp = append(resp, dto.FormationWithSegmentsIDs{
			FormationID: formationID,
			Segments:    segmentsIDs,
		})
	}

	return resp, nil
}

func (u *UziUseCase) GetFormationWithSegments(ctx context.Context, id uuid.UUID) (*dto.FormationWithSegments, error) {
	u.logger.Debug("[Request] Get formation by ID", zap.Any("id", id))
	formation, err := u.uziRepo.GetFormationByID(ctx, id)
	if err != nil {
		u.logger.Error("Get formation by ID", zap.Error(err))
		return nil, fmt.Errorf("get formation by ID: %w", err)
	}
	u.logger.Debug("[Response] Got formation by ID", zap.Any("formation", formation))

	u.logger.Debug("[Request] Get formation segments", zap.Any("uzi id", id))
	segments, err := u.uziRepo.GetFormationSegments(ctx, id)
	if err != nil {
		u.logger.Error("Get formation segments", zap.Error(err))
		return nil, fmt.Errorf("get formation segments: %w", err)
	}
	u.logger.Debug("[Response] Get formation segments", zap.Any("Segments", segments))

	dtoFormation, err := u.GetDTOFormation(ctx, []entity.Formation{*formation})
	if err != nil {
		return nil, fmt.Errorf("get dto formation: %w", err)
	}

	dtoSegments, err := u.GetDTOSegments(ctx, segments)
	if err != nil {
		return nil, fmt.Errorf("get dto segments: %w", err)
	}

	return &dto.FormationWithSegments{
		Formation: &dtoFormation[0],
		Segments:  dtoSegments,
	}, nil
}

func (u *UziUseCase) UpdateFormation(ctx context.Context, id uuid.UUID, req *dto.Formation) error {
	formation, err := mappers.TransformObj[dto.Formation, entity.Formation](req, func(src *dto.Formation, dst *entity.Formation) error {
		if src.Tirads != nil {
			u.logger.Debug("[Request] Create tirads")
			tiradsID, err := u.uziRepo.CreateTirads(ctx, src.Tirads)
			if err != nil {
				u.logger.Error("Create tirads", zap.Error(err))
				return fmt.Errorf("create tirads: %w", err)
			}
			u.logger.Debug("[Response] Created tirads", zap.Int("id", tiradsID))

			dst.TiradsID = tiradsID
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("get entity formation: %w", err)
	}

	u.logger.Debug("[Request] Update Formation", zap.Any("id", id), zap.Any("data", req))
	if err := u.uziRepo.UpdateFormation(ctx, id, formation); err != nil {
		u.logger.Error("Update Formation", zap.Error(err))
		return fmt.Errorf("update formation: %w", err)
	}
	u.logger.Debug("[Response] Updated Formation")

	return nil
}

func (u *UziUseCase) InsertDTOFormations(ctx context.Context, formations []dto.Formation) error {
	u.logger.Debug("[Request] Create formations tirads")
	entityFormations, err := mappers.TransformSlice[dto.Formation, entity.Formation](formations, func(src *dto.Formation, dst *entity.Formation) error {
		tiradsID, err := u.uziRepo.CreateTirads(ctx, src.Tirads)
		if err != nil {
			return fmt.Errorf("create tirads: %w", err)
		}

		dst.TiradsID = tiradsID
		return nil
	})
	if err != nil {
		u.logger.Error("Create formations tirads", zap.Error(err))
		return fmt.Errorf("create formations tirads: %w", err)
	}
	u.logger.Debug("[Response] Created formations tirads")

	u.logger.Debug("[Request] Insert formations")
	if err := u.uziRepo.InsertFormations(ctx, entityFormations); err != nil {
		u.logger.Error("Insert formations", zap.Error(err))
		return fmt.Errorf("insert formations: %w", err)
	}
	u.logger.Debug("[Response] Inserted formations")

	return nil
}

func (u *UziUseCase) GetDTOFormation(ctx context.Context, formations []entity.Formation) ([]dto.Formation, error) {
	u.logger.Debug("[Request] Get formations tirads")
	dtoFormations, err := mappers.TransformSlice[entity.Formation, dto.Formation](formations, func(src *entity.Formation, dst *dto.Formation) error {
		tirads, err := u.uziRepo.GetTiradsByID(ctx, src.TiradsID)
		if err != nil {
			return fmt.Errorf("get tirads [id %q]: %w", src.TiradsID, err)
		}

		dst.Tirads = tirads
		return nil
	})
	if err != nil {
		u.logger.Error("Get formations tirads", zap.Error(err))
		return nil, fmt.Errorf("get formations tirads: %w", err)
	}
	u.logger.Debug("[Response] Got formations tirads")

	return dtoFormations, nil
}
