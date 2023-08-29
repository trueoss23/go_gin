package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

var Cfg Config

func init() {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if err := env.Parse(&Cfg); err != nil {
		panic(err)
	}
}

type Config struct {
	AppHost string `env:"APP_HOST"`
	AppPort string `env:"APP_PORT"`
	DbName  string `env:"DB_NAME"`
	DbUser  string `env:"DB_USER"`
	DbPass  string `env:"DB_PASS"`
}
