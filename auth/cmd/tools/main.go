package main

import (
	"context"
	"flag"
	"yir/auth/internal/config"
	"yir/auth/internal/entity"
	"yir/auth/internal/repositories/db/repositories"

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

	login := "smt@smt.ri"
	pass := "Wr0ngpass"

	salt := "pensioner"

	hash, _ := entity.HashByScrypt(pass, salt)

	UUID, _ := uuid.NewRandom()

	authRepo.CreateUser(context.TODO(), &entity.UserCreditals{
		UUID:             UUID,
		Login:            login,
		PasswordHash:     hash + salt,
		RefreshTokenWord: "tachanka",
	})
}
