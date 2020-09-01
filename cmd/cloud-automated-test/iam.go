package main

import (
	"fmt"

	"github.com/golang/protobuf/jsonpb"
	admin "github.com/nokamoto/demo20-apis/cloud/iam/admin/v1alpha"
	"github.com/nokamoto/demo20-apis/cloud/iam/v1alpha"
	"github.com/nokamoto/demo20-cli/internal/automatedtest"
	"go.uber.org/zap"
	"google.golang.org/protobuf/testing/protocmp"
)

var configSetMachineUserAPIKey = automatedtest.Scenario{
	Name: "config set --machine-user-api-key",
	Run: func(state automatedtest.State, logger *zap.Logger) (automatedtest.State, error) {
		var machineUser v1alpha.MachineUser
		err := jsonpb.UnmarshalString(state["machineuser"], &machineUser)
		if err != nil {
			return nil, err
		}

		expected, err := currentConfig(logger)
		if err != nil {
			return nil, err
		}

		expected.MachineUserAPIKey = machineUser.GetApiKey()

		_, err = automatedtest.CloudF(logger, "config", "set", "--machine-user-api-key", expected.MachineUserAPIKey)
		if err != nil {
			return nil, err
		}

		return state, configView(logger, expected)
	},
}

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
		Name: "iam admin machineusers create",
		Run: func(state automatedtest.State, logger *zap.Logger) (automatedtest.State, error) {
			stdout, err := automatedtest.CloudF(logger, "iam", "admin", "machineusers", "create", "--display-name", "test machine user")
			if err != nil {
				return nil, err
			}

			state["machineuser"] = stdout

			expected := &v1alpha.MachineUser{
				DisplayName: "test machine user",
				Parent:      "projects//",
			}

			return state, automatedtest.Diff(stdout, expected, &v1alpha.MachineUser{}, protocmp.IgnoreFields(&v1alpha.MachineUser{}, "name", "api_key"))
		},
	},
	configSetMachineUserAPIKey,
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

			var machineUser v1alpha.MachineUser
			err := jsonpb.UnmarshalString(state["machineuser"], &machineUser)
			if err != nil {
				return nil, err
			}

			stdout, err := automatedtest.CloudF(logger, "iam", "admin", "rolebindings", "add", "--role", fmt.Sprintf("roles/%s", roleID), "--user", machineUser.GetName())
			if err != nil {
				return nil, err
			}

			expected := &v1alpha.RoleBinding{
				Role:   fmt.Sprintf("roles/%s", roleID),
				User:   machineUser.GetName(),
				Parent: "projects//",
			}

			return state, automatedtest.Diff(stdout, expected, &v1alpha.RoleBinding{})
		},
	},
	{
		Name: "iam admin auth can-i",
		Run: func(state automatedtest.State, logger *zap.Logger) (automatedtest.State, error) {
			stdout, err := automatedtest.CloudF(logger, "iam", "admin", "auth", "can-i", "--permission-id", state["permissionid"])
			if err != nil {
				return nil, err
			}

			var machineUser v1alpha.MachineUser
			err = jsonpb.UnmarshalString(state["machineuser"], &machineUser)
			if err != nil {
				return nil, err
			}

			expected := &admin.AuthorizeMachineUserResponse{
				MachineUser: &machineUser,
			}
			expected.MachineUser.ApiKey = ""

			return state, automatedtest.Diff(stdout, expected, &admin.AuthorizeMachineUserResponse{})
		},
	},
	{
		Name: "iam machineusers create",
		Run: func(state automatedtest.State, logger *zap.Logger) (automatedtest.State, error) {
			stdout, err := automatedtest.CloudF(logger, "resourcemanager", "projects", "machineusers", "create", "--display-name", "test machine user")
			if err != nil {
				return nil, err
			}

			state["machineuser"] = stdout

			expected := &v1alpha.MachineUser{
				DisplayName: "test machine user",
				Parent:      fmt.Sprintf("projects/%s", state[testProjectIDState]),
			}

			return state, automatedtest.Diff(stdout, expected, &v1alpha.MachineUser{}, protocmp.IgnoreFields(&v1alpha.MachineUser{}, "name", "api_key"))
		},
	},
	configSetMachineUserAPIKey,
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

			var machineUser v1alpha.MachineUser
			err := jsonpb.UnmarshalString(state["machineuser"], &machineUser)
			if err != nil {
				return nil, err
			}

			stdout, err := automatedtest.CloudF(logger, "resourcemanager", "projects", "rolebindings", "add", "--role", fmt.Sprintf("projects/%s/roles/%s", projectID, roleID), "--user", machineUser.GetName())
			if err != nil {
				return nil, err
			}

			expected := &v1alpha.RoleBinding{
				Role:   fmt.Sprintf("projects/%s/roles/%s", projectID, roleID),
				User:   machineUser.GetName(),
				Parent: fmt.Sprintf("projects/%s", projectID),
			}

			return state, automatedtest.Diff(stdout, expected, &v1alpha.RoleBinding{})
		},
	},
	{
		Name: "iam auth can-i",
		Run: func(state automatedtest.State, logger *zap.Logger) (automatedtest.State, error) {
			stdout, err := automatedtest.CloudF(logger, "resourcemanager", "projects", "auth", "can-i", "--permission-id", state["permissionid"])
			if err != nil {
				return nil, err
			}

			var machineUser v1alpha.MachineUser
			err = jsonpb.UnmarshalString(state["machineuser"], &machineUser)
			if err != nil {
				return nil, err
			}

			expected := &admin.AuthorizeMachineUserResponse{
				MachineUser: &machineUser,
			}
			expected.MachineUser.ApiKey = ""

			return state, automatedtest.Diff(stdout, expected, &admin.AuthorizeMachineUserResponse{})
		},
	},
}
