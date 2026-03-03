package app

import (
	"grpc-basics/apps/orders/internal/storage"
	"log/slog"
	"os"
)

func setupStorage() storage.OrderStorage {
	return storage.NewOrderStorage()
}

func Setup(
	addr string,
) {
	logger := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}),
	)
	slog.SetDefault(logger)

	storage := setupStorage()
	gRPCServer := NewGrpcServer(addr)

	logger.Info("gRPC server is started on", "addr", addr)

	gRPCServer.Run(storage)
}
