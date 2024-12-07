package config

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"
)

type Config struct {
	App App
	DB  DB
	JWT JWT
}

type App struct {
	Url string `env:"APP_URL" env-default:"localhost:50055"`
}

type DB struct {
	Dsn string `env:"DB_DSN" env-required:"true"`
}

type JWT struct {
	AccessTokenTime  time.Duration `env:"JWT_TOKEN_AC_TIME" env-default:"5m"`
	RefreshTokenTime time.Duration `env:"JWT_TOKEN_RE_TIME" env-default:"24h"`
	RsaPublicKey     string        `env:"JWT_KEY_PUBLIC" env-required:"true"`
	RsaPrivateKey    string        `env:"JWT_KEY_PRIVATE" env-required:"true"`
}

func (c Config) ParseRsaKeys() (*rsa.PublicKey, *rsa.PrivateKey, error) {
	publicBlock, _ := pem.Decode([]byte(c.JWT.RsaPublicKey))
	publicKey, err := x509.ParsePKIXPublicKey(publicBlock.Bytes)
	if err != nil {
		return nil, nil, fmt.Errorf("parse public key: %w", err)
	}

	privateBlock, _ := pem.Decode([]byte(c.JWT.RsaPrivateKey))
	privateKey, err := x509.ParsePKCS1PrivateKey(privateBlock.Bytes)
	if err != nil {
		return nil, nil, fmt.Errorf("parse public key: %w", err)
	}

	return publicKey.(*rsa.PublicKey), privateKey, nil
}
