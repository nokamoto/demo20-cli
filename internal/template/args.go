package template

import (
	"github.com/golang/protobuf/proto"
	"github.com/nokamoto/demo20-cli/internal/printer"
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

// NewArg0Proto returns a template sub command with ExactArgs(0).
func NewArg0Proto(use string, short string, run func(*cobra.Command) (proto.Message, error)) *cobra.Command {
	return NewArg0(use, short, func(cmd *cobra.Command) error {
		res, err := run(cmd)
		if err != nil {
			return err
		}
		return printer.Proto(cmd, res)
	})
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

// NewArg1Proto returns a template sub command with ExactArgs(1).
func NewArg1Proto(use string, short string, run func(*cobra.Command, string) (proto.Message, error)) *cobra.Command {
	return NewArg1(use, short, func(cmd *cobra.Command, arg string) error {
		res, err := run(cmd, arg)
		if err != nil {
			return err
		}
		return printer.Proto(cmd, res)
	})
}
