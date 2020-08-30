package instances

import (
	"context"

	"github.com/nokamoto/demo20-cli/internal/template"

	"github.com/nokamoto/demo20-apis/cloud/compute/v1alpha"
	"github.com/nokamoto/demo20-cli/internal/client"
	"github.com/nokamoto/demo20-cli/internal/config"
	"github.com/nokamoto/demo20-cli/internal/printer"
	"github.com/spf13/cobra"
)

func newCreate(value *config.Value, client client.Client) *cobra.Command {
	var (
		labels []string
	)

	cmd := template.NewArg0("create", "Create a new instance", func(cmd *cobra.Command) error {
		res, err := client.Compute().CreateInstance(context.Background(), &v1alpha.CreateInstanceRequest{
			Instance: &v1alpha.Instance{
				Labels: labels,
			},
		})
		if err != nil {
			return err
		}
		return printer.Proto(cmd, res)
	})

	cmd.Flags().StringArrayVar(&labels, "labels", nil, "labels")

	return cmd
}

func init() {
	RootCmd.AddCommand(newCreate(&config.Default, client.NewClient(&config.Default)))
}
