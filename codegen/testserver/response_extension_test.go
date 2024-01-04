package testserver

import (
	"context"
	"testing"

	"github.com/DoNotPayHQ/gqlgen-fork/client"
	"github.com/DoNotPayHQ/gqlgen-fork/graphql"
	"github.com/DoNotPayHQ/gqlgen-fork/handler"
	"github.com/stretchr/testify/require"
)

func TestResponseExtension(t *testing.T) {
	resolvers := &Stub{}
	resolvers.QueryResolver.Valid = func(ctx context.Context) (s string, e error) {
		return "Ok", nil
	}

	srv := handler.GraphQL(
		NewExecutableSchema(Config{Resolvers: resolvers}),
		handler.RequestMiddleware(func(ctx context.Context, next func(ctx context.Context) []byte) []byte {
			rctx := graphql.GetRequestContext(ctx)
			if err := rctx.RegisterExtension("example", "value"); err != nil {
				panic(err)
			}
			return next(ctx)
		}),
	)
	c := client.New(srv)

	raw, _ := c.RawPost(`query { valid }`)
	require.Equal(t, raw.Extensions["example"], "value")
}
