package config

type Config struct {
	App App
	DB  DB
}

type App struct {
	Url string `env:"APP_URL" env-default:"localhost:50060"`
}

type DB struct {
	Dsn string `env:"DB_DSN" env-required:"true"`
}
