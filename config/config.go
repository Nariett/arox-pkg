package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func New() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file %v", err)

		return nil, err
	}

	return &Config{
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		DBName:   os.Getenv("DBNAME"),
		SSLMode:  os.Getenv("SSLMODE"),
		Port:     os.Getenv("PORT"),
		Host:     os.Getenv("HOST"),
	}, nil

}
