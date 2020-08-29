package iam

import (
	"github.com/nokamoto/demo20-cli/internal/cmd/iam/admin"
	"github.com/spf13/cobra"
)

// RootCmd is a root of v1alpha sub commands.
var RootCmd = &cobra.Command{
	Use:   "iam",
	Short: "An iam service management tool",
}

func init() {
	RootCmd.AddCommand(admin.RootCmd)
}
