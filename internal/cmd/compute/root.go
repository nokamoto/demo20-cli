package compute

import (
	"github.com/nokamoto/demo20-cli/internal/cmd/compute/instances"
	"github.com/nokamoto/demo20-cli/internal/template"
)

// RootCmd is a root of v1alpha sub commands.
var RootCmd = template.NewRoot("compute", "A compute service management tool")

func init() {
	RootCmd.AddCommand(instances.RootCmd)
}
