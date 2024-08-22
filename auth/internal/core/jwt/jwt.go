package jwt

import (
	"fmt"
	"time"
	"yir/auth/internal/enity"

	"github.com/golang-jwt/jwt/v5"
)

type Service struct {
	accessLifeTime  time.Duration
	refreshLifeTime time.Duration
	privateKey      string
}

func NewService(
	accessLifeTime time.Duration,
	refreshLifeTime time.Duration,
	privateKey string,
) *Service {
	return &Service{
		accessLifeTime:  accessLifeTime,
		refreshLifeTime: refreshLifeTime,
		privateKey:      privateKey,
	}
}

// rt claims add "rtw"!
func (s *Service) Generate(claims map[string]any, refreshWord string) (*enity.TokenPair, error) {
	var token *jwt.Token

	accessClaims := jwt.MapClaims{}
	for k, v := range claims {
		accessClaims[k] = v
	}
	accessClaims["exp"] = time.Now().Add(s.accessLifeTime).Unix()

	token = jwt.NewWithClaims(jwt.SigningMethodRS256, accessClaims)
	accessToken, err := token.SignedString(s.privateKey)
	if err != nil {
		return nil, fmt.Errorf("access token signed: %w", err)
	}

	refreshClaims := jwt.MapClaims{}
	for k, v := range claims {
		refreshClaims[k] = v
	}
	refreshClaims["rtw"] = refreshWord
	refreshClaims["exp"] = time.Now().Add(s.refreshLifeTime).Unix()

	token = jwt.NewWithClaims(jwt.SigningMethodRS256, refreshClaims)
	refreshToken, err := token.SignedString(s.privateKey)
	if err != nil {
		return nil, fmt.Errorf("refresh token signed: %w", err)
	}

	return &enity.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
