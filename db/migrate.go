package db

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func findMigrationsPath() string {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		testPath := filepath.Join(currentDir, "schema", "migrations")
		if _, err := os.Stat(testPath); err == nil {
			path := "file://" + testPath

			path = strings.Replace(path, "\\", "/", -1)
			return path
		}
		currentDir = filepath.Dir(currentDir)
	}

	log.Fatal("Migrations directory not found")
	return ""
}

func Migrate(db *sqlx.DB) {
	migrationsPath := findMigrationsPath()

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}
}
