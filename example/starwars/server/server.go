package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/DoNotPayHQ/gqlgen-fork/example/starwars"
	"github.com/DoNotPayHQ/gqlgen-fork/example/starwars/generated"
	"github.com/DoNotPayHQ/gqlgen-fork/graphql"
	"github.com/DoNotPayHQ/gqlgen-fork/handler"
)

func main() {
	http.Handle("/", handler.Playground("Starwars", "/query"))
	http.Handle("/query", handler.GraphQL(generated.NewExecutableSchema(starwars.NewResolver()),
		handler.ResolverMiddleware(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
			rc := graphql.GetResolverContext(ctx)
			fmt.Println("Entered", rc.Object, rc.Field.Name)
			res, err = next(ctx)
			fmt.Println("Left", rc.Object, rc.Field.Name, "=>", res, err)
			return res, err
		}),
	))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
