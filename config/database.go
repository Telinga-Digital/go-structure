package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Connection string `default:"mysql"`
	Host       string `default:"127.0.0.1"`
	Port       string `default:"3306"`
	Database   string `default:"go-structure"`
	Username   string `default:"root"`
	Password   string `default:""`
}

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file!")
	}
}

func GetDBConfig() *DBConfig {
	return &DBConfig{
		Connection: os.Getenv("DB_CONNECTION"),
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		Database:   os.Getenv("DB_DATABASE"),
		Username:   os.Getenv("DB_USERNAME"),
		Password:   os.Getenv("DB_PASSWORD"),
	}
}
