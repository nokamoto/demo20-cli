package instances

import (
	"github.com/spf13/cobra"
)

// RootCmd is a root of v1alpha sub commands.
var RootCmd = &cobra.Command{
	Use:   "instances",
	Short: "A cloud compute instance management tool",
}
