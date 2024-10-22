package uzi

import (
	"context"
	"fmt"
	"io"
	"path/filepath"

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
		s3:     s3,
		logger: logger,
	}
}

func (u *UziUseCase) UploadAndSplitUziFile(ctx context.Context, img []byte) (uuid.UUID, uuid.UUIDs, error) {
	mainFile, err := addMetaToImageData(img)
	if err != nil {
		return uuid.Nil, nil, fmt.Errorf("convert main file: %w", err)
	}

	splitted, err := splitImageWithMeta(img)
	if err != nil {
		return uuid.Nil, nil, fmt.Errorf("splitting main file: %w", err)
	}

	return u.uploadSplittingUzi(ctx, mainFile, splitted)
}

func (u *UziUseCase) GetByPath(ctx context.Context, path string) (io.ReadCloser, error) {
	u.logger.Info("[Request] Get file from S3", zap.String("path", path))
	stream, err := u.s3.Get(ctx, path)
	if err != nil {
		u.logger.Error("Get file from S3", zap.Error(err))
		return nil, fmt.Errorf("get file [path %q]: %w", path, err)
	}
	u.logger.Info("[Response] Got file from S3")

	return stream, nil
}

func (u *UziUseCase) uploadSplittingUzi(ctx context.Context, mainFile *entity.ImageDataWithFormat, splitted []entity.ImageDataWithFormat) (uuid.UUID, uuid.UUIDs, error) {
	// сюда нужен очевидно стейт machine, что бы ретраить, потом прикрутим temporal наверное
	mainID, err := uuid.NewRandom()
	if err != nil {
		return uuid.Nil, nil, fmt.Errorf("generate main id: %w", err)
	}

	u.logger.Info("[Request] Insert main file in S3", zap.Any("id", mainID))
	err = u.s3.Upload(ctx,
		mainID.String(),
		mainID.String(),
		mainFile.Image,
		&entity.ImageMetaData{ContentType: mainFile.ContentType},
	)
	if err != nil {
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

		err = u.s3.Upload(ctx,
			filepath.Join(mainID.String(), splittedID.String()),
			splittedID.String(),
			v.Image,
			&entity.ImageMetaData{ContentType: v.ContentType},
		)
		if err != nil {
			u.logger.Error("Insert splitted file in S3", zap.Int("number of splitted", i+1), zap.Error(err))
			return uuid.Nil, nil, fmt.Errorf("insert splitted file [index %q] to S3: %w", i, err)
		}
	}
	u.logger.Info("[Response] Inserted splitted files in S3", zap.Any("id", mainID))

	return mainID, splittedIDs, nil
}
