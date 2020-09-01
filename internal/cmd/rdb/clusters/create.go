package clusters

import (
	"github.com/golang/protobuf/proto"
	"github.com/nokamoto/demo20-apis/cloud/rdb/v1alpha"
	"github.com/nokamoto/demo20-cli/internal/client"
	"github.com/nokamoto/demo20-cli/internal/config"
	"github.com/nokamoto/demo20-cli/internal/template"
	"github.com/spf13/cobra"
)

func newCreate(value *config.Value, client client.Client) *cobra.Command {
	var (
		replicas int32
	)

	cmd := template.NewArg1Proto("create CLUSTER_ID", "Create a cluster", func(cmd *cobra.Command, arg string) (proto.Message, error) {
		return client.Rdb().CreateCluster(value.OutgoingContext(), &v1alpha.CreateClusterRequest{
			ClusterId: arg,
			Cluster: &v1alpha.Cluster{
				Replicas: replicas,
			},
		})
	})

	cmd.Flags().Int32Var(&replicas, "replicas", 0, "replicas")

	return cmd
}

func init() {
	RootCmd.AddCommand(newCreate(&config.Default, client.NewClient(&config.Default)))
}
