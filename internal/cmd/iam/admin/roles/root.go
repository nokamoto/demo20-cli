package roles

import (
	"github.com/nokamoto/demo20-cli/internal/template"
)

// RootCmd is a root of v1alpha sub commands.
var RootCmd = template.NewRoot("roles", "A cloud iam role management tool")
