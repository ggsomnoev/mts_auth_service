package config

import (
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Port             string
	WebAPICertFile   string
	WebAPIKeyFile    string
	CACertFile       string
	TrustedClientCNs []string
}

func Load() *Config {
	return &Config{
		Port:             envOrDefault("API_PORT", "8443"),
		WebAPICertFile:   envOrDefault("WEB_API_CERT_FILE", "certs/server/server.crt"),
		WebAPIKeyFile:    envOrDefault("WEB_API_KEY_FILE", "certs/server/server.key"),
		CACertFile:       envOrDefault("CA_CERT_FILE", "certs/ca/ca.crt"),
		TrustedClientCNs: parseCSVEnvOrDefault("TRUSTED_CLIENT_CNS", []string{"authorized-client"}),
	}
}

func envOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func parseCSVEnvOrDefault(key string, defaultVals []string) []string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVals
	}
	parts := strings.Split(val, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}
