package config

import "os"

type Config struct {
	DBURL string
	Port  string
}

func LoadConfig() *Config {
	return &Config{
		DBURL: os.Getenv("DB_URL"),
		Port:  getEnv("PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
