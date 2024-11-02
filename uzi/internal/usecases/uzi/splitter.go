package uzi

import (
	"context"
	"fmt"
	"path/filepath"

	"yir/uzi/internal/entity"
	"yir/uzi/internal/entity/imagesplitter"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

// 1.загрузить из S3 && 2.Split uzi && 3.load to S3 && 4.load to postgres
// все декомпозировано, usecase реализует uzi_splitted_event
func (u *UziUseCase) SplitLoadSaveUzi(ctx context.Context, uziID uuid.UUID) (uuid.UUIDs, error) {
	// какой же колхоз то блять, надо все это в мидлвары обернуть
	u.logger.Info("[Request] SplitUziAndLoadS3 request")

	uziURL := filepath.Join(uziID.String(), uziID.String())

	u.logger.Info("[Request] Get uzi file from S3", zap.String("Uzi ID", uziID.String()), zap.String("Path", uziURL))
	uzi, err := u.s3Repo.GetFile(ctx, uziURL)
	if err != nil {
		u.logger.Error("Get uzi file from S3", zap.Error(err))
		return nil, fmt.Errorf("get uzi from s3: %w", err)
	}
	u.logger.Info("[Response] Got uzi file from S3")

	splitted, err := imagesplitter.SplitToPng(uzi)
	if err != nil {
		return nil, fmt.Errorf("split uzi file: %w", err)
	}

	imagesIDs := uuid.UUIDs{}
	images := make([]entity.Image, 0, len(splitted))

	u.logger.Info("[Request] Upload pages to S3", zap.Int("pages count", len(splitted)))
	for i, split := range splitted {
		id, _ := uuid.NewRandom()
		imagesIDs = append(imagesIDs, id)		
		url := filepath.Join(uziID.String(), id.String(), id.String())

		if err := u.s3Repo.Upload(ctx, url, &split); err != nil {
			u.logger.Error("page upload err", zap.Int("page number", i+1), zap.Error(err))
			return nil, fmt.Errorf("upload page [page number %q]: %w", i+1, err)
		}

		images = append(images, entity.Image{
			Id:    id,
			Url:   url,
			UziID: uziID,
			Page:  i + 1,
		})
	}
	u.logger.Info("[Response] Uploaded all pages to S3")

	if _, err := u.CreateImages(ctx, images); err != nil {
		return nil, fmt.Errorf("insert images to postgres: %w", err)
	}

	return imagesIDs, nil
}
