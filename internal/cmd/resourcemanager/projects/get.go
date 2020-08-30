package projects

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"

	"github.com/nokamoto/demo20-apis/cloud/resourcemanager/v1alpha"
	"github.com/nokamoto/demo20-cli/internal/client"
	"github.com/nokamoto/demo20-cli/internal/config"
	"github.com/nokamoto/demo20-cli/internal/template"
	"github.com/spf13/cobra"
)

func newGet(value *config.Value, client client.Client) *cobra.Command {
	return template.NewArg1Proto("get PROJECT_ID", "Get a project", func(cmd *cobra.Command, arg string) (proto.Message, error) {
		return client.ResourceManager().GetProject(context.Background(), &v1alpha.GetProjectRequest{
			Name: fmt.Sprintf("projects/%s", arg),
		})
	})
}

func init() {
	RootCmd.AddCommand(newGet(&config.Default, client.NewClient(&config.Default)))
}
