package jwt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"time"
	"yir/auth/internal/config"
	"yir/auth/internal/enity"

	"github.com/golang-jwt/jwt/v5"
)

type Service struct {
	accessLifeTime  time.Duration
	refreshLifeTime time.Duration
	privateKey      *rsa.PrivateKey
	publicKey       *rsa.PublicKey
}

func NewService(cfg *config.Token) (*Service, error) {
	privateBlock, _ := pem.Decode([]byte(cfg.PrivateKey))
	privateKey, err := x509.ParsePKCS8PrivateKey(privateBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parse private key: %w", err)
	}

	publicBlock, _ := pem.Decode([]byte(cfg.PublicKey))
	publicKey, err := x509.ParsePKIXPublicKey(publicBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parse public key: %w", err)
	}

	return &Service{
		accessLifeTime:  cfg.AccessLifeTime,
		refreshLifeTime: cfg.RefreshLifeTime,
		privateKey:      privateKey.(*rsa.PrivateKey),
		publicKey:       publicKey.(*rsa.PublicKey),
	}, nil
}

// Add refresh word as "rtw" key.
func (s *Service) GeneratePair(claims map[string]any, refreshWord string) (*enity.TokensPair, error) {
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

	return &enity.TokensPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Service) ParseUserData(refreshToken string) (*enity.UserTokenVerify, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(
		refreshToken,
		claims,
		func(t *jwt.Token) (interface{}, error) { return s.publicKey, nil },
		jwt.WithExpirationRequired(),
	)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, enity.ErrTokenExpired
		}
		return nil, err
	}

	rtw, err := s.parseRTW(claims)
	if err != nil {
		return nil, enity.ErrInvalidToken
	}

	ID, err := s.parseID(claims)
	if err != nil {
		return nil, enity.ErrInvalidToken
	}

	return &enity.UserTokenVerify{
		UserID:           ID,
		RefreshTokenWord: rtw,
	}, nil
}

func (s *Service) parseRTW(claims jwt.MapClaims) (string, error) {
	rtwInterface, ok := claims["rtw"]
	if !ok {
		return "", enity.ErrInvalidToken
	}

	tokenWord, ok := rtwInterface.(string)
	if !ok {
		return "", enity.ErrInvalidToken
	}

	return tokenWord, nil
}

func (s *Service) parseID(claims jwt.MapClaims) (int, error) {
	IDInterface, ok := claims["id"]
	if !ok {
		return 0, enity.ErrInvalidToken
	}

	ID, ok := IDInterface.(float64)
	if !ok {
		return 0, enity.ErrInvalidToken
	}

	return int(ID), nil
}
