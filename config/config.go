package config

import (
	"os"
)

// Config holds all application configuration.
type Config struct {
	Server ServerConfig
}

// ServerConfig holds server-related configuration.
type ServerConfig struct {
	Port string
}

// Load reads configuration from environment variables with sensible defaults.
func Load() (*Config, error) {
	cfg := &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
		},
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
