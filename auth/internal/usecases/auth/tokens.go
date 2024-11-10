package usecases

import (
	"yir/auth/internal/entity"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
)

func (u *AuthUseCase) generateUserTokenPair(id uuid.UUID) (*entity.TokensPair, string, error) {
	refreshTokenWord := gofakeit.MinecraftVillagerJob()

	tokenPair, err := u.jwt.GeneratePair(
		map[string]any{
			"id":        id.String(),
			"likesFood": gofakeit.MinecraftFood(),
		},
		refreshTokenWord,
	)
	if err != nil {
		return nil, "", err
	}

	return tokenPair, refreshTokenWord, nil
}
