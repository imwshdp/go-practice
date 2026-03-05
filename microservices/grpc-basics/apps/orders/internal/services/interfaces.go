package services

import (
	"context"
	"grpc-basics/apps/orders/internal/repositories"
	"grpc-basics/apps/orders/internal/services/models"
)

type OrderService interface {
	CreateOrder(ctx context.Context, payload *models.Order) error
}

type Services struct {
	Order OrderService
}

func NewServices(
	repos *repositories.Repositories,
) *Services {
	return &Services{
		MewOrderService(repos.Order),
	}
}
