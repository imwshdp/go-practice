package main

import "grpc-basics/apps/kitchen/internal/app"

func main() {
	app.Setup(
		":6000",
		":7000",
	)
}
