package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	HttpPort string
	Postgres ConfigPostgres
}

type ConfigPostgres struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func Load(path string) Config {
	godotenv.Load(path + "/.env")
	cfg := viper.New()
	cfg.AutomaticEnv()
	config := Config{
		HttpPort: cfg.GetString("HTTP_PORT"),
		Postgres: ConfigPostgres{
			Host:     cfg.GetString("POSTGRES_HOST"),
			Port:     cfg.GetString("POSTGRES_PORT"),
			User:     cfg.GetString("POSTGRES_USER"),
			Password: cfg.GetString("POSTGRES_PASSWORD"),
			Database: cfg.GetString("POSTGRES_DATABASE"),
		},
	}

	return config
}
