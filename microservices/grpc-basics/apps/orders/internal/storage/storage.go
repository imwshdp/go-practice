package storage

import (
	"context"
	"errors"
	"grpc-basics/apps/orders/internal/services/models"
)

var ErrNotFound = errors.New("not found")

type storage struct {
	data map[int]models.Order
}

func NewOrderStorage() *storage {
	return &storage{
		data: map[int]models.Order{},
	}
}

func (s *storage) GetAll(ctx context.Context) []*models.Order {
	result := make([]*models.Order, 0, len(s.data))

	for _, order := range s.data {
		result = append(result, &order)
	}

	return result
}

func (s *storage) Get(ctx context.Context, orderID int) (*models.Order, error) {
	order, exists := s.data[orderID]
	if !exists {
		return nil, ErrNotFound
	}

	return &order, nil
}

func (s *storage) Set(ctx context.Context, order models.Order) {
	s.data[order.OrderID] = order
}

func (s *storage) GetLen(ctx context.Context) int {
	return len(s.data)
}
