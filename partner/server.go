package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jailtonjunior94/go-challenge/partner/graph"
	"github.com/jailtonjunior94/go-challenge/partner/graph/generated"
	"github.com/jailtonjunior94/go-challenge/partner/infrastructure/environments"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	environments.NewSetupEnvironments()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver()}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
