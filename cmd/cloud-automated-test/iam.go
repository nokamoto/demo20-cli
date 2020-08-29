package main

import (
	"errors"
	"fmt"

	"github.com/nokamoto/demo20-apis/cloud/iam/v1alpha"

	"github.com/nokamoto/demo20-cli/internal/automatedtest"
	"go.uber.org/zap"
)

var iamScenarios = automatedtest.Scenarios{
	configSet(iamGrpcAddress),
	{
		Name: "iam admin permissions create",
		Run: func(state automatedtest.State, logger *zap.Logger) (automatedtest.State, error) {
			id := automatedtest.RandomID()
			stdout, stderr, err := automatedtest.Cloud(logger, "iam", "admin", "permissions", "create", id)
			if len(stderr) != 0 {
				return nil, errors.New(stderr)
			}
			if err != nil {
				return nil, err
			}

			return state, automatedtest.Diff(
				stdout,
				&v1alpha.Permission{
					Name: fmt.Sprintf("permissions/%s", id),
				},
				&v1alpha.Permission{},
			)
		},
	},
}
