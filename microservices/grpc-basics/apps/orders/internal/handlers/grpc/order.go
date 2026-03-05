package grpc

import (
	"context"
	"grpc-basics/apps/common/genproto/orders"
	"grpc-basics/apps/orders/internal/services"
	"grpc-basics/apps/orders/internal/services/models"
)

type orderHandler struct {
	orders.UnimplementedOrderServiceServer
	service services.OrderService
}

func NewOrderHandler(
	orderService services.OrderService,
) *orderHandler {
	return &orderHandler{
		service: orderService,
	}
}

func (h *orderHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := models.CreateOrder(1, 2, 3, 4)

	err := h.service.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return &orders.CreateOrderResponse{
		Status: "success",
	}, nil
}
