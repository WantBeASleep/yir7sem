package enity

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/scrypt"
)

// return hash + salt
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

func HashBySHA256(password string) (string, error) {
	hash := sha256.Sum256([]byte(password))
	// [32] не слайс а именно МАССИВ: курите разницу массивов и срезов golang
	return string(hash[:]), nil
}
