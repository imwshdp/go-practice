package main

import (
	"grpc-basics/apps/common/fallbacks"
	"grpc-basics/apps/orders/internal/app"
	"os"
)

func main() {
	grpcPort, ok := os.LookupEnv("GRPC_PORT")
	if !ok {
		grpcPort = fallbacks.GrpcPort
	}

	ordersPort, ok := os.LookupEnv("ORDERS_PORT")
	if !ok {
		ordersPort = fallbacks.OrdersPort
	}

	app.Setup(
		":"+ordersPort,
		":"+grpcPort,
	)
}
