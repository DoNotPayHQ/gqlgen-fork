package main

import (
	"log"
	"net/http"

	"github.com/DoNotPayHQ/gqlgen-fork/example/scalars"
	"github.com/DoNotPayHQ/gqlgen-fork/handler"
)

func main() {
	http.Handle("/", handler.Playground("Starwars", "/query"))
	http.Handle("/query", handler.GraphQL(scalars.NewExecutableSchema(scalars.Config{Resolvers: &scalars.Resolver{}})))

	log.Fatal(http.ListenAndServe(":8084", nil))
}
