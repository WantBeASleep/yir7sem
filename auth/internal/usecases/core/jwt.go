package core

import (
	"yir/auth/internal/enity"
)

type JWTService interface {
	Generate(claims map[string]any, refreshWord string) (*enity.TokenPair, error)
}
