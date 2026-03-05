package http

import (
	"grpc-basics/apps/common/genproto/orders"
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

func NewHandlers(
	client orders.OrderServiceClient,
) *Handlers {
	return &Handlers{
		NewOrderHandler(client),
	}
}
