package main

import (
	"flag"

	"yir/auth/internal/config"
	"yir/internal/log"
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

	logger := log.New(cfg.App.Env, "")

	logger.Debug("cfg && logger load!")
}
