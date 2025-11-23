package user

import (
	"database/sql"
	"fmt"
	"rest/internal/dto"
)

type UserRepository interface {
	GetByEmail(email string) (*dto.User, error)
	GetById(id int) (*dto.User, error)
	Create(user dto.User) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func scanRowsIntoUser(rows *sql.Rows) (*dto.User, error) {
	user := new(dto.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *userRepository) GetById(id int) (*dto.User, error) {
	rows, err := repo.db.Query("SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	user := new(dto.User)

	for rows.Next() {
		user, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (repo *userRepository) GetByEmail(email string) (*dto.User, error) {
	rows, err := repo.db.Query("SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}

	user := new(dto.User)

	for rows.Next() {
		user, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (repo *userRepository) Create(user dto.User) error {
	_, err := repo.db.Exec("INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4)", user.FirstName, user.LastName, user.Email, user.Password)

	if err != nil {
		return fmt.Errorf("user creation: %s", err)
	}

	return nil
}
