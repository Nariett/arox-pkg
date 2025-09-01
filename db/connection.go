package db

import (
	"fmt"
	"github.com/Nariett/arox-pkg/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgres(cfg *config.Config) (*sqlx.DB, error) {
	err := CreateDatabase(cfg)
	if err != nil {
		return nil, err
	}

	connStr := cfg.BuildConnStr()

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error while connecting to postgres: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping the database: %w", err)

	}

	return db, nil
}
