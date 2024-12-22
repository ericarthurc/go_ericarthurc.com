package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DbPool struct {
	*pgxpool.Pool
}

func NewDbPool() (*DbPool, error) {
	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return &DbPool{dbPool}, nil
}
