package app

import (
	"grpc-basics/apps/orders/internal/repositories"
	"grpc-basics/apps/orders/internal/services"
	"grpc-basics/apps/orders/internal/storage"
	"log/slog"
	"os"
)

func setupLogger() *slog.Logger {
	logger := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}),
	)
	slog.SetDefault(logger)
	return logger
}

func setupStorage() storage.OrderStorage {
	return storage.NewOrderStorage()
}

func setupServices(
	db storage.OrderStorage,
) *services.Services {
	repos := repositories.NewRepositories(db)
	return services.NewServices(repos)
}

func Setup(
	grpcAddr string,
	httpAddr string,
) {
	logger := setupLogger()

	storage := setupStorage()
	services := setupServices(storage)

	httpServer := NewHttpServer(httpAddr)
	logger.Info("http server is started on", "addr", httpAddr)
	go httpServer.Run(services)

	gRPCServer := NewGrpcServer(grpcAddr)
	logger.Info("gRPC server is started on", "addr", grpcAddr)
	gRPCServer.Run(services)
}
