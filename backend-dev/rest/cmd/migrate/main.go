package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"rest/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	log.Println("Starting migration...")

	cfg := config.Load()

	db, err := sql.Open("postgres", cfg.DbConfig.Addr)
	if err != nil {
		log.Fatal(fmt.Errorf("Postgres connection establish error: %s", err))
	}

	defer func() {
		db.Close()
		log.Println("Database connection closed after migration")
	}()

	if err := db.Ping(); err != nil {
		log.Fatal(fmt.Errorf("Database ping failed: %s", err))
	}

	log.Println("Database connection opened for migration")

	m, err := migrate.New(
		"file://cmd/migrate/migrations",
		cfg.DbConfig.Addr,
	)

	if err != nil {
		log.Fatal(fmt.Errorf("Migration init failed: %s", err))
	}

	cmd := os.Args[(len(os.Args) - 1)]

	switch cmd {
	case "up":
		if err := m.Up(); err != nil {
			log.Fatal("Migration up failed", err)
		}

	case "down":
		if err := m.Down(); err != nil {
			log.Fatal("Migration down failed", err)
		}

	default:
		log.Fatal("Unknown command")
	}

	log.Println("Migration completed successfully")
}
