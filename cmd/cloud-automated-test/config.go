package main

import (
	"fmt"
	"os"

	"github.com/nokamoto/demo20-cli/internal/automatedtest"

	"github.com/google/go-cmp/cmp"
	"github.com/nokamoto/demo20-cli/internal/config"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

func currentConfig(logger *zap.Logger) (config.Value, error) {
	var actual config.Value
	stdout, err := automatedtest.CloudF(logger, "config", "view")
	if err != nil {
		return actual, err
	}

	err = yaml.Unmarshal([]byte(stdout), &actual)
	if err != nil {
		return actual, err
	}

	return actual, nil
}

func configView(logger *zap.Logger, expected config.Value) error {
	actual, err := currentConfig(logger)
	if err != nil {
		return err
	}

	if diff := cmp.Diff(expected, actual); len(diff) != 0 {
		return fmt.Errorf("unexpected configuration: %s", diff)
	}

	return nil
}

func configSet(env string) automatedtest.Scenario {
	return automatedtest.Scenario{
		Name: fmt.Sprintf("config set --grpc-address $%s", env),
		Run: func(state automatedtest.State, logger *zap.Logger) (automatedtest.State, error) {
			expected, err := currentConfig(logger)
			if err != nil {
				return nil, err
			}

			address := os.Getenv(env)
			_, err = automatedtest.CloudF(logger, "config", "set", "--grpc-address", address)
			if err != nil {
				return nil, err
			}

			expected.GrpcAddress = address

			return state, configView(logger, expected)
		},
	}
}

var configScenarios = automatedtest.Scenarios{
	{
		Name: "config view - initial",
		Run: func(state automatedtest.State, logger *zap.Logger) (automatedtest.State, error) {
			return state, configView(logger, config.Value{
				GrpcAddress: "localhost:9000",
			})
		},
	},
}
