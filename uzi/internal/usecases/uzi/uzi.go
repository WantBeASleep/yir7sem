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

func (u *UziUseCase) InsertUzi(ctx context.Context, req *dto.Uzi) error {
	u.logger.Debug("[Request] Insert Uzi", zap.Any("Data", req.UziInfo))
	if err := u.uziRepo.InsertUzi(ctx, req.UziInfo); err != nil {
		u.logger.Error("Insert Uzi", zap.Error(err))
		return fmt.Errorf("insert uzi: %w", err)
	}
	u.logger.Debug("[Response] Inserted Uzi")

	u.logger.Debug("[Request] Insert images")
	if err := u.uziRepo.InsertImages(ctx, req.Images); err != nil {
		u.logger.Error("Insert images", zap.Error(err))
		return fmt.Errorf("insert images: %w", err)
	}
	u.logger.Debug("[Response] Inserted images")

	u.logger.Debug("[Request] Create formations tirads")
	formations, err := mappers.TransformSlice[dto.Formation, entity.Formation](req.Formations, func(src *dto.Formation, dst *entity.Formation) error {
		tiradsID, err := u.uziRepo.CreateTirads(ctx, &src.Tirads)
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
	if err := u.uziRepo.InsertFormations(ctx, formations); err != nil {
		u.logger.Error("Insert formations", zap.Error(err))
		return fmt.Errorf("insert formations: %w", err)
	}
	u.logger.Debug("[Response] Inserted formations")

	u.logger.Debug("[Request] Create segments tirads")
	segments, err := mappers.TransformSlice[dto.Segment, entity.Segment](req.Segments, func(src *dto.Segment, dst *entity.Segment) error {
		tiradsID, err := u.uziRepo.CreateTirads(ctx, &src.Tirads)
		if err != nil {
			return fmt.Errorf("create tirads: %w", err)
		}

		dst.TiradsID = tiradsID
		return nil
	})
	if err != nil {
		u.logger.Error("Create segments tirads", zap.Error(err))
		return fmt.Errorf("create segments tirads: %w", err)
	}
	u.logger.Debug("[Response] Created segments tirads")

	u.logger.Debug("[Request] Insert segments")
	if err := u.uziRepo.InsertSegments(ctx, segments); err != nil {
		u.logger.Error("Insert segments", zap.Error(err))
		return fmt.Errorf("insert segments: %w", err)
	}
	u.logger.Debug("[Response] Inserted segments")

	return nil
}

func (u *UziUseCase) GetUzi(ctx context.Context, id uuid.UUID) (*dto.Uzi, error) {
	u.logger.Debug("[Request] Get Uzi", zap.Any("uzi id", id))
	uzi, err := u.uziRepo.GetUzi(ctx, id)
	if err != nil {
		u.logger.Error("Get uzi", zap.Error(err))
		return nil, fmt.Errorf("get uzi: %w", err)
	}
	u.logger.Debug("[Response] Get uzi", zap.Any("Uzi", uzi))

	u.logger.Debug("[Request] Get uzi images", zap.Any("uzi id", id))
	images, err := u.uziRepo.GetUziImages(ctx, id)
	if err != nil {
		u.logger.Error("Get uzi images", zap.Error(err))
		return nil, fmt.Errorf("get uzi images: %w", err)
	}
	u.logger.Debug("[Response] Get uzi images", zap.Any("Images", images))

	u.logger.Debug("[Request] Get uzi formations", zap.Any("uzi id", id))
	formations, err := u.uziRepo.GetUziFormations(ctx, id)
	if err != nil {
		u.logger.Error("Get uzi formations", zap.Error(err))
		return nil, fmt.Errorf("get uzi formations: %w", err)
	}
	u.logger.Debug("[Response] Get uzi formations", zap.Any("Formations", formations))

	u.logger.Debug("[Request] Get uzi segments", zap.Any("uzi id", id))
	segments, err := u.uziRepo.GetUziSegments(ctx, id)
	if err != nil {
		u.logger.Error("Get uzi segments", zap.Error(err))
		return nil, fmt.Errorf("get uzi segments: %w", err)
	}
	u.logger.Debug("[Response] Get uzi segments", zap.Any("Segments", segments))

	u.logger.Debug("[Request] Get formations tirads")
	dtoFormations, err := mappers.TransformSlice[entity.Formation, dto.Formation](formations, func(src *entity.Formation, dst *dto.Formation) error {
		tirads, err := u.uziRepo.GetTirads(ctx, src.TiradsID)
		if err != nil {
			return fmt.Errorf("get tirads [id %q]: %w", src.TiradsID, err)
		}

		dst.Tirads = *tirads
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("get ")
	}

	return &dto.Uzi{
		UziInfo: uzi,
		Images:  images,
	}, nil
}
