package database

import (
	"context"
	"fmt"

	"github.com/Oik17/gRPC-game/config"
	db "github.com/Oik17/gRPC-game/db/sqlc"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB(ctx context.Context) (*db.Queries, error) {

	dsn := config.Config("DSN")
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}
	fmt.Println("Database connected")
	sqlDB := stdlib.OpenDBFromPool(pool)

	queries := db.New(sqlDB)

	return queries, nil
}
