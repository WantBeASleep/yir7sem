package secret

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func LoadPublicKey() (*rsa.PublicKey, error) {
	publicKeyEnv := os.Getenv("TOKEN_PUBLIC_KEY")
	if publicKeyEnv == "" {
		return nil, fmt.Errorf("TOKEN_PUBLIC_KEY environment variable is not set")
	}

	block, _ := pem.Decode([]byte(publicKeyEnv))
	if block == nil {
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
