package app

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"grpc-basics/apps/common/genproto/orders"
	grpcInterface "grpc-basics/apps/orders/internal/handlers/grpc"
	"grpc-basics/apps/orders/internal/services"
)

type gRPCServer struct {
	addr string
}

func NewGrpcServer(addr string) *gRPCServer {
	return &gRPCServer{addr}
}

func (s *gRPCServer) setup(
	grpcServer *grpc.Server,
	services *services.Services,
) {
	handlers := grpcInterface.NewHandlers(services)
	orders.RegisterOrderServiceServer(grpcServer, handlers.Order)
}

func (s *gRPCServer) Run(
	services *services.Services,
) error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	s.setup(grpcServer, services)

	return grpcServer.Serve(listener)
}
