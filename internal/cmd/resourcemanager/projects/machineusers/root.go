package machineusers

import (
	"github.com/nokamoto/demo20-cli/internal/template"
)

// RootCmd is a root of v1alpha sub commands.
var RootCmd = template.NewRoot("machineusers", "A cloud iam project machine user management tool")
