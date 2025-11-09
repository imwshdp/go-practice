package repositories

import (
	"context"
	"database/sql"
	"e-commerce/internal/dtos"
)

type UsersRepository interface {
	Create(ctx context.Context, user *dtos.User) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UsersRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Create(ctx context.Context, user *dtos.User) error {
	return nil
}
