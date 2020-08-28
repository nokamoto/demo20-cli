package compute

import (
	"context"

	"github.com/nokamoto/demo20-apis/cloud/compute/v1alpha"
	"github.com/nokamoto/demo20-cli/internal/client"
	"github.com/nokamoto/demo20-cli/internal/config"
	"github.com/nokamoto/demo20-cli/internal/printer"
	"github.com/spf13/cobra"
)

func newCreate(value *config.Value, client *client.Client) *cobra.Command {
	var (
		labels []string
	)

	cmd := &cobra.Command{
		Use:           "create",
		Short:         "Create a new instance",
		Args:          cobra.ExactArgs(0),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			res, err := client.Compute().CreateInstance(context.Background(), &v1alpha.CreateInstanceRequest{
				Instance: &v1alpha.Instance{
					Labels: labels,
				},
			})
			if err != nil {
				return err
			}
			return printer.Proto(cmd, res)
		},
	}

	cmd.Flags().StringArrayVar(&labels, "labels", nil, "labels")

	return cmd
}

func init() {
	RootCmd.AddCommand(newCreate(&config.Default, client.NewClient(&config.Default)))
}
