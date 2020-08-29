package main

import (
	"errors"

	"github.com/nokamoto/demo20-apis/cloud/compute/v1alpha"
	"github.com/nokamoto/demo20-cli/internal/automatedtest"
	"go.uber.org/zap"
	"google.golang.org/protobuf/testing/protocmp"
)

var computeScenarios = automatedtest.Scenarios{
	configSet(computeGrpcAddress),
	{
		Name: "compute instances create",
		Run: func(state automatedtest.State, logger *zap.Logger) (automatedtest.State, error) {
			stdout, stderr, err := automatedtest.Cloud(logger, "compute", "instances", "create", "--labels", "foo")
			if len(stderr) != 0 {
				return nil, errors.New(stderr)
			}
			if err != nil {
				return nil, err
			}

			return state, automatedtest.Diff(
				stdout,
				&v1alpha.Instance{
					Parent: "projects/todo",
					Labels: []string{"foo"},
				},
				&v1alpha.Instance{},
				protocmp.IgnoreFields(&v1alpha.Instance{}, "name"),
			)
		},
	},
}
