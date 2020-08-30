package permissions

import (
	"context"

	"github.com/nokamoto/demo20-cli/internal/template"

	"github.com/nokamoto/demo20-apis/cloud/iam/admin/v1alpha"

	"github.com/nokamoto/demo20-cli/internal/client"
	"github.com/nokamoto/demo20-cli/internal/config"
	"github.com/nokamoto/demo20-cli/internal/printer"
	"github.com/spf13/cobra"
)

func newCreate(value *config.Value, client client.Client) *cobra.Command {
	cmd := template.NewArg1("create PERMISSION_ID", "Create a new permission", func(cmd *cobra.Command, arg string) error {
		res, err := client.AdminIam().CreatePermission(context.Background(), &v1alpha.CreatePermissionRequest{
			PermissionId: arg,
		})
		if err != nil {
			return err
		}

		return printer.Proto(cmd, res)
	})

	return cmd
}

func init() {
	RootCmd.AddCommand(newCreate(&config.Default, client.NewClient(&config.Default)))
}
