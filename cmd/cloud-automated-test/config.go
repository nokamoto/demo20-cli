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

func configView(logger *zap.Logger, expected config.Value) error {
	stdout, stderr, err := automatedtest.Cloud(logger, "config", "view")
	if err != nil {
		return err
	}

	var actual config.Value
	err = yaml.Unmarshal([]byte(stdout), &actual)
	if err != nil {
		return err
	}

	if diff := cmp.Diff(expected, actual); len(diff) != 0 {
		return fmt.Errorf("unexpected configuration: %s", diff)
	}

	if len(stderr) != 0 {
		return fmt.Errorf("expected no stderr: %s", stderr)
	}

	return nil
}

func configSet(env string) automatedtest.Scenario {
	return automatedtest.Scenario{
		Name: "config set",
		Run: func(state automatedtest.State, logger *zap.Logger) (automatedtest.State, error) {
			address := os.Getenv(env)
			_, _, err := automatedtest.Cloud(logger, "config", "set", "--grpc-address", address)
			if err != nil {
				return nil, err
			}

			return state, configView(logger, config.Value{
				GrpcAddress: address,
			})
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
