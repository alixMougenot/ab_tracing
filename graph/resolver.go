package graph

import "github.com/jackc/pgx/v5/pgxpool"

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DBPool *pgxpool.Pool
}
