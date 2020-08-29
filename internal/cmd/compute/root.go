package compute

import (
	"github.com/nokamoto/demo20-cli/internal/cmd/compute/instances"
	"github.com/spf13/cobra"
)

// RootCmd is a root of v1alpha sub commands.
var RootCmd = &cobra.Command{
	Use:   "compute",
	Short: "A compute service management tool",
}

func init() {
	RootCmd.AddCommand(instances.RootCmd)
}
