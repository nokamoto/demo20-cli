package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/nokamoto/demo20-cli/internal/automatedtest"
	"go.uber.org/zap"
)

const (
	loggerDebug                = "LOGGER_DEBUG"
	computeGrpcAddress         = "COMPUTE_GRPC_ADDRESS"
	iamGrpcAddress             = "IAM_GRPC_ADDRESS"
	resourcemanagerGrpcAddress = "RESOURCEMANAGER_GRPC_ADDRESS"
	rdbGrpcAddress             = "RDB_GRPC_ADDRESS"
)

func main() {
	rand.Seed(time.Now().Unix())

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

	var xs automatedtest.Scenarios
	xs = append(xs, configScenarios...)
	xs = append(xs, computeScenarios...)
	xs = append(xs, iamScenarios...)
	xs = append(xs, rdbScenarios...)
	xs.Run(logger)
}
