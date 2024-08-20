package config

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
}
