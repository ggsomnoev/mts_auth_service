package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Env            string
	Port           string
	WebAPICertFile string
	WebAPIKeyFile  string
}

func Load() *Config {
	return &Config{
		Env:            envOrDefault("API_ENV", "local"),
		Port:           envOrDefault("API_PORT", "8080"),
		WebAPICertFile: envOrDefault("WEB_API_CERT_FILE", "certs/tls.crt"),
		WebAPIKeyFile:  envOrDefault("WEB_API_KEY_FILE", "certs/tls.key"),
	}
}

func envOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
