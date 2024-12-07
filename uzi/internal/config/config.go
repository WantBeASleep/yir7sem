package config

type Config struct {
	App    App
	DB     DB
	S3     S3
	Broker Broker
}

type App struct {
	Url string `env:"APP_URL" env-default:"localhost:50060"`
}

type DB struct {
	Dsn string `env:"DB_DSN" env-required:"true"`
}

type S3 struct {
	Endpoint     string `env:"S3_ENDPOINT" env-required:"true"`
	Access_Token string `env:"S3_TOKEN_ACCESS" env-required:"true"`
	Secret_Token string `env:"S3_TOKEN_SECRET" env-required:"true"`
}

type Broker struct {
	Addrs []string `env:"BROKER_ADDRS" env-required:"true"`
}
