package testserver

import (
	"context"
	"testing"

	"github.com/DoNotPayHQ/gqlgen-fork/client"
	"github.com/DoNotPayHQ/gqlgen-fork/handler"
	"github.com/stretchr/testify/require"
)

func TestTypeFallback(t *testing.T) {
	resolvers := &Stub{}

	c := client.New(handler.GraphQL(NewExecutableSchema(Config{Resolvers: resolvers})))

	resolvers.QueryResolver.Fallback = func(ctx context.Context, arg FallbackToStringEncoding) (FallbackToStringEncoding, error) {
		return arg, nil
	}

	t.Run("fallback to string passthrough", func(t *testing.T) {
		var resp struct {
			Fallback string
		}
		c.MustPost(`query { fallback(arg: A) }`, &resp)
		require.Equal(t, "A", resp.Fallback)
	})
}
