package admin

import (
	"github.com/nokamoto/demo20-cli/internal/cmd/iam/admin/permissions"
	"github.com/nokamoto/demo20-cli/internal/template"
)

// RootCmd is a root of v1alpha sub commands.
var RootCmd = template.NewRoot("admin", "An iam service management tool for administration")

func init() {
	RootCmd.AddCommand(permissions.RootCmd)
}
