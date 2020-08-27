package config

import (
	"github.com/nokamoto/demo20-cli/internal/config"
	"github.com/spf13/cobra"
)

func newView(value *config.Value) *cobra.Command {
	return &cobra.Command{
		Use:           "view",
		Short:         "View current configration values",
		Args:          cobra.ExactArgs(0),
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			cmd.Print(value.View())
			return nil
		},
	}
}

func init() {
	RootCmd.AddCommand(newView(&config.Default))
}
