package rolebindings

import (
	"github.com/nokamoto/demo20-cli/internal/template"
)

// RootCmd is a root of v1alpha sub commands.
var RootCmd = template.NewRoot("rolebindings", "A cloud iam role binding management tool")
