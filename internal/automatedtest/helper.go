package automatedtest

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"

	"go.uber.org/zap"
)

// Cloud executes the cloud command.
func Cloud(logger *zap.Logger, args ...string) (string, string, error) {
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

// CloudF executes the cloud command expected no stderr.
func CloudF(logger *zap.Logger, args ...string) (string, error) {
	stdout, stderr, err := Cloud(logger, args...)
	if len(stderr) != 0 {
		return "", fmt.Errorf("unexpected stderr: %s: %w", stderr, err)
	}
	return stdout, err
}

// Diff asserts that the standard output is the expected proto message.
func Diff(stdout string, expected proto.Message, actual proto.Message, opts ...cmp.Option) error {
	err := jsonpb.UnmarshalString(stdout, actual)
	if err != nil {
		return err
	}

	opts = append(opts, protocmp.Transform())

	if diff := cmp.Diff(expected, actual, opts...); len(diff) != 0 {
		return fmt.Errorf("diff=%s", diff)
	}

	return nil
}
