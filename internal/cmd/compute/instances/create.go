package instances

import (
	"context"

	"github.com/golang/protobuf/proto"

	"github.com/nokamoto/demo20-cli/internal/template"

	"github.com/nokamoto/demo20-apis/cloud/compute/v1alpha"
	"github.com/nokamoto/demo20-cli/internal/client"
	"github.com/nokamoto/demo20-cli/internal/config"
	"github.com/spf13/cobra"
)

func newCreate(value *config.Value, client client.Client) *cobra.Command {
	var (
		labels []string
	)

	cmd := template.NewArg0Proto("create", "Create a new instance", func(cmd *cobra.Command) (proto.Message, error) {
		return client.Compute().CreateInstance(context.Background(), &v1alpha.CreateInstanceRequest{
			Instance: &v1alpha.Instance{
				Labels: labels,
			},
		})
	})

	cmd.Flags().StringArrayVar(&labels, "labels", nil, "labels")

	return cmd
}

func init() {
	RootCmd.AddCommand(newCreate(&config.Default, client.NewClient(&config.Default)))
}
