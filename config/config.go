package config

type Config struct {
	AppHost     string `env:"APP_HOST"`
	AppPort     string `env:"APP_PORT"`
	DatabaseURL string `env:"DB_URL"`
	APIKey      string `env:"API_KEY"`
	ServerPort  int    `env:"SERVER_PORT"`
}
