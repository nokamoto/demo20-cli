package config

import (
	"fmt"
	"os"
	"path"

	"github.com/nokamoto/demo20-cli/internal/config"
	"github.com/spf13/cobra"
)

func newSet(value *config.Value) *cobra.Command {
	var (
		grpcAddress string
	)

	cmd := &cobra.Command{
		Use:           "set",
		Short:         "Set configration values to the default configration file",
		Args:          cobra.ExactArgs(0),
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if cmd.Flags().Changed("grpc-address") {
				value.GrpcAddress = grpcAddress
			}

			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			file := path.Join(home, fmt.Sprintf("%s.%s", config.Filename, config.Extension))
			err = value.Write(file)
			if err != nil {
				return err
			}

			cmd.Println("OK")

			return nil
		},
	}

	cmd.Flags().StringVar(&grpcAddress, "grpc-address", "", "gRPC server address")

	return cmd
}

func init() {
	RootCmd.AddCommand(newSet(&config.Default))
}
