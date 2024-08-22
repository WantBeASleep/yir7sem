package enity

import (
	"fmt"
	"crypto/sha256"

	"github.com/brianvoe/gofakeit/v7"
	"golang.org/x/crypto/scrypt"
)

func HashByScrypt(password string) (string, error) {
	hash, err := scrypt.Key(
		[]byte(password),
		[]byte(gofakeit.MinecraftFood()),
		1<<13,
		8,
		1,
		200, // 50 unicode символов по 4 байта
	)
	if err != nil {
		return "", fmt.Errorf("hash pass: %v", err)
	}

	return string(hash), nil
}

func HashBySHA256(password string) (string, error) {
	hash := sha256.Sum256([]byte(password))
	// [32] не слайс а именно МАССИВ: курите разницу массивов и срезов golang
	return string(hash[:]), nil
}