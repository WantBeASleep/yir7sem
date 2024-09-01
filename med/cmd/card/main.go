package main

import (
	"flag"
	"log"
	"yir/med/internal/config"
)

const (
	envDev  = "dev"
	envProd = "prod"
)

const (
	defaultCfgPath = "med/config/config.yaml"
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
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err.Error())
	}

	log.Printf("config: %+v", cfg)
}
