package auth

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/nokamoto/demo20-apis/cloud/iam/admin/v1alpha"
	"github.com/nokamoto/demo20-cli/internal/client"
	"github.com/nokamoto/demo20-cli/internal/cmd/iam/auth"
	"github.com/nokamoto/demo20-cli/internal/config"
	"github.com/spf13/cobra"
)

func newCanI(value *config.Value, client client.Client) *cobra.Command {
	return auth.NewCanI(func(permissionID string) (proto.Message, error) {
		return client.AdminIam().AuthorizeMachineUser(context.Background(), &v1alpha.AuthorizeMachineUserRequest{
			ApiKey:     value.MachineUserAPIKey,
			Permission: fmt.Sprintf("permissions/%s", permissionID),
			Parent:     fmt.Sprintf("projects/%s", value.ProjectID),
		})
	})
}

func init() {
	RootCmd.AddCommand(newCanI(&config.Default, client.NewClient(&config.Default)))
}
