package machineusers

import (
	"github.com/golang/protobuf/proto"
	"github.com/nokamoto/demo20-apis/cloud/iam/v1alpha"
	"github.com/nokamoto/demo20-cli/internal/client"
	"github.com/nokamoto/demo20-cli/internal/cmd/iam/machineusers"
	"github.com/nokamoto/demo20-cli/internal/config"
	"github.com/spf13/cobra"
)

func newCreate(value *config.Value, client client.Client) *cobra.Command {
	return machineusers.NewCreate(func(displayName string) (proto.Message, error) {
		return client.Iam().CreateMachineUser(value.OutgoingContext(), &v1alpha.CreateMachineUserRequest{
			MachineUser: &v1alpha.MachineUser{
				DisplayName: displayName,
			},
		})
	})
}

func init() {
	RootCmd.AddCommand(newCreate(&config.Default, client.NewClient(&config.Default)))
}
