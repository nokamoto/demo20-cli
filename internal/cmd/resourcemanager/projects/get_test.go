package projects

import (
	"context"
	"testing"

	"github.com/nokamoto/demo20-apis/cloud/api"
	"github.com/nokamoto/demo20-apis/cloud/resourcemanager/v1alpha"
	"github.com/nokamoto/demo20-apps/pkg/sdk/metadata"
	"github.com/nokamoto/demo20-cli/internal/config"
	"github.com/nokamoto/demo20-cli/internal/test"
)

func Test_newGet(t *testing.T) {
	xs := test.Cases{
		test.Case{
			Name: "OK",
			Args: test.Args("foo"),
			Value: &config.Value{
				ProjectID:         "test",
				MachineUserAPIKey: "mu",
			},
			Cmd: newGet,
			Mock: func(c *test.MockClient) {
				c.MockResourceManager.EXPECT().GetProject(
					metadata.AppendToOutgoingContextF(context.Background(), &api.Metadata{
						Credential: &api.Metadata_MachineUserApiKey{
							MachineUserApiKey: "mu",
						},
						Parent: "projects/test",
					}),
					&v1alpha.GetProjectRequest{}).Return(&v1alpha.Project{
					Name:        "projects/foo",
					DisplayName: "bar",
				}, nil)
			},
			Check: test.Succeeded(&v1alpha.Project{
				Name:        "projects/foo",
				DisplayName: "bar",
			}),
		},
	}

	xs.Run(t)
}
