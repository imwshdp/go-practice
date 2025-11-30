package orders

import (
	"context"
	repo "e-commerce-api/internal/adapters/postgresql/sqlc"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var (
	ErrProductNotFound = errors.New("product not found")
	ErrProductNotStock = errors.New("product has not enough stock")
)

type Service interface {
	PlaceOrder(ctx context.Context, payload createOrderParams) (repo.Order, error)
}

type service struct {
	db   *pgx.Conn
	repo *repo.Queries
}

func NewService(r *repo.Queries, db *pgx.Conn) Service {
	return &service{
		repo: r,
		db:   db,
	}
}

func (s *service) PlaceOrder(ctx context.Context, payload createOrderParams) (repo.Order, error) {
	if payload.CustomerID == 0 {
		return repo.Order{}, fmt.Errorf("customer ID is required")
	}

	if len(payload.Items) == 0 {
		return repo.Order{}, fmt.Errorf("at least minimum one order item is required")
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return repo.Order{}, err
	}
	defer tx.Rollback(ctx)

	qtx := s.repo.WithTx(tx)

	order, err := qtx.CreateOrder(ctx, payload.CustomerID)
	if err != nil {
		return repo.Order{}, err
	}

	for _, item := range payload.Items {
		product, err := qtx.FindProductByID(ctx, item.ProductID)
		if err != nil {
			return repo.Order{}, ErrProductNotFound
		}

		if product.Quantity < item.Quantity {
			return repo.Order{}, ErrProductNotStock
		}

		_, err = qtx.CreateOrderItem(ctx, repo.CreateOrderItemParams{
			OrderID:      order.ID,
			ProductID:    item.ProductID,
			Quantity:     item.Quantity,
			PriceInCents: product.PriceInCents,
		})

		if err != nil {
			return repo.Order{}, err
		}

		// update quantity of product
	}

	tx.Commit(ctx)

	return order, nil
}
