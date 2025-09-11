package events

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (event *Event) CreateEvent(ctx context.Context, dbPool *pgxpool.Pool) error {
	stmt := `INSERT INTO events (name, description, location, date_time, user_id) 
	VALUES ($1, $2, $3, $4 ,$5) RETURNING id;`
	conn, err := dbPool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	result := conn.QueryRow(ctx, stmt, event.Name, event.Description, event.Location, event.DateTime, event.UserId)
	err = result.Scan(&event.ID)

	if err != nil {
		return err
	}
	return nil

}

func GetEvents(ctx context.Context, dbPool *pgxpool.Pool) ([]Event, error) {
	stmt := `SELECT * FROM events ORDER BY id ASC;`
	conn, err := dbPool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	result, err := conn.Query(ctx, stmt)

	if err != nil {
		return nil, err

	}
	defer result.Close()
	var events []Event
	for result.Next() {
		var event Event
		err = result.Scan(&event.ID, &event.Name, &event.Location, &event.Description, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil

}

func GetEventById(ctx context.Context, eventId int64, dbPool *pgxpool.Pool) (*Event, error) {
	stmt := `SELECT * FROM events WHERE id=$1`
	conn, err := dbPool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	result := conn.QueryRow(ctx, stmt, eventId)

	var event Event
	err = result.Scan(&event.ID, &event.Name, &event.Location, &event.Description, &event.DateTime, &event.UserId)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event Event) UpdateEvent(ctx context.Context, dbPool *pgxpool.Pool) error {
	stmt := `UPDATE events SET name = $1, location = $2, description = $3, date_time = $4 WHERE id=$5`
	conn, err := dbPool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(ctx, stmt, event.Name, event.Location, event.Description, event.DateTime, event.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEvent(ctx context.Context, eventId int64, dbPool *pgxpool.Pool) error {
	stmt := `DELETE FROM events WHERE id=$1`
	conn, err := dbPool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(ctx, stmt, eventId)

	if err != nil {
		return err
	}
	return nil
}
