package uzi

import (
	"context"
	"fmt"
	"path/filepath"

	"yir/uzi/internal/entity/imagesplitter"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (u *UziUseCase) SplitUziAndLoadS3(ctx context.Context, uziID uuid.UUID) (uuid.UUIDs, error) {
	// какой же колхоз то блять, надо все это в мидлвары обернуть
	u.logger.Info("[Request] SplitUziAndLoadS3 request")

	u.logger.Info("[Request] Get uzi file from S3", zap.String("Uzi ID", uziID.String()))
	uzi, err := u.s3Repo.FullGetByPath(ctx, filepath.Join(uziID.String(), uziID.String()))
	if err != nil {
		u.logger.Error("Get uzi file from S3", zap.Error(err))
		return nil, fmt.Errorf("get uzi from s3: %w", err)
	}
	u.logger.Info("[Response] Got uzi file from S3")

	splitted, err := imagesplitter.SplitToPng(uzi)
	if err != nil {
		return nil, fmt.Errorf("split uzi file: %w", err)
	}

	splittedIDs := uuid.UUIDs{}

	u.logger.Info("[Request] Upload pages to S3", zap.Int("pages count", len(splitted)))
	for i, split := range splitted {
		id, _ := uuid.NewRandom()
		splittedIDs = append(splittedIDs, id)

		if err := u.s3Repo.Upload(ctx, filepath.Join(uziID.String(), id.String(), id.String()), split); err != nil {
			u.logger.Error("page upload err", zap.Int("page number", i+1), zap.Error(err))
			return nil, fmt.Errorf("upload page [page number %q]: %w", i+1, err)
		}
	}
	u.logger.Info("[Response] Uploaded all pages to S3")

	return splittedIDs, nil
}
