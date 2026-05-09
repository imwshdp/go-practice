package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Event struct {
	EventID   string    `db:"event_id"`
	EventName string    `db:"event_name"`
	CreatedAt time.Time `db:"created_at"`
}

func NewEvent() *Event {
	id := GenerateRandomString(15)
	return &Event{
		EventID:   id,
		EventName: "test_event",
		CreatedAt: time.Now(),
	}
}

type EventRepo struct {
	repo      *sqlx.DB
	tableName string
}

func NewEventRepo(db *sqlx.DB) *EventRepo {
	return &EventRepo{
		repo:      db,
		tableName: "app.events",
	}
}

func (r *EventRepo) Insert(ctx context.Context, tx *sqlx.Tx, e *Event) (string, error) {
	query := fmt.Sprintf(`
		INSERT INTO %s (event_id, event_name, created_at)
		VALUES (:event_id, :event_name, :created_at)
        ON CONFLICT (event_id) DO NOTHING
	`, r.tableName)

	_, err := tx.NamedExecContext(ctx, query, e)
	if err != nil {
		return "", fmt.Errorf("insert event: %w", err)
	}
	return e.EventID, nil
}

func (r *EventRepo) Get(ctx context.Context, tx *sqlx.Tx, eventID string) (*Event, error) {
	var event Event

	query := fmt.Sprintf(`
		SELECT event_id, event_name, created_at
		FROM %s
		WHERE event_id = $1
	`, r.tableName)

	err := tx.GetContext(ctx, &event, query, eventID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("get event: %w", err)
	}

	return &event, nil
}

func TxClosure[T any](ctx context.Context, r *EventRepo, fn func(ctx context.Context, tx *sqlx.Tx) (T, error)) (T, error) {
	tx, err := r.repo.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	})
	if err != nil {
		panic("unable to start TX")
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}

		if err != nil {
			tx.Rollback()
			return
		}

		err = tx.Commit()
		if err != nil {
			fmt.Printf("err on commit = %v\n", err)
		}
	}()

	res, err := fn(ctx, tx)
	return res, err
}
