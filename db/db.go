package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(ctx context.Context) (*pgxpool.Pool, error) {
	DatabaseURL := os.Getenv("DATABASE_URL")
	
	if DatabaseURL == "" {
		return nil, fmt.Errorf("Database URL Belum terisi di ENV")
	}

	pool, err := pgxpool.New(ctx, DatabaseURL)
	 
	if err != nil {
		return nil, err
	}

	// Cek ping

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	fmt.Println("Database Connected!")

	return pool, nil
}