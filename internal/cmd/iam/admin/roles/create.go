package roles

import (
	"context"

	"github.com/golang/protobuf/proto"
	admin "github.com/nokamoto/demo20-apis/cloud/iam/admin/v1alpha"
	"github.com/nokamoto/demo20-apis/cloud/iam/v1alpha"
	"github.com/nokamoto/demo20-cli/internal/client"
	"github.com/nokamoto/demo20-cli/internal/cmd/iam/roles"
	"github.com/nokamoto/demo20-cli/internal/config"
	"github.com/spf13/cobra"
)

func newCreate(value *config.Value, client client.Client) *cobra.Command {
	return roles.NewCreate(func(arg string, displayName string, permissions []string) (proto.Message, error) {
		return client.AdminIam().CreateRole(context.Background(), &admin.CreateRoleRequest{
			RoleId: arg,
			Role: &v1alpha.Role{
				DisplayName: displayName,
				Permissions: permissions,
			},
		})
	})
}

func init() {
	RootCmd.AddCommand(newCreate(&config.Default, client.NewClient(&config.Default)))
}
