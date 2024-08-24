package jwt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"
	"yir/auth/internal/config"
	"yir/auth/internal/enity"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type Service struct {
	accessLifeTime  time.Duration
	refreshLifeTime time.Duration
	privateKey      *rsa.PrivateKey

	logger *zap.Logger
}

func NewService(cfg *config.Token, logger *zap.Logger) (*Service, error) {
	block, _ := pem.Decode([]byte(cfg.PrivateKey))

	pkey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parse private key: %w", err)
	}

	return &Service{
		accessLifeTime:  cfg.AccessLifeTime,
		refreshLifeTime: cfg.RefreshLifeTime,
		privateKey:      pkey.(*rsa.PrivateKey),

		logger: logger,
	}, nil
}

// rt claims add "rtw"!
func (s *Service) Generate(claims map[string]any, refreshWord string) (*enity.TokenPair, error) {
	var token *jwt.Token

	accessClaims := jwt.MapClaims{}
	for k, v := range claims {
		accessClaims[k] = v
	}
	accessClaims["exp"] = time.Now().Add(s.accessLifeTime).Unix()

	s.logger.Debug("Access claims ready", zap.Any("claims", accessClaims))

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

	s.logger.Debug("Refresh claims ready", zap.Any("claims", refreshClaims))

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
