package permissions

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/nokamoto/demo20-apis/cloud/iam/admin/v1alpha"
	"github.com/nokamoto/demo20-cli/internal/client"
	"github.com/nokamoto/demo20-cli/internal/config"
	"github.com/nokamoto/demo20-cli/internal/template"
	"github.com/spf13/cobra"
)

func newCreate(value *config.Value, client client.Client) *cobra.Command {
	cmd := template.NewArg1Proto("create PERMISSION_ID", "Create a new permission", func(cmd *cobra.Command, arg string) (proto.Message, error) {
		return client.AdminIam().CreatePermission(context.Background(), &v1alpha.CreatePermissionRequest{
			PermissionId: arg,
		})
	})

	return cmd
}

func init() {
	RootCmd.AddCommand(newCreate(&config.Default, client.NewClient(&config.Default)))
}
