package roles

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/nokamoto/demo20-cli/internal/template"
	"github.com/spf13/cobra"
)

// NewCreate returns a template of `roles create`.
func NewCreate(f func(string, string, []string) (proto.Message, error)) *cobra.Command {
	var (
		displayName   string
		permissionIDs []string
	)

	cmd := template.NewArg1Proto("create ROLE_ID", "Create a new role", func(cmd *cobra.Command, arg string) (proto.Message, error) {
		var names []string
		for _, p := range permissionIDs {
			names = append(names, fmt.Sprintf("permissions/%s", p))
		}
		return f(arg, displayName, names)
	})

	cmd.Flags().StringVar(&displayName, "display-name", "", "display name")
	cmd.Flags().StringSliceVar(&permissionIDs, "permission-ids", nil, "permission ids")

	return cmd
}
