package app

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGrpcConnection(addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// func (s *gRPCServer) setup(
// 	grpcServer *grpc.Server,
// 	services *services.Services,
// ) {
// 	handlers := grpcInterface.NewHandlers(services)
// 	orders.RegisterOrderServiceServer(grpcServer, handlers.Order)
// }

// func (s *gRPCServer) Run(
// 	services *services.Services,
// ) error {
// 	listener, err := net.Listen("tcp", s.addr)
// 	if err != nil {
// 		return fmt.Errorf("failed to listen: %v", err)
// 	}

// 	grpcServer := grpc.NewServer()
// 	s.setup(grpcServer, services)

// 	return grpcServer.Serve(listener)
// }
