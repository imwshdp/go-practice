package repositories

import (
	"context"
	"grpc-basics/apps/orders/internal/services/models"
	"grpc-basics/apps/orders/internal/storage"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *models.Order) error
}

type Repositories struct {
	Order OrderRepository
}

func NewRepositories(
	db storage.OrderStorage,
) *Repositories {
	return &Repositories{
		NewOrderRepository(db),
	}
}
