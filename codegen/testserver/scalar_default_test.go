package testserver

import (
	"context"
	"testing"

	"github.com/DoNotPayHQ/gqlgen-fork/client"
	"github.com/DoNotPayHQ/gqlgen-fork/handler"
	"github.com/stretchr/testify/require"
)

func TestDefaultScalarImplementation(t *testing.T) {
	resolvers := &Stub{}

	c := client.New(handler.GraphQL(NewExecutableSchema(Config{Resolvers: resolvers})))

	resolvers.QueryResolver.DefaultScalar = func(ctx context.Context, arg string) (i string, e error) {
		return arg, nil
	}

	t.Run("with arg value", func(t *testing.T) {
		var resp struct{ DefaultScalar string }
		c.MustPost(`query { defaultScalar(arg: "fff") }`, &resp)
		require.Equal(t, "fff", resp.DefaultScalar)
	})

	t.Run("with default value", func(t *testing.T) {
		var resp struct{ DefaultScalar string }
		c.MustPost(`query { defaultScalar  }`, &resp)
		require.Equal(t, "default", resp.DefaultScalar)
	})
}
