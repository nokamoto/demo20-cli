package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/go-cmp/cmp"
	"github.com/nokamoto/demo20-cli/internal/config"
	"gopkg.in/yaml.v2"

	"go.uber.org/zap"
)

const (
	loggerDebug = "LOGGER_DEBUG"
	grpcAddress = "GRPC_ADDRESS"
)

func main() {
	cfg := zap.NewProductionConfig()
	if len(os.Getenv(loggerDebug)) != 0 {
		cfg.Level.SetLevel(zap.DebugLevel)
	}

	logger, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	logger.Info("logger created", zap.Any("level", cfg.Level))

	view := func(expected config.Value) error {
		stdout, stderr, err := cloud(logger, "config", "view")
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

	xs := Scenarios{
		{
			Name: "config view - initial",
			Run: func(state State, logger *zap.Logger) (State, error) {
				return state, view(config.Value{
					GrpcAddress: "localhost:9000",
				})
			},
		},
		{
			Name: "config set",
			Run: func(state State, logger *zap.Logger) (State, error) {
				address := os.Getenv(grpcAddress)
				_, _, err := cloud(logger, "config", "set", "--grpc-address", address)
				if err != nil {
					return nil, err
				}

				return state, view(config.Value{
					GrpcAddress: address,
				})
			},
		},
	}

	xs.run(logger)
}
