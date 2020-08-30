package resourcemanager

import (
	"github.com/nokamoto/demo20-cli/internal/cmd/resourcemanager/projects"
	"github.com/nokamoto/demo20-cli/internal/template"
)

// RootCmd is a root of v1alpha sub commands.
var RootCmd = template.NewRoot("resourcemanager", "A resourcemanager service management tool")

func init() {
	RootCmd.AddCommand(projects.RootCmd)
}
