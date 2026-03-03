package repositories

import (
	"context"
	"grpc-basics/apps/orders/internal/models"
	"grpc-basics/apps/orders/internal/storage"
)

// order
type OrderRepository interface {
	CreateOrder(ctx context.Context, order *models.Order) error
}

// Repositories
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
