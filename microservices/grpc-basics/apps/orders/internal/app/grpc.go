package app

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"grpc-basics/apps/common/genproto/orders"
	grpcInterface "grpc-basics/apps/orders/internal/handlers/grpc"
	"grpc-basics/apps/orders/internal/repositories"
	"grpc-basics/apps/orders/internal/services"
	"grpc-basics/apps/orders/internal/storage"
)

type gRPCServer struct {
	addr string
}

func NewGrpcServer(addr string) *gRPCServer {
	return &gRPCServer{addr}
}

func (s *gRPCServer) setup(
	grpcServer *grpc.Server,
	db storage.OrderStorage,
) {
	repos := repositories.NewRepositories(db)
	services := services.NewServices(repos)
	handlers := grpcInterface.NewHandlers(services)

	orders.RegisterOrderServiceServer(grpcServer, handlers.Order)
}

func (s *gRPCServer) Run(
	db storage.OrderStorage,
) error {

	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	s.setup(grpcServer, db)

	return grpcServer.Serve(listener)
}
