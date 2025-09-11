package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

/* const defaultMaxConns = 10
const defaultMinConns = 0
const defaultMaxConnLifetime = time.Hour * 1
const defaultMaxConnIdleTime = time.Minute * 30
const defaultHealthCheckPeriod = time.Minute
const defaultConnectTimeout = time.Second * 5 */

func CreateConnectionPool(ctx context.Context, dbUser string, dbPassword string, dbHost string, dbPort int, dbName string) (*pgxpool.Pool, error) {

	databaseUrl := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	dbPool, err := pgxpool.New(ctx, databaseUrl)
	if err != nil {
		return nil, err
	}
	return dbPool, nil
}
