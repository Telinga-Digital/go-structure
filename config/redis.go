package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type RedisConfig struct {
	Enabled  string `default:"false"`
	Host     string `default:"127.0.0.1"`
	Port     string `default:"6379"`
	Password string `default:""`
	DB       string `default:"0"`
}

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetRedisConfig() *RedisConfig {
	return &RedisConfig{
		Enabled:  os.Getenv("REDIS_ENABLED"),
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       os.Getenv("REDIS_DB"),
	}
}
