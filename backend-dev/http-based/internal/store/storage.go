package store

import (
	"context"
	"database/sql"
)

type PostsRepository interface {
	Create(ctx context.Context, post *Post) error
}

type UsersRepository interface {
	Create(ctx context.Context, user *User) error
}

type Storage struct {
	Posts PostsRepository
	Users UsersRepository
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostsStore{db},
		Users: &UsersStore{db},
	}
}
