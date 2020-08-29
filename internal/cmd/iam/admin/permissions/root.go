package permissions

import (
	"github.com/spf13/cobra"
)

// RootCmd is a root of v1alpha sub commands.
var RootCmd = &cobra.Command{
	Use:   "permissions",
	Short: "A cloud iam permission management tool",
}
