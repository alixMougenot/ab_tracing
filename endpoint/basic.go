package endpoint

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/alixMougenot/ab_tracing/graph"
	"github.com/jackc/pgx/v5/pgxpool"
)

// HandlerFunc for the GraphQL endpoint
func MakeGraphQLHandler(pool *pgxpool.Pool) http.Handler {
	return handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DBPool: pool}}))
}
