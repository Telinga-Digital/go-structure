package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Connection string
	Host       string
	Port       string
	Database   string
	Username   string
	Password   string
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
