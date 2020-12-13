package config

import (
  "os"
)

type Config struct {
  Environment string
	Port string
  ServiceName string
  Version string
}

func New() *Config {
	return &Config{
		Environment: GetEnv("ENV", ""),
		Port: GetEnv("PORT", ""),
    ServiceName: GetEnv("SERVICE_NAME", ""),
    Version: GetEnv("VERSION", ""),
	}
}

func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
