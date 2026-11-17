package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(databaseUrl string) (*sql.DB, error) {
	// Open is connection pool
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("sql.Open : %w", err)
	}
	// SetMaxOpenConns ,default =0 ,mean unlimited
	db.SetMaxOpenConns(25)
	// default 2 // 25-2 = 23 will be created and closed
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// fail fast -> pattern
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("db.ping : %w", err)
	}
	return db, nil
}
