package products

import "context"

type Service interface {
	ListProducts(ctx context.Context) ([]string, error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) ListProducts(ctx context.Context) ([]string, error) {
	return []string{"Product 1", "Product 2", "Product 3"}, nil
}
