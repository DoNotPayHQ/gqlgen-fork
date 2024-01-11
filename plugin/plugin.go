// plugin package interfaces are EXPERIMENTAL.

package plugin

import (
	"github.com/DoNotPayHQ/gqlgen-fork/codegen"
	"github.com/DoNotPayHQ/gqlgen-fork/codegen/config"
)

type Plugin interface {
	Name() string
}

type ConfigMutator interface {
	MutateConfig(cfg *config.Config) error
}

type CodeGenerator interface {
	GenerateCode(cfg *codegen.Data) error
}
