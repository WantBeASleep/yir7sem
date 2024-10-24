package uzi

import (
	"yir/uzi/internal/usecases/repositories"

	"go.uber.org/zap"
)

type UziUseCase struct {
	uziRepo repositories.UziRepo
	s3Repo  repositories.S3

	logger *zap.Logger
}

func NewUziUseCase(
	uziRepo repositories.UziRepo,
	s3Repo repositories.S3,

	logger *zap.Logger,
) *UziUseCase {
	return &UziUseCase{
		uziRepo: uziRepo,
		s3Repo:  s3Repo,
		logger:  logger,
	}
}
