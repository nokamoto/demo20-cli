package main

import (
	"fmt"

	"github.com/golang/protobuf/jsonpb"
	"github.com/nokamoto/demo20-apis/cloud/resourcemanager/v1alpha"
	"github.com/nokamoto/demo20-cli/internal/automatedtest"
	"go.uber.org/zap"
)

var resourcemanagerScenarios = automatedtest.Scenarios{
	configSet(resourcemanagerGrpcAddress),
	{
		Name: "resourcemanager projects create",
		Run: func(state automatedtest.State, logger *zap.Logger) (automatedtest.State, error) {
			id := automatedtest.RandomID()
			displayName := fmt.Sprintf("test %s display name", id)

			stdout, err := automatedtest.CloudF(logger, "resourcemanager", "projects", "create", id, "--display-name", displayName)
			if err != nil {
				return nil, err
			}

			state[testProjectIDState] = id
			state["project"] = stdout

			return state, automatedtest.Diff(
				stdout,
				&v1alpha.Project{
					Name:        fmt.Sprintf("projects/%s", id),
					DisplayName: displayName,
				},
				&v1alpha.Project{},
			)
		},
	},
	{
		Name: "config set --project-id",
		Run: func(state automatedtest.State, logger *zap.Logger) (automatedtest.State, error) {
			expected, err := currentConfig(logger)
			if err != nil {
				return nil, err
			}

			expected.ProjectID = state[testProjectIDState]

			_, err = automatedtest.CloudF(logger, "config", "set", "--project-id", expected.ProjectID)
			if err != nil {
				return nil, err
			}

			return state, configView(logger, expected)
		},
	},
	{
		Name: "resourcemanager projects get",
		Run: func(state automatedtest.State, logger *zap.Logger) (automatedtest.State, error) {
			var expected v1alpha.Project
			err := jsonpb.UnmarshalString(state["project"], &expected)
			if err != nil {
				return nil, err
			}

			stdout, err := automatedtest.CloudF(logger, "resourcemanager", "projects", "get", state[testProjectIDState])
			if err != nil {
				return nil, err
			}

			return state, automatedtest.Diff(
				stdout,
				&expected,
				&v1alpha.Project{},
			)
		},
	},
}
