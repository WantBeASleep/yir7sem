package uzi

import (
	"context"
	"fmt"
	"yir/uzi/internal/entity"
	"yir/uzi/internal/usecases/mappers"
	"yir/uzi/internal/utils"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (u *UziUseCase) InsertUzi(ctx context.Context, req *entity.InsertUziRequest) error {
	u.logger.Debug("[Request] Insert Uzi", zap.Any("Data", req.Uzi))
	if err := u.uziRepo.InsertUzi(ctx, &req.Uzi); err != nil {
		u.logger.Error("Insert Uzi", zap.Error(err))
		return fmt.Errorf("insert uzi: %w", err)
	}
	u.logger.Debug("[Response] Inserted Uzi")

	images := mappers.HttpImagesToImages(req.Images, req.Uzi.Id)
	u.logger.Debug("[Request] Insert images")
	if err := u.uziRepo.InsertImages(ctx, images); err != nil {
		u.logger.Error("Insert images", zap.Error(err))
		return fmt.Errorf("insert images: %w", err)
	}
	u.logger.Debug("[Response] Inserted images")

	formations := utils.MustTransformSlice[entity.HttpFormation, entity.DBFormation](req.Formations)
	u.logger.Debug("[Request] Insert formations with image-formation")
	if err := u.uziRepo.InsertFormationsWithImageFormations(ctx, formations); err != nil {
		u.logger.Error("Insert formations with image-formations", zap.Error(err))
		return fmt.Errorf("insert formations with image-formations: %w", err)
	}
	u.logger.Debug("[Response] Inserted formations with image-formation")

	return nil
}

func (u *UziUseCase) GetMetaUzi(ctx context.Context, id uuid.UUID) (*entity.GetMetaUziResponse, error) {
	u.logger.Debug("[Request] Get Uzi", zap.Any("uzi id", id))
	uzi, err := u.uziRepo.GetUzi(ctx, id)
	if err != nil {
		u.logger.Error("Get uzi", zap.Error(err))
		return nil, fmt.Errorf("get uzi: %w", err)
	}
	u.logger.Debug("[Response] Get uzi", zap.Any("Uzi", uzi))

	u.logger.Debug("[Request] Get device", zap.Int("device id", uzi.DeviceID))
	device, err := u.uziRepo.GetDevice(ctx, uzi.DeviceID)
	if err != nil {
		u.logger.Error("Get device", zap.Error(err))
		return nil, fmt.Errorf("get device: %w", err)
	}
	u.logger.Debug("[Response] Get device", zap.Any("device", device))

	u.logger.Debug("[Request] Get uzi images", zap.Any("uzi id", id))
	images, err := u.uziRepo.GetUziImages(ctx, id)
	if err != nil {
		u.logger.Error("Get uzi images", zap.Error(err))
		return nil, fmt.Errorf("get uzi images: %w", err)
	}
	u.logger.Debug("[Response] Get uzi images")

	return mappers.UziDeviceImagesToUziMeta(uzi, device, images), nil
}
