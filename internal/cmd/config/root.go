package config

import (
	"github.com/nokamoto/demo20-cli/internal/template"
)

// RootCmd is a root of config sub commands.
var RootCmd = template.NewRoot("config", "A commandline configration management tool")
