package config

import (
	"github.com/spf13/cobra"
)

// RootCmd is a root of config sub commands.
var RootCmd = &cobra.Command{
	Use:   "config",
	Short: "A commandline configration management tool",
}
