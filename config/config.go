package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	defaultDBPort   = "5432"
	defautlHttpPort = "8080"
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

func (cfg *Config) BuildConnStr() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)
}
