package app

import (
	"grpc-basics/apps/common/genproto/orders"
	"grpc-basics/apps/kitchen/internal/handlers/http"
	"log"
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

func Setup(
	httpAddr string,
	grpcAddr string,
) {
	logger := setupLogger()

	grpcConn, err := NewGrpcConnection(grpcAddr)
	if err != nil {
		log.Fatal("gRPC client creation error: ", err)
	}

	grpcClient := orders.NewOrderServiceClient(grpcConn)
	handlers := http.NewHandlers(grpcClient)

	httpServer := NewHttpServer(httpAddr)
	logger.Info("http server is started on", "addr", httpAddr)

	httpServer.Run(handlers)
}
