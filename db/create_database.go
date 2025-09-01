package db

import (
	"fmt"
	"github.com/Nariett/arox-pkg/config"
	"github.com/jmoiron/sqlx"
	"log"
)

func CreateDatabase(cfg *config.Config) error {
	dbname := cfg.DBName

	strConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=%s", cfg.Host, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.SSLMode)

	db, err := sqlx.Open("postgres", strConn)
	if err != nil {
		return fmt.Errorf("failed to connect to PostgreSQL server: %w", err)
	}
	defer func(db *sqlx.DB) error {
		err := db.Close()
		if err != nil {
			return fmt.Errorf("failed to close DB connection: %w", err)
		}
		return nil
	}(db)

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping PostgreSQL server: %w", err)
	}

	var exists bool
	err = db.Get(&exists, `
		SELECT EXISTS(
			SELECT 1 FROM pg_database WHERE datname = $1
		)
	`, dbname)
	if err != nil {
		return fmt.Errorf("failed to check database existence: %w", err)
	}

	if !exists {
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbname))
		if err != nil {
			return fmt.Errorf("failed to create database: %w", err)
		}

		log.Println("Database created successfully")
	}

	return nil
}
