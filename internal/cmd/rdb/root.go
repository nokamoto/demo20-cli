package rdb

import (
	"github.com/nokamoto/demo20-cli/internal/cmd/rdb/clusters"
	"github.com/nokamoto/demo20-cli/internal/template"
)

// RootCmd is a root of v1alpha sub commands.
var RootCmd = template.NewRoot("rdb", "A rdb service management tool")

func init() {
	RootCmd.AddCommand(clusters.RootCmd)
}
