package app

import (
	"log/slog"
	"os"
)

func Setup(
	addr string,
) {
	logger := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}),
	)

	slog.SetDefault(logger)

	gRPCServer := NewGRPCServer(addr)
	logger.Info("gRPC server is started on", "addr", addr)

	gRPCServer.Run()
}
