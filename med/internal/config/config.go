package config

type Config struct {
	App App `yaml:"app"`
	DB  DB  `yaml:"db"`
}

type App struct {
	Url string `yaml:"url" env:"url" env-default:"localhost:50060"`
}

type DB struct {
	Dsn string `yaml:"dsn" env:"DB_DSN" env-required:"true"`
}
