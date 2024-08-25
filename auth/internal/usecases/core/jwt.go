package core

import (
	"yir/auth/internal/entity"
)

type JWTService interface {
	GeneratePair(claims map[string]any, refreshWord string) (*entity.TokensPair, error)
	ParseUserData(refreshToken string) (*entity.UserTokenVerify, error)
}
