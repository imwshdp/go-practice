package services

import (
	"context"
	"grpc-basics/apps/orders/internal/models"
	"grpc-basics/apps/orders/internal/repositories"
)

type ordersService struct {
	repository repositories.OrderRepository
}

func MewOrdersService(
	repository repositories.OrderRepository,
) *ordersService {
	return &ordersService{repository}
}

func (s *ordersService) CreateOrder(ctx context.Context, payload *models.Order) error {
	err := s.repository.CreateOrder(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}
