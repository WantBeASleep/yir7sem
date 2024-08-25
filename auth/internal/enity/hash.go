package enity

import (
	"crypto/sha256"
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

// Deprecated: В бд соль пишем после пароля, по фиксированному 64 в длинну паролю.
// Получается пароль(64 символа) + соль
// Этот вернет пароль 32 длинной, использовать, если понимаете что делаете
func HashBySHA256(password string, salt string) (string, error) {
	hash := sha256.Sum256([]byte(password + salt))
	return string(hash[:]), nil
}
