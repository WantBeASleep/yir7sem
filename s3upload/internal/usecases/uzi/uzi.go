package uzi

import (
	"bytes"
	"context"
	"fmt"
	"path/filepath"

	"image"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/tiff"

	"yir/s3upload/internal/entity"
	"yir/s3upload/internal/usecases/repo"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type UziUseCase struct {
	s3 repo.S3

	logger *zap.Logger
}

func NewUziUseCase(
	s3 repo.S3,

	logger *zap.Logger,
) *UziUseCase {
	return &UziUseCase{
		s3: s3,
		logger: logger,
	}
}


func (u *UziUseCase) UploadAndSplitUziFile(ctx context.Context, img []byte) (uuid.UUID, uuid.UUIDs, error) {
	
	_, format, err := image.DecodeConfig(bytes.NewBuffer(img))
	if err != nil {
		return uuid.Nil, nil, fmt.Errorf("decode img format: %w", err)
	}

	// сюда нужен очевидно стейт machine, что бы ретраить, потом прикрутим temporal наверное
	switch format {
	case "tiff":
		
	default:
		return uuid.Nil, nil, fmt.Errorf("decode img format: %w", err)
	}

	return uuid.Nil, nil, nil
}

func (u *UziUseCase) UploadSplittingUzi(ctx context.Context, mainFile *entity.ImageDataWithFormat, splitted []entity.ImageDataWithFormat) (uuid.UUID, uuid.UUIDs, error) {
	mainID, err := uuid.NewRandom()
	if err != nil {
		return uuid.Nil, nil, fmt.Errorf("generate main id: %w", err)
	}

	u.logger.Info("[Request] Insert main file in S3", zap.Any("id", mainID))
	if err := u.s3.Upload(ctx, mainID.String(), fmt.Sprintf("%s.%s", mainID.String(), mainFile.Format), mainFile.Image); err != nil {
		u.logger.Error("Insert main file in S3", zap.Error(err))
		return uuid.Nil, nil, fmt.Errorf("insert main file to S3: %w", err)
	}
	u.logger.Info("[Response] Insert edsplitted files in S3", zap.Any("id", mainID))

	splittedIDs := uuid.UUIDs{}

	u.logger.Info("[Request] Insert splitted files in S3", zap.Any("id", mainID))
	for i, v := range splitted {

		splittedID, err := uuid.NewRandom()
		if err != nil {
			return uuid.Nil, nil, fmt.Errorf("generate split id: %w", err)
		}
		splittedIDs = append(splittedIDs, splittedID)

		if err := u.s3.Upload(ctx, filepath.Join(mainID.String(), splittedID.String()), fmt.Sprintf("%s.%s", splittedID.String(), v.Format), v.Image); err != nil {
			u.logger.Error("Insert splitted file in S3", zap.Int("number of splitted", i + 1), zap.Error(err))
			return uuid.Nil, nil, fmt.Errorf("insert splitted file [index %q] to S3: %w", i, err)
		}
	}
	u.logger.Info("[Response] Inserted splitted files in S3", zap.Any("id", mainID))

	return mainID, splittedIDs, nil
}