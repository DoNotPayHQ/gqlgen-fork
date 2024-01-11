package main

import (
	"log"
	"net/http"

	"github.com/DoNotPayHQ/gqlgen-fork/example/selection"
	"github.com/DoNotPayHQ/gqlgen-fork/handler"
)

func main() {
	http.Handle("/", handler.Playground("Selection Demo", "/query"))
	http.Handle("/query", handler.GraphQL(selection.NewExecutableSchema(selection.Config{Resolvers: &selection.Resolver{}})))
	log.Fatal(http.ListenAndServe(":8086", nil))
}
