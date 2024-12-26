package config

type Config struct {
	App   App
	DB    DB
	Mongo Mongo
}

type App struct {
	Url string `env:"APP_URL" env-default:"localhost:50060"`
}

type DB struct {
	Dsn string `env:"DB_DSN" env-required:"true"`
}

type Mongo struct {
	URI      string `env:"MONGO_URI" env-required:"true"`
	Database string `env:"MONGO_DATABASE" env-required:"true"`
}
