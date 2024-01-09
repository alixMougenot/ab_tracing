// Copyright (c) 2023, Alix Mougenot (Au pied des Arbres), alix.mougenot@gmail.com
// See license file.
package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreatePool(connStr string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	log.Println("DB is setup")
	return pool, nil
}
