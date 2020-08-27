package main

import (
	"bytes"
	"os/exec"

	"go.uber.org/zap"
)

func cloud(logger *zap.Logger, args ...string) (string, string, error) {
	cmd := exec.Command("cloud", args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	so := stdout.String()
	se := stderr.String()

	logger.Debug("execute", zap.Strings("args", args), zap.String("stdout", so), zap.String("stderr", se), zap.Error(err))

	return so, se, err
}
