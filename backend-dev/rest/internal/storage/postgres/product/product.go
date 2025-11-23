package product

import (
	"database/sql"
	"fmt"
	"rest/internal/dto"
	"strings"
)

type ProductRepository interface {
	GetAll() ([]*dto.Product, error)
	GetByIds(ids []int) ([]*dto.Product, error)
	Create(product dto.Product) error
	Update(product *dto.Product) error
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func scanRowsIntoProduct(rows *sql.Rows) (*dto.Product, error) {
	product := new(dto.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (repo *productRepository) GetAll() ([]*dto.Product, error) {
	rows, err := repo.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]*dto.Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func (repo *productRepository) GetByIds(productIDs []int) ([]*dto.Product, error) {
	if len(productIDs) == 0 {
		return []*dto.Product{}, nil
	}

	placeholders := make([]string, len(productIDs))
	args := make([]any, len(productIDs))

	for i, id := range productIDs {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = id
	}

	query := fmt.Sprintf(
		"SELECT id, name, description, image, price, quantity, created_at FROM products WHERE id IN (%s)",
		strings.Join(placeholders, ","),
	)

	rows, err := repo.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("products retrieval: %s", err)
	}

	products := []*dto.Product{}
	for rows.Next() {
		product, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, fmt.Errorf("products retrieval: %s", err)
		}

		products = append(products, product)
	}

	return products, nil
}

func (repo *productRepository) Create(product dto.Product) error {
	_, err := repo.db.Exec("INSERT INTO products (name, description, image, price, quantity) VALUES ($1, $2, $3, $4, $5)", product.Name, product.Description, product.Image, product.Price, product.Quantity)

	if err != nil {
		return fmt.Errorf("product creation: %s", err)
	}

	return nil
}

func (repo *productRepository) Update(product *dto.Product) error {
	_, err := repo.db.Exec("UPDATE products SET name = $1, price = $2, image = $3, description = $4, quantity = $5 WHERE id = $6", product.Name, product.Price, product.Image, product.Description, product.Quantity, product.ID)

	if err != nil {
		return fmt.Errorf("product update: %s", err)
	}

	return nil
}
