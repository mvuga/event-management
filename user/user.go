package user

import (
	"context"
	"errors"

	"rest-api/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

func (usr User) Create(ctx context.Context, dbPool *pgxpool.Pool) error {
	stmt := `
	INSERT INTO users (email, password) VALUES ($1, $2)
	`
	conn, err := dbPool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	hashPassword, err := utils.HashValue(usr.Password)
	if err != nil {
		return err
	}
	_, err = conn.Exec(ctx, stmt, usr.Email, hashPassword)

	if err != nil {
		return err
	}

	return nil
}

func (usr *User) ValidateCredentials(ctx context.Context, dbPool *pgxpool.Pool) error {

	stmt := "SELECT id, password from users where email = $1"
	conn, err := dbPool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	result := conn.QueryRow(ctx, stmt, usr.Email)

	var password string
	err = result.Scan(&usr.ID, &password)
	if err != nil {
		return err
	}

	if utils.CheckPasswordHash(usr.Password, password) != nil {
		return errors.New("invalid login credentials")
	}

	return nil
}
