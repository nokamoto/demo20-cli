package iam

import (
	"github.com/nokamoto/demo20-cli/internal/cmd/iam/admin"
	"github.com/nokamoto/demo20-cli/internal/template"
)

// RootCmd is a root of v1alpha sub commands.
var RootCmd = template.NewRoot("iam", "An iam service management tool")

func init() {
	RootCmd.AddCommand(admin.RootCmd)
}
