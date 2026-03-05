package services

import (
	"context"
	"grpc-basics/apps/orders/internal/repositories"
	"grpc-basics/apps/orders/internal/services/models"
)

type orderService struct {
	repository repositories.OrderRepository
}

func MewOrderService(
	repository repositories.OrderRepository,
) *orderService {
	return &orderService{repository}
}

func (s *orderService) CreateOrder(ctx context.Context, payload *models.Order) error {
	err := s.repository.CreateOrder(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}
