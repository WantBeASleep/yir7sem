package main

import (
	"context"
	"flag"
	"fmt"
	"yir/auth/internal/config"
	"yir/auth/internal/entity"
	"yir/auth/internal/repositories/db/repositories"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
)

const (
	defaultCfgPath = "config/config.yml"
	shorthand      = " (shorthand)"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", defaultCfgPath, "set config path")
	flag.StringVar(&configPath, "c", defaultCfgPath, "set config path"+shorthand)
}

func main() {
	flag.Parse()
	cfg := config.MustLoad(configPath)

	authRepo, _ := repositories.NewRepository(&cfg.DB)

	for range 100 {
		UUID := uuid.New()
		login := gofakeit.Email()
		pass := gofakeit.Password(true, true, true, false, false, 15)
		salt := gofakeit.Word()
		refresh := gofakeit.Word()
		hash, _ := entity.HashByScrypt(pass, salt)

		fmt.Println(login, pass)
		authRepo.CreateUser(context.Background(), &entity.UserCreditals{
			UUID:             UUID,
			Login:            login,
			PasswordHash:     hash + salt,
			RefreshTokenWord: refresh,
			MedWorkerUUID:    uuid.New(),
		})
	}
}
