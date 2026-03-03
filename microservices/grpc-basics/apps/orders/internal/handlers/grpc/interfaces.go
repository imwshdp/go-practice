package grpc

import (
	"context"
	"grpc-basics/apps/common/genproto/orders"
	"grpc-basics/apps/orders/internal/services"
)

// order
type OrderHandler interface {
	orders.UnimplementedOrderServiceServer
	CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error)
}

// Handlers
type Handlers struct {
	Order orders.OrderServiceServer
}

func NewHandlers(services *services.Services) *Handlers {
	return &Handlers{
		NewOrdersHandler(services.Order),
	}
}
