package instances

import (
	"context"
	"testing"

	"github.com/nokamoto/demo20-apis/cloud/api"
	"github.com/nokamoto/demo20-apps/pkg/sdk/metadata"

	"github.com/nokamoto/demo20-apis/cloud/compute/v1alpha"
	"github.com/nokamoto/demo20-cli/internal/config"
	"github.com/nokamoto/demo20-cli/internal/test"
)

func Test_newCreate(t *testing.T) {
	xs := test.Cases{
		{
			Name: "OK",
			Args: test.Args("--labels", "foo", "--labels", "bar"),
			Value: &config.Value{
				ProjectID:         "test",
				MachineUserAPIKey: "mu",
			},
			Cmd: newCreate,
			Mock: func(mock *test.MockClient) {
				mock.MockCompute.EXPECT().CreateInstance(
					metadata.AppendToOutgoingContextF(context.Background(), &api.Metadata{
						Credential: &api.Metadata_MachineUserApiKey{
							MachineUserApiKey: "mu",
						},
						Parent: "projects/test",
					}),
					&v1alpha.CreateInstanceRequest{
						Instance: &v1alpha.Instance{
							Labels: []string{"foo", "bar"},
						},
					}).Return(&v1alpha.Instance{
					Name:   "instances/baz",
					Parent: "projects/qux",
					Labels: []string{"foo", "bar"},
				}, nil)
			},
			Check: test.Succeeded(&v1alpha.Instance{
				Name:   "instances/baz",
				Parent: "projects/qux",
				Labels: []string{"foo", "bar"},
			}),
		},
	}

	xs.Run(t)
}
