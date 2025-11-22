package main

import (
	"log"
	"rest/cmd/api"
	"rest/internal/config"
	"rest/internal/storage/postgres"
)

func main() {
	cfg := config.Load()

	db, err := postgres.New(
		cfg.DbConfig.Addr,
		cfg.DbConfig.MaxOpenConns,
		cfg.DbConfig.MaxIdleConns,
		cfg.DbConfig.MaxIdleTime,
	)

	defer func() {
		db.Close()
		log.Println("Database connection closed")
	}()
	log.Println("Database connection opened")

	if err != nil {
		log.Panic(err)
	}

	server := api.New(cfg.Addr, nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
