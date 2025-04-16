package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/furre-dev/homelaan-go-backend/auth"
	"github.com/furre-dev/homelaan-go-backend/database"
	"github.com/furre-dev/homelaan-go-backend/graph"
	"github.com/furre-dev/homelaan-go-backend/internal"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/vektah/gqlparser/v2/ast"
)

func protectedDirective(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	_, ok := internal.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("not authorized")
	}
	// Add more checks if needed (role-based, etc)
	return next(ctx)
}

const defaultPort = "8000"

func main() {
	db := database.ConnectDB()
	defer db.Close(context.Background())

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	router.Use(auth.AuthMiddleware)

	srv := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{DB: db},
		Directives: graph.DirectiveRoot{
			Protected: protectedDirective,
		}}))

	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
