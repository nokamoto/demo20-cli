package admin

import (
	"github.com/nokamoto/demo20-cli/internal/cmd/iam/admin/permissions"
	"github.com/spf13/cobra"
)

// RootCmd is a root of v1alpha sub commands.
var RootCmd = &cobra.Command{
	Use:   "admin",
	Short: "An iam service management tool for administration",
}

func init() {
	RootCmd.AddCommand(permissions.RootCmd)
}
