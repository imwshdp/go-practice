package storage

import (
	"context"
	"grpc-basics/apps/orders/internal/services/models"
)

type OrderStorage interface {
	Get(ctx context.Context, orderID int) (*models.Order, error)
	Set(ctx context.Context, order models.Order)
	GetAll(ctx context.Context) []*models.Order
	GetLen(ctx context.Context) int
}
