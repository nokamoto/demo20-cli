package main

import (
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
			stdout, err := automatedtest.CloudF(logger, "iam", "admin", "permissions", "create", id)
			if err != nil {
				return nil, err
			}

			state["permissionid"] = id
			state["permission"] = stdout

			return state, automatedtest.Diff(
				stdout,
				&v1alpha.Permission{
					Name: fmt.Sprintf("permissions/%s", id),
				},
				&v1alpha.Permission{},
			)
		},
	},
	{
		Name: "iam admin roles create",
		Run: func(state automatedtest.State, logger *zap.Logger) (automatedtest.State, error) {
			id := automatedtest.RandomID()
			permissionID := state["permissionid"]

			stdout, err := automatedtest.CloudF(logger, "iam", "admin", "roles", "create", id, "--permission-ids", permissionID, "--display-name", "test role")
			if err != nil {
				return nil, err
			}

			expected := &v1alpha.Role{
				Name:        fmt.Sprintf("roles/%s", id),
				DisplayName: "test role",
				Permissions: []string{fmt.Sprintf("permissions/%s", permissionID)},
				Parent:      fmt.Sprintf("projects//"),
			}

			state["roleid"] = id
			state["role"] = stdout

			return state, automatedtest.Diff(stdout, expected, &v1alpha.Role{})
		},
	},
	{
		Name: "iam admin rolebindings add",
		Run: func(state automatedtest.State, logger *zap.Logger) (automatedtest.State, error) {
			roleID := state["roleid"]

			stdout, err := automatedtest.CloudF(logger, "iam", "admin", "rolebindings", "add", "--role", fmt.Sprintf("roles/%s", roleID), "--user", "todo")
			if err != nil {
				return nil, err
			}

			expected := &v1alpha.RoleBinding{
				Role:   fmt.Sprintf("roles/%s", roleID),
				User:   "todo",
				Parent: "projects//",
			}

			return state, automatedtest.Diff(stdout, expected, &v1alpha.RoleBinding{})
		},
	},
	{
		Name: "iam roles create",
		Run: func(state automatedtest.State, logger *zap.Logger) (automatedtest.State, error) {
			id := automatedtest.RandomID()
			permissionID := state["permissionid"]
			projectID := state[testProjectIDState]

			stdout, err := automatedtest.CloudF(logger, "resourcemanager", "projects", "roles", "create", id, "--permission-ids", permissionID, "--display-name", "test role")
			if err != nil {
				return nil, err
			}

			expected := &v1alpha.Role{
				Name:        fmt.Sprintf("projects/%s/roles/%s", projectID, id),
				DisplayName: "test role",
				Permissions: []string{fmt.Sprintf("permissions/%s", permissionID)},
				Parent:      fmt.Sprintf("projects/%s", projectID),
			}

			state["roleid"] = id
			state["role"] = stdout

			return state, automatedtest.Diff(stdout, expected, &v1alpha.Role{})
		},
	},
	{
		Name: "iam rolebindings add",
		Run: func(state automatedtest.State, logger *zap.Logger) (automatedtest.State, error) {
			roleID := state["roleid"]
			projectID := state[testProjectIDState]

			stdout, err := automatedtest.CloudF(logger, "resourcemanager", "projects", "rolebindings", "add", "--role", fmt.Sprintf("projects/%s/roles/%s", projectID, roleID), "--user", "todo")
			if err != nil {
				return nil, err
			}

			expected := &v1alpha.RoleBinding{
				Role:   fmt.Sprintf("projects/%s/roles/%s", projectID, roleID),
				User:   "todo",
				Parent: fmt.Sprintf("projects/%s", projectID),
			}

			return state, automatedtest.Diff(stdout, expected, &v1alpha.RoleBinding{})
		},
	},
}
