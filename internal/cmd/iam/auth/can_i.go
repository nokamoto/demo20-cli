package auth

import (
	"github.com/golang/protobuf/proto"
	"github.com/nokamoto/demo20-cli/internal/template"
	"github.com/spf13/cobra"
)

// NewCanI returns a template of `auth can-i`
func NewCanI(f func(string) (proto.Message, error)) *cobra.Command {
	var (
		permissionID string
	)

	cmd := template.NewArg0Proto("can-i", "Checks whether I can do an action", func(cmd *cobra.Command) (proto.Message, error) {
		return f(permissionID)
	})

	cmd.Flags().StringVar(&permissionID, "permission-id", "", "permission id")

	return cmd
}
