package usecases

import (
	"context"
	"yir/auth/internal/enity"

	"github.com/brianvoe/gofakeit/v7"
	"go.uber.org/zap"
)

// внедрить контекст в jwt
func (a *AuthUseCase) generateTokenPair(_ context.Context, userID int, refreshTokenWord string) (*enity.TokensPair, error) {
	a.logger.Info("[Request] Generate new JWT pair tokens")
	tokenPair, err := a.jwtService.GeneratePair(
		map[string]any{
			"id":        userID,
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
