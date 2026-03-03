package services

import (
	"context"
	"grpc-basics/apps/orders/internal/services/orders"
)

type OrderService interface {
	CreateOrder(ctx context.Context, payload orders.Order) error
}
