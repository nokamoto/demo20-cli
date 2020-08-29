package permissions

import (
	"context"

	"github.com/nokamoto/demo20-apis/cloud/iam/admin/v1alpha"

	"github.com/nokamoto/demo20-cli/internal/client"
	"github.com/nokamoto/demo20-cli/internal/config"
	"github.com/nokamoto/demo20-cli/internal/printer"
	"github.com/spf13/cobra"
)

func newCreate(value *config.Value, client client.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "create PERMISSION_ID",
		Short:         "Create a new permission",
		Args:          cobra.ExactArgs(1),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			res, err := client.AdminIam().CreatePermission(context.Background(), &v1alpha.CreatePermissionRequest{
				PermissionId: args[0],
			})
			if err != nil {
				return err
			}

			return printer.Proto(cmd, res)
		},
	}
	return cmd
}

func init() {
	RootCmd.AddCommand(newCreate(&config.Default, client.NewClient(&config.Default)))
}
