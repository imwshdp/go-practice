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
	order := models.CreateOrder(
		int(req.GetCustomerID()),
		int(req.GetProductID()),
		int(req.GetQuantity()),
	)

	err := h.service.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return &orders.CreateOrderResponse{
		Status: "success",
	}, nil
}

func (h *orderHandler) GetOrders(ctx context.Context, _ *orders.GetOrdersRequest) (*orders.GetOrderResponse, error) {
	data := h.service.GetOrders(ctx)
	dtoOrders := make([]*orders.Order, 0, len(data))

	for _, or := range data {
		dtoOrders = append(dtoOrders, &orders.Order{
			OrderID:    int32(or.OrderID),
			CustomerID: int32(or.CustomerID),
			ProductID:  int32(or.ProductID),
			Quantity:   int32(or.Quantity),
		})
	}

	return &orders.GetOrderResponse{
		Orders: dtoOrders,
	}, nil
}
