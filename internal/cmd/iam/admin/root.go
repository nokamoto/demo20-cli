package admin

import (
	"github.com/nokamoto/demo20-cli/internal/cmd/iam/admin/auth"
	"github.com/nokamoto/demo20-cli/internal/cmd/iam/admin/machineusers"
	"github.com/nokamoto/demo20-cli/internal/cmd/iam/admin/permissions"
	"github.com/nokamoto/demo20-cli/internal/cmd/iam/admin/rolebindings"
	"github.com/nokamoto/demo20-cli/internal/cmd/iam/admin/roles"
	"github.com/nokamoto/demo20-cli/internal/template"
)

// RootCmd is a root of v1alpha sub commands.
var RootCmd = template.NewRoot("admin", "An iam service management tool for administration")

func init() {
	RootCmd.AddCommand(permissions.RootCmd, roles.RootCmd, rolebindings.RootCmd, machineusers.RootCmd, auth.RootCmd)
}
