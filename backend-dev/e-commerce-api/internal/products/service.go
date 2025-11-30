package products

import (
	"context"
	repo "e-commerce-api/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
}

type service struct {
	repo repo.Querier
}

func NewService(r repo.Querier) Service {
	return &service{
		repo: r,
	}
}

func (s *service) ListProducts(ctx context.Context) ([]repo.Product, error) {
	products, err := s.repo.ListProducts(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}
