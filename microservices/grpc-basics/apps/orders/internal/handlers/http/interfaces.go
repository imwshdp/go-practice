package http

import (
	"grpc-basics/apps/orders/internal/services"
	"net/http"
)

type BaseHandler interface {
	RegisterRouter(router *http.ServeMux)
}

type OrderHandler interface {
	BaseHandler
	CreateOrder(w http.ResponseWriter, r *http.Request)
}

type Handlers struct {
	Order OrderHandler
}

func NewHandlers(services *services.Services) *Handlers {
	return &Handlers{
		NewOrderHandler(services.Order),
	}
}
