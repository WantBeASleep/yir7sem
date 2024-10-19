package jwt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func ValidateJWT(tokenString string, publicKey *rsa.PublicKey) error {
	Logger
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Убедитесь, что используется алгоритм RSA
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return err
	}

	// Проверка, валиден ли токен
	if !token.Valid {
		return fmt.Errorf("token is invalid")
	}

	// Дополнительная проверка срока действия токена
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("couldn't parse claims")
	}

	if exp, ok := claims["exp"].(float64); ok {
		expTime := time.Unix(int64(exp), 0)
		if time.Now().After(expTime) {
			return fmt.Errorf("token has expired")
		}
	}

	return nil
}

func LoadPublicKey() (*rsa.PublicKey, error) {
	publicKeyEnv := os.Getenv("PUBLIC_KEY")
	if publicKeyEnv == "" {
		return nil, fmt.Errorf("PUBLIC_KEY environment variable is not set")
	}

	block, _ := pem.Decode([]byte(publicKeyEnv))
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %v", err)
	}

	rsaPub, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA public key")
	}

	return rsaPub, nil
}
