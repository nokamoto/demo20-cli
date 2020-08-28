package test

import (
	"bytes"
	"testing"

	mockcompute "github.com/nokamoto/demo20-cli/internal/test/mock/compute"

	"github.com/golang/mock/gomock"

	"github.com/nokamoto/demo20-cli/internal/client"
	"github.com/nokamoto/demo20-cli/internal/config"
	"github.com/spf13/cobra"
)

// Case represents a single test case.
type Case struct {
	Name  string
	Args  []string
	Value *config.Value
	Cmd   func(*config.Value, client.Client) *cobra.Command
	Mock  func(*MockClient)
	Check Check
}

// Cases represent a list of test cases.
type Cases []Case

// Run runs all test cases.
func (xs Cases) Run(t *testing.T) {
	for _, x := range xs {
		t.Run(x.Name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			c := MockClient{
				MockCompute: mockcompute.NewMockComputeClient(ctrl),
			}

			if x.Mock == nil {
				t.Fatal("no mock")
			}
			x.Mock(&c)

			if x.Cmd == nil {
				t.Fatal("no command")
			}
			cmd := x.Cmd(x.Value, &c)

			var stdout bytes.Buffer
			cmd.SetOut(&stdout)

			cmd.SetArgs(x.Args)

			err := cmd.Execute()

			if x.Check == nil {
				t.Fatal("no check")
			}
			x.Check(t, stdout.String(), err)
		})
	}
}

// Args returns given arguments.
func Args(args ...string) []string {
	return args
}
