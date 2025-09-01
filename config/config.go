package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	defaultDBPort   = "5432"
	defaultHttpPort = "8080"
)

func New() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file %v", err)

		return nil, err
	}

	return &Config{
		DBUser:       os.Getenv("DB_USER"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		DBName:       os.Getenv("DB_NAME"),
		SSLMode:      os.Getenv("DB_SSLMODE"),
		DBPort:       os.Getenv("DB_PORT"),
		Port:         os.Getenv("PORT"),
		Host:         os.Getenv("HOST"),
		Protocol:     os.Getenv("PROTOCOL"),
		LPort:        os.Getenv("LPORT"),
		ProductsPort: os.Getenv("PRODUCTS_PORT"),
	}, nil
}

func (cfg *Config) BuildConnStr() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.SSLMode)
}
