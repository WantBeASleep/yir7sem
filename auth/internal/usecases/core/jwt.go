package core

import (
	"yir/auth/internal/enity"
)

type JWTService interface {
	GeneratePair(claims map[string]any, refreshWord string) (*enity.TokensPair, error)
	ParseUserData(refreshToken string) (*enity.UserTokenVerify, error)
}
