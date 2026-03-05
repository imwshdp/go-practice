package main

import (
	"grpc-basics/apps/common/fallbacks"
	"grpc-basics/apps/kitchen/internal/app"
	"os"
)

func main() {
	grpcPort, ok := os.LookupEnv("GRPC_PORT")
	if !ok {
		grpcPort = fallbacks.GrpcPort
	}

	kitchenPort, ok := os.LookupEnv("ORDERS_PORT")
	if !ok {
		kitchenPort = fallbacks.KitchenPort
	}

	app.Setup(
		":"+kitchenPort,
		":"+grpcPort,
	)
}
