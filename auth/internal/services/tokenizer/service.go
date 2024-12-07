package tokenizer

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GeneratePair(ctx context.Context, data map[string]any) (string, string, error)
	ParseClaims(ctx context.Context, token string) (map[string]any, error)
}

type service struct {
	accessLifeTime  time.Duration
	refreshLifeTime time.Duration
	privateKey      *rsa.PrivateKey
	publicKey       *rsa.PublicKey
}

func New(
	accessLifeTime time.Duration,
	refreshLifeTime time.Duration,
	privateKey *rsa.PrivateKey,
	publicKey *rsa.PublicKey,
) Service {
	return &service{
		accessLifeTime:  accessLifeTime,
		refreshLifeTime: refreshLifeTime,
		privateKey:      privateKey,
		publicKey:       publicKey,
	}
}

func (s *service) GeneratePair(ctx context.Context, data map[string]any) (string, string, error) {
	var token *jwt.Token
	claims := jwt.MapClaims(data)

	claims["exp"] = time.Now().Add(s.accessLifeTime).Unix()
	token = jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	accessToken, err := token.SignedString(s.privateKey)
	if err != nil {
		return "", "", fmt.Errorf("signed access_token: %w", err)
	}

	claims["exp"] = time.Now().Add(s.refreshLifeTime).Unix()
	claims["arcane_2"] = "disappointment" // change my mind
	token = jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	refreshToken, err := token.SignedString(s.privateKey)
	if err != nil {
		return "", "", fmt.Errorf("signed refresh_token: %w", err)
	}

	return accessToken, refreshToken, nil
}

func (s *service) ParseClaims(ctx context.Context, token string) (map[string]any, error) {
	parsed, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) { return s.publicKey, nil })
	if err != nil || !parsed.Valid {
		return nil, err
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("unknow jwt error")
	}

	return claims, nil
}
