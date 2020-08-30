package config

import (
	"github.com/nokamoto/demo20-cli/internal/config"
	"github.com/nokamoto/demo20-cli/internal/template"
	"github.com/spf13/cobra"
)

func newView(value *config.Value) *cobra.Command {
	return template.NewArg0("view", "View current configration values", func(cmd *cobra.Command) error {
		cmd.Print(value.View())
		return nil
	})
}

func init() {
	RootCmd.AddCommand(newView(&config.Default))
}
