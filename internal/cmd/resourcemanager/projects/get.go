package projects

import (
	"github.com/nokamoto/demo20-cli/internal/client"
	"github.com/nokamoto/demo20-cli/internal/config"
	"github.com/nokamoto/demo20-cli/internal/template"
	"github.com/spf13/cobra"
)

func newGet(value *config.Value, client client.Client) *cobra.Command {
	return template.NewArg1("get PROJECT_ID", "Get a project", func(cmd *cobra.Command, arg string) error {
		return nil
	})
}

func init() {
	RootCmd.AddCommand(newGet(&config.Default, client.NewClient(&config.Default)))
}
