package config

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "config",
	Short: "A commandline configration management tool",
}
