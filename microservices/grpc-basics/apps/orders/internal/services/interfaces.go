package services

import (
	"context"
	"grpc-basics/apps/orders/internal/models"
	"grpc-basics/apps/orders/internal/repositories"
)

// order
type OrderService interface {
	CreateOrder(ctx context.Context, payload *models.Order) error
}

// Services
type Services struct {
	Order OrderService
}

func NewServices(
	repos *repositories.Repositories,
) *Services {
	return &Services{
		MewOrdersService(repos.Order),
	}
}
