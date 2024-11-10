package entity

import (
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/scrypt"
)

func HashByScrypt(password string, salt string) (string, error) {
	hexHash, err := scrypt.Key(
		[]byte(password),
		[]byte(salt),
		1<<13,
		8,
		1,
		32,
	)
	if err != nil {
		return "", fmt.Errorf("hash pass: %v", err)
	}

	hash := hex.EncodeToString(hexHash)
	return string(hash), nil
}
