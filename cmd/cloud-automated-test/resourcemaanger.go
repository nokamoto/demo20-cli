package main

import (
	"errors"
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

			stdout, stderr, err := automatedtest.Cloud(logger, "resourcemanager", "projects", "create", id, "--display-name", displayName)
			if len(stderr) != 0 {
				return nil, errors.New(stderr)
			}
			if err != nil {
				return nil, err
			}

			state["project-id"] = id
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
		Name: "resourcemanager projects get",
		Run: func(state automatedtest.State, logger *zap.Logger) (automatedtest.State, error) {
			var expected v1alpha.Project
			err := jsonpb.UnmarshalString(state["project"], &expected)
			if err != nil {
				return nil, err
			}

			stdout, stderr, err := automatedtest.Cloud(logger, "resourcemanager", "projects", "get", state["project-id"])
			if len(stderr) != 0 {
				return nil, errors.New(stderr)
			}
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
