package events

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateRegistration(ctx context.Context, eventId, userId int64, dbPool *pgxpool.Pool) error {
	stmt := `INSERT INTO registrations (event_id, user_id) 
	VALUES ($1, $2);`
	conn, err := dbPool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(ctx, stmt, eventId, userId)

	if err != nil {
		return err
	}
	return nil

}

func CancelRegistration(ctx context.Context, eventId, userId int64, dbPool *pgxpool.Pool) error {
	stmt := `DELETE FROM registrations WHERE event_id=$1 AND user_id=$2`
	conn, err := dbPool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	result, err := conn.Exec(context.Background(), stmt, eventId, userId)
	if result.RowsAffected() == 0 {
		return errors.New("no rows matching that event id and user id")
	}
	if err != nil {
		return err
	}
	return nil

}
