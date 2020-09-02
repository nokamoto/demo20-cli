package machineusers

import (
	"github.com/golang/protobuf/proto"
	"github.com/nokamoto/demo20-cli/internal/template"
	"github.com/spf13/cobra"
)

// NewCreate returns a template of `machineusers create`
func NewCreate(f func(string) (proto.Message, error)) *cobra.Command {
	var (
		displayName string
	)

	cmd := template.NewArg0Proto("create", "Create a new machine user", func(cmd *cobra.Command) (proto.Message, error) {
		return f(displayName)
	})

	cmd.Flags().StringVar(&displayName, "display-name", "", "display name")

	return cmd
}
