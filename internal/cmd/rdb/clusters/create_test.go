package clusters

import (
	"context"
	"testing"

	"github.com/nokamoto/demo20-apis/cloud/api"
	"github.com/nokamoto/demo20-apis/cloud/rdb/v1alpha"
	"github.com/nokamoto/demo20-apps/pkg/sdk/metadata"
	"github.com/nokamoto/demo20-cli/internal/config"
	"github.com/nokamoto/demo20-cli/internal/test"
)

func Test_newCreate(t *testing.T) {
	xs := test.Cases{
		{
			Name: "OK",
			Args: test.Args("foo", "--replicas", "1"),
			Value: &config.Value{
				ProjectID:         "test",
				MachineUserAPIKey: "mu",
			},
			Cmd: newCreate,
			Mock: func(c *test.MockClient) {
				c.MockRdb.EXPECT().CreateCluster(
					metadata.AppendToOutgoingContextF(context.Background(), &api.Metadata{
						Credential: &api.Metadata_MachineUserApiKey{
							MachineUserApiKey: "mu",
						},
						Parent: "projects/test",
					}),
					&v1alpha.CreateClusterRequest{
						ClusterId: "foo",
						Cluster: &v1alpha.Cluster{
							Replicas: 1,
						},
					}).Return(&v1alpha.Cluster{
					Name:      "clusters/foo",
					Replicas:  1,
					Instances: []string{"bar", "baz"},
					Parent:    "projects/todo",
				}, nil)
			},
			Check: test.Succeeded(&v1alpha.Cluster{
				Name:      "clusters/foo",
				Replicas:  1,
				Instances: []string{"bar", "baz"},
				Parent:    "projects/todo",
			}),
		},
	}

	xs.Run(t)
}
