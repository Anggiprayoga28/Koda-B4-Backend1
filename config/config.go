package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port             string
	ShowPasswordHash bool
}

var AppConfig *Config

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using defaults")
	}

	config := &Config{
		Port:             getEnv("APP_PORT", "8080"),
		ShowPasswordHash: getEnv("SHOW_PASSWORD_HASH", "true") == "true",
	}

	AppConfig = config
	return config
}

func getEnv(key, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}
