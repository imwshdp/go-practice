package repositories

import (
	"context"
	"grpc-basics/apps/orders/internal/models"
	"grpc-basics/apps/orders/internal/storage"
)

type orderRepository struct {
	storage storage.OrderStorage
}

func NewOrderRepository(storage storage.OrderStorage) *orderRepository {
	return &orderRepository{storage}
}

func (r *orderRepository) CreateOrder(ctx context.Context, order *models.Order) error {
	r.storage.Set(ctx, *order)
	return nil
}
