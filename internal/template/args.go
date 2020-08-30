package template

import (
	"github.com/spf13/cobra"
)

// NewArg0 returns a template sub command with ExactArgs(0).
func NewArg0(use string, short string, run func(*cobra.Command) error) *cobra.Command {
	return &cobra.Command{
		Use:           use,
		Short:         short,
		Long:          short + ".",
		Args:          cobra.ExactArgs(0),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd)
		},
	}
}

// NewArg1 returns a template sub command with ExactArgs(1).
func NewArg1(use string, short string, run func(*cobra.Command, string) error) *cobra.Command {
	return &cobra.Command{
		Use:           use,
		Short:         short,
		Long:          short + ".",
		Args:          cobra.ExactArgs(1),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd, args[0])
		},
	}
}
