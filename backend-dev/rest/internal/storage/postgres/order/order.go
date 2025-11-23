package order

import (
	"database/sql"
	"fmt"
	"rest/internal/dto"
)

type OrderRepository interface {
	Create(order dto.Order) (int, error)
	CreateOrderItem(item dto.OrderItem) error
}

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (repo *orderRepository) Create(order dto.Order) (int, error) {
	const query = `
		INSERT INTO orders (user_id, total, status, address)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	var id int
	err := repo.db.QueryRow(
		query,
		order.UserID,
		order.Total,
		order.Status,
		order.Address,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("order creation: %w", err)
	}

	return int(id), nil
}

func (repo *orderRepository) CreateOrderItem(item dto.OrderItem) error {
	_, err := repo.db.Exec("INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)", item.OrderID, item.ProductID, item.Quantity, item.Price)

	if err != nil {
		return fmt.Errorf("order item creation: %s", err)
	}

	return nil
}
