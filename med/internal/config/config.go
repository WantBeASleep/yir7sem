package config

type Config struct {
	App App
	DB  DB
}

type App struct {
	Url string `env:"url" env-default:"localhost:50060"`
}

type DB struct {
	Dsn string `env:"DB_DSN" env-required:"true"`
}
