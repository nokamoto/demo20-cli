package test

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/google/go-cmp/cmp"
	"github.com/nokamoto/demo20-cli/internal/printer"
)

// Check asserts stdout and error.
type Check func(*testing.T, string, error)

// Succeeded asserts that the expected message in stdout.
func Succeeded(expected proto.Message) Check {
	return func(t *testing.T, stdout string, err error) {
		t.Helper()
		if err != nil {
			t.Fatalf("expected no error but actual %v", err)
		}

		s, err := printer.ProtoString(expected)
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(s+"\n", stdout); len(diff) != 0 {
			t.Error(diff)
		}
	}
}
