package app

import (
	httpInterface "grpc-basics/apps/kitchen/internal/handlers/http"
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
	handlers *httpInterface.Handlers,
) {
	handlers.Order.RegisterRouter(mux)
}

func (s *httpServer) Run(
	handlers *httpInterface.Handlers,
) error {
	rootMux := http.NewServeMux()
	apiMux := http.NewServeMux()

	s.setup(apiMux, handlers)
	rootMux.Handle("/api/", http.StripPrefix("/api", apiMux))

	return http.ListenAndServe(s.addr, rootMux)
}
