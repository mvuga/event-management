package handlers

import "github.com/jackc/pgx/v5/pgxpool"

type EventHandler struct {
	dbPool *pgxpool.Pool
}

func NewEventHandler(db *pgxpool.Pool) *EventHandler {
	return &EventHandler{dbPool: db}
}
