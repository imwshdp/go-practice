package orders

import (
	"grpc-basics/apps/common/genproto/orders"
	"grpc-basics/apps/orders/internal/services"
)

type ordersHandler struct{
	service services.OrderService
	// TODO: maybe this interface must be in repo
	orders.UnimplementedOrderServiceServer
}

func NewOrdersHandler() *ordersHandler {
	return &ordersHandler{}
}
