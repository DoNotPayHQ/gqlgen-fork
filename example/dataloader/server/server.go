package main

import (
	"log"
	"net/http"

	"github.com/DoNotPayHQ/gqlgen-fork/example/dataloader"
	"github.com/DoNotPayHQ/gqlgen-fork/handler"
	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()
	router.Use(dataloader.LoaderMiddleware)

	router.Handle("/", handler.Playground("Dataloader", "/query"))
	router.Handle("/query", handler.GraphQL(
		dataloader.NewExecutableSchema(dataloader.Config{Resolvers: &dataloader.Resolver{}}),
	))

	log.Println("connect to http://localhost:8082/ for graphql playground")
	log.Fatal(http.ListenAndServe(":8082", router))
}
