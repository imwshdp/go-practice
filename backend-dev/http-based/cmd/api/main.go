package main

import (
	"http-server/internal/database"
	"http-server/internal/env"
	"http-server/internal/store"
	"log"
	"time"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://user:password@localhost:5433/go_backend_dev?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetDuration("DB_MAX_IDLE_TIME", time.Minute*15),
		},
	}

	db, err := database.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	defer func() {
		db.Close()
		log.Println("Database connection closed")
	}()
	log.Println("Database connection opened")

	if err != nil {
		log.Panic(err)
	}

	storage := store.NewStorage(db)

	app := &application{
		config:  cfg,
		storage: storage,
	}

	router := app.mount()
	log.Fatal(app.run(router))
}
