package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alixMougenot/ab_tracing/db"
	"github.com/alixMougenot/ab_tracing/endpoint"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// PG conf
	pgStr := "postgresql://postgres:postgress@localhost:5433/ab_tracing"
	pool, err := db.CreatePool(pgStr)
	if err != nil {
		log.Printf("Could not connect to database: %s", err.Error())
		os.Exit(35)
	}

	// Handling CORS
	allowdOrigins := []string{"http://localhost:8080", "http://localhost:5173", "192.168.85.133:5173"}
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: allowdOrigins,
	})

	srv := endpoint.MakeGraphQLHandler(pool)
	withCors := corsMiddleware.Handler(srv)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", withCors)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
