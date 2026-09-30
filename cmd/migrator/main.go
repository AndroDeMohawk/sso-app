package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/url" // Добавили стандартный пакет для работы с URL

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var dbURL, migrationsPath, migrationsTable string
	var down bool
	flag.StringVar(&dbURL, "db-url", "", "PostgreSQL connection URL")
	flag.StringVar(&migrationsPath, "migrations-path", "", "Path to migration files")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "Name of the table to track migrations")
	flag.BoolVar(&down, "down", false, "Rollback all migrations")
	flag.Parse()

	if dbURL == "" || migrationsPath == "" {
		log.Fatal("Error: --db-url and --migrations-path are required")
	}

	// ПРАВИЛЬНЫЙ СПОСОБ: внедряем x-migrations-table внутрь URL парсером Go
	u, err := url.Parse(dbURL)
	if err != nil {
		log.Fatalf("failed to parse db url: %v", err)
	}
	q := u.Query()
	q.Set("x-migrations-table", migrationsTable)
	u.RawQuery = q.Encode()

	// Инициализация мигратора
	m, err := migrate.New("file://"+migrationsPath, u.String())
	if err != nil {
		panic(err)
	}
	if down {
		if err := m.Down(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				fmt.Println("No migrations to rollback")
				return
			}
			panic(err)
		}
		fmt.Println("Rolled back migrations successfully")
		return
	}

	// Накатывание миграций
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("No migrations to apply")
			return
		}
		panic(err)
	}

	fmt.Println("Applied migrations successfully")
}
