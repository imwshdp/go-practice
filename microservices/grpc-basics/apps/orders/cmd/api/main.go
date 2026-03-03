package main

import "grpc-basics/apps/orders/internal/app"

func main() {
	app.Setup(
		":9000",
	)
}
