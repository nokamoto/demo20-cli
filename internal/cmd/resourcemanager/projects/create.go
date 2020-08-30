package projects

import (
	"context"

	"github.com/golang/protobuf/proto"

	"github.com/nokamoto/demo20-apis/cloud/resourcemanager/v1alpha"
	"github.com/nokamoto/demo20-cli/internal/client"
	"github.com/nokamoto/demo20-cli/internal/config"
	"github.com/nokamoto/demo20-cli/internal/template"
	"github.com/spf13/cobra"
)

func newCreate(value *config.Value, client client.Client) *cobra.Command {
	var (
		displayName string
	)

	cmd := template.NewArg1Proto("create PROJECT_ID", "Create a project", func(cmd *cobra.Command, arg string) (proto.Message, error) {
		return client.ResourceManager().CreateProject(context.Background(), &v1alpha.CreateProjectRequest{
			ProjectId: arg,
			Project: &v1alpha.Project{
				DisplayName: displayName,
			},
		})
	})

	cmd.Flags().StringVar(&displayName, "display-name", "", "display name")

	return cmd
}

func init() {
	RootCmd.AddCommand(newCreate(&config.Default, client.NewClient(&config.Default)))
}
