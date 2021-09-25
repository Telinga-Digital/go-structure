package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Env   string
	Url   string
	Debug string
	Port  string
}

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file!")
	}
}

func GetAppConfig() *AppConfig {
	return &AppConfig{
		Env:   os.Getenv("APP_ENV"),
		Url:   os.Getenv("APP_URL"),
		Debug: os.Getenv("APP_DEBUG"),
		Port:  os.Getenv("APP_PORT"),
	}
}
