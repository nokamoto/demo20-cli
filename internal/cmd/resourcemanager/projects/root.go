package projects

import (
	"github.com/nokamoto/demo20-cli/internal/cmd/resourcemanager/projects/rolebindings"
	"github.com/nokamoto/demo20-cli/internal/cmd/resourcemanager/projects/roles"
	"github.com/nokamoto/demo20-cli/internal/template"
)

// RootCmd is a root of v1alpha sub commands.
var RootCmd = template.NewRoot("projects", "A cloud resourcemanager projects management tool")

func init() {
	RootCmd.AddCommand(roles.RootCmd, rolebindings.RootCmd)
}
