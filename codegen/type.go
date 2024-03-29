package codegen

import (
	"github.com/DoNotPayHQ/gqlgen-fork/codegen/config"
)

func (b *builder) buildTypes() map[string]*config.TypeReference {
	ret := map[string]*config.TypeReference{}

	for _, ref := range b.Binder.References {
		for ref != nil {
			ret[ref.UniquenessKey()] = ref

			ref = ref.Elem()
		}
	}
	return ret
}
