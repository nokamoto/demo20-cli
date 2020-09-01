package rolebindings

import (
	"github.com/golang/protobuf/proto"
	"github.com/nokamoto/demo20-apis/cloud/iam/v1alpha"
	"github.com/nokamoto/demo20-cli/internal/client"
	"github.com/nokamoto/demo20-cli/internal/cmd/iam/rolebindings"
	"github.com/nokamoto/demo20-cli/internal/config"
	"github.com/spf13/cobra"
)

func newAdd(value *config.Value, client client.Client) *cobra.Command {
	return rolebindings.NewAdd(func(role, user string) (proto.Message, error) {
		return client.Iam().AddRoleBinding(value.OutgoingContext(), &v1alpha.AddRoleBindingRequest{
			RoleBinding: &v1alpha.RoleBinding{
				Role: role,
				User: user,
			},
		})
	})
}

func init() {
	RootCmd.AddCommand(newAdd(&config.Default, client.NewClient(&config.Default)))
}
