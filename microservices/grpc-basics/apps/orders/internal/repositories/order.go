package repositories

import (
	"context"
	"grpc-basics/apps/orders/internal/services/models"
	"grpc-basics/apps/orders/internal/storage"
)

type orderRepository struct {
	storage storage.OrderStorage
}

func NewOrderRepository(storage storage.OrderStorage) *orderRepository {
	return &orderRepository{storage}
}

func (r *orderRepository) CreateOrder(ctx context.Context, order *models.Order) error {
	length := r.storage.GetLen(ctx)
	order.OrderID = length + 1

	r.storage.Set(ctx, *order)
	return nil
}

func (r *orderRepository) GetOrders(ctx context.Context) []*models.Order {
	return r.storage.GetAll(ctx)
}
