package rolebindings

import (
	"github.com/golang/protobuf/proto"
	"github.com/nokamoto/demo20-cli/internal/template"
	"github.com/spf13/cobra"
)

// NewAdd returns a template of `rolebindings add`.
func NewAdd(f func(string, string) (proto.Message, error)) *cobra.Command {
	var (
		role string
		user string
	)

	cmd := template.NewArg0Proto("add", "Add a new role binding", func(cmd *cobra.Command) (proto.Message, error) {
		return f(role, user)
	})

	cmd.Flags().StringVar(&role, "role", "", "role name")
	cmd.Flags().StringVar(&user, "user", "", "user name")

	return cmd
}
