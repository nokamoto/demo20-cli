package template

import (
	"github.com/spf13/cobra"
)

// NewRoot returns a template root command.
func NewRoot(use string, short string) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		Long:  short + ".",
	}
}
