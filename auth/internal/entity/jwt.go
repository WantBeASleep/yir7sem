package entity

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"time"
	"yir/auth/internal/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWT struct {
	accessLifeTime  time.Duration
	refreshLifeTime time.Duration
	privateKey      *rsa.PrivateKey
	publicKey       *rsa.PublicKey
}

func NewService(cfg *config.Token) (*JWT, error) {
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

	return &JWT{
		accessLifeTime:  cfg.AccessLifeTime,
		refreshLifeTime: cfg.RefreshLifeTime,
		privateKey:      privateKey.(*rsa.PrivateKey),
		publicKey:       publicKey.(*rsa.PublicKey),
	}, nil
}

// Add refresh word as "rtw" key.
func (s *JWT) GeneratePair(claims map[string]any, refreshWord string) (*TokensPair, error) {
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

	return &TokensPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// парсит user id, и refresh token
func (s *JWT) ParseUserIDAndRW(refreshToken string) (uuid.UUID, string, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(
		refreshToken,
		claims,
		func(t *jwt.Token) (interface{}, error) { return s.publicKey, nil },
		jwt.WithExpirationRequired(),
	)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return uuid.Nil, "", ErrTokenExpired
		}
		return uuid.Nil, "", err
	}

	rtw, err := s.parseRTW(claims)
	if err != nil {
		return uuid.Nil, "", ErrInvalidToken
	}

	id, err := s.parseID(claims)
	if err != nil {
		return uuid.Nil, "", ErrInvalidToken
	}

	return id, rtw, nil
}

func (s *JWT) parseRTW(claims jwt.MapClaims) (string, error) {
	rtwInterface, ok := claims["rtw"]
	if !ok {
		return "", ErrInvalidToken
	}

	tokenWord, ok := rtwInterface.(string)
	if !ok {
		return "", ErrInvalidToken
	}

	return tokenWord, nil
}

func (s *JWT) parseID(claims jwt.MapClaims) (uuid.UUID, error) {
	idInterface, ok := claims["id"]
	if !ok {
		return uuid.Nil, ErrInvalidToken
	}

	idStr, ok := idInterface.(string)
	if !ok {
		return uuid.Nil, ErrInvalidToken
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.Nil, ErrInvalidToken
	}

	return id, nil
}
