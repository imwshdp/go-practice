package grpc

import (
	"context"
	"grpc-basics/apps/common/genproto/orders"
	"grpc-basics/apps/orders/internal/models"
	"grpc-basics/apps/orders/internal/services"
)

type ordersHandler struct {
	orders.UnimplementedOrderServiceServer
	service services.OrderService
}

func NewOrdersHandler(
	orderService services.OrderService,
) *ordersHandler {
	return &ordersHandler{
		service: orderService,
	}
}

func (h *ordersHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &models.Order{
		OrderID:    1,
		CustomerID: 2,
		ProductID:  3,
		Quantity:   4,
	}

	err := h.service.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return &orders.CreateOrderResponse{
		Status: "success",
	}, nil
}
