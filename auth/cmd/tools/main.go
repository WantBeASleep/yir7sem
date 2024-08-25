package main

import (
	"context"
	"flag"
	"yir/auth/internal/config"
	"yir/auth/internal/entity"
	"yir/auth/internal/repositories/db/repositories"
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

	authRepo.CreateUser(context.TODO(), &entity.User{
		Login:            login,
		PasswordHash:     hash + salt,
		RefreshTokenWord: "tachanka",
	})
}
