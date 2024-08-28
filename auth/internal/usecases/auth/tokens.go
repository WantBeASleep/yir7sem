package usecases

import (
	"yir/auth/internal/entity"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (a *AuthUseCase) generateUserTokenPair(UUID uuid.UUID, refreshTokenWord string) (*entity.TokensPair, error) {
	a.logger.Info("[Request] Generate new JWT pair tokens")
	tokenPair, err := a.jwtService.GeneratePair(
		map[string]any{
			"uuid":      UUID,
			"likesFood": gofakeit.MinecraftFood(),
		},
		refreshTokenWord,
	)
	if err != nil {
		a.logger.Error("Generate tokens", zap.Error(err))
		return nil, err
	}
	a.logger.Info("[Response] Generated JWT Tokens")

	return tokenPair, nil
}
