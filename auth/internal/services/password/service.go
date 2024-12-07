package password

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"golang.org/x/crypto/scrypt"
)

const (
	iterationCount   = 1 << 13
	memoryCountForOp = 8
	parralelism      = 1
	passlen          = 32
)

type Service interface {
	Hash(ctx context.Context, password, salt string) (string, error)
	GetSalt(ctx context.Context, hash string) (string, error)
	GenerateSalt(ctx context.Context) string
}

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) Hash(ctx context.Context, password, salt string) (string, error) {
	hexHash, err := scrypt.Key(
		[]byte(password),
		[]byte(salt),
		iterationCount,
		memoryCountForOp,
		parralelism,
		passlen,
	)
	if err != nil {
		return "", err
	}

	hash := hex.EncodeToString(hexHash)
	return fmt.Sprintf("%s%s", hash, salt), err
}

func (s *service) GetSalt(ctx context.Context, hash string) (string, error) {
	if len(hash) <= 64 {
		return "", errors.New("pass is too short, not out hashing")
	}

	return hash[64:], nil
}

func (s *service) GenerateSalt(ctx context.Context) string {
	salt := gofakeit.MinecraftAnimal()

	hash := sha256.New()
	hash.Write([]byte(salt))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	if len(hashString) > 64 {
		hashString = hashString[64:]
	}

	return hashString
}
