package main

import (
	"context"
	"e-commerce-api/internal/env"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DB_STRING", "user=user password=user_password host=127.0.0.1 port=5435 dbname=ecom sslmode=disable"),
		},
	}

	logger := slog.New(
		slog.NewTextHandler(os.Stdout, nil),
	)
	slog.SetDefault(logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		slog.Error("Database connection failed", "error", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	api := application{
		config: cfg,
		db:     conn,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
