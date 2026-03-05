package app

import (
	httpInterface "grpc-basics/apps/orders/internal/handlers/http"
	"grpc-basics/apps/orders/internal/services"
	"net/http"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) setup(
	mux *http.ServeMux,
	services *services.Services,
) {
	handlers := httpInterface.NewHandlers(services)
	handlers.Order.RegisterRouter(mux)
}

func (s *httpServer) Run(
	services *services.Services,
) error {
	rootMux := http.NewServeMux()
	apiMux := http.NewServeMux()

	s.setup(apiMux, services)
	rootMux.Handle("/api/", http.StripPrefix("/api", apiMux))

	return http.ListenAndServe(s.addr, rootMux)
}
