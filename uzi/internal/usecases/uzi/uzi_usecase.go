package uzi

import (
	"yir/uzi/internal/usecases/repositories"

	"go.uber.org/zap"
)

type UziUseCase struct {
	uziRepo repositories.UziRepo

	logger *zap.Logger
}

func NewUziUseCase(
	uziRepo repositories.UziRepo,

	logger *zap.Logger,
) *UziUseCase {
	return &UziUseCase{
		uziRepo: uziRepo,
		logger:  logger,
	}
}
