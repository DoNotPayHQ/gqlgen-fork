package config

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vektah/gqlparser"
	"github.com/vektah/gqlparser/ast"
)

func TestSlicePointerBinding(t *testing.T) {
	t.Run("without OmitSliceElementPointers", func(t *testing.T) {
		binder, schema := createBinder(Config{
			OmitSliceElementPointers: false,
		})

		ta, err := binder.TypeReference(schema.Query.Fields.ForName("messages").Type, nil)
		if err != nil {
			panic(err)
		}

		require.Equal(t, ta.GO.String(), "[]*github.com/DoNotPayHQ/gqlgen-fork/example/chat.Message")
	})

	t.Run("with OmitSliceElementPointers", func(t *testing.T) {
		binder, schema := createBinder(Config{
			OmitSliceElementPointers: true,
		})

		ta, err := binder.TypeReference(schema.Query.Fields.ForName("messages").Type, nil)
		if err != nil {
			panic(err)
		}

		require.Equal(t, ta.GO.String(), "[]github.com/DoNotPayHQ/gqlgen-fork/example/chat.Message")
	})
}

func createBinder(cfg Config) (*Binder, *ast.Schema) {
	cfg.Models = TypeMap{
		"Message": TypeMapEntry{
			Model: []string{"github.com/DoNotPayHQ/gqlgen-fork/example/chat.Message"},
		},
	}

	s := gqlparser.MustLoadSchema(&ast.Source{Name: "TestAutobinding.schema", Input: `
		type Message { id: ID }

		type Query {
			messages: [Message!]!
		}
	`})

	b, err := cfg.NewBinder(s)
	if err != nil {
		panic(err)
	}

	return b, s
}
