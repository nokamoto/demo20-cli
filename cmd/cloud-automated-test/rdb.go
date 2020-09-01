package main

import (
	"fmt"

	"github.com/nokamoto/demo20-apis/cloud/rdb/v1alpha"
	"github.com/nokamoto/demo20-cli/internal/automatedtest"
	"go.uber.org/zap"
	"google.golang.org/protobuf/testing/protocmp"
)

var rdbScenarios = automatedtest.Scenarios{
	configSet(rdbGrpcAddress),
	{
		Name: "rdb clusters create",
		Run: func(state automatedtest.State, logger *zap.Logger) (automatedtest.State, error) {
			id := automatedtest.RandomID()

			stdout, err := automatedtest.CloudF(logger, "rdb", "clusters", "create", id, "--replicas", "1")
			if err != nil {
				return nil, err
			}

			return state, automatedtest.Diff(
				stdout,
				&v1alpha.Cluster{
					Name:     fmt.Sprintf("clusters/%s", id),
					Replicas: 1,
					Parent:   fmt.Sprintf("projects/%s", state[testProjectIDState]),
				},
				&v1alpha.Cluster{},
				protocmp.IgnoreFields(&v1alpha.Cluster{}, "instances"),
			)
		},
	},
}
