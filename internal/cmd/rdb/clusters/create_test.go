package clusters

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nokamoto/demo20-apis/cloud/rdb/v1alpha"
	"github.com/nokamoto/demo20-cli/internal/config"
	"github.com/nokamoto/demo20-cli/internal/test"
)

func Test_newCreate(t *testing.T) {
	xs := test.Cases{
		{
			Name:  "OK",
			Args:  test.Args("foo", "--replicas", "1"),
			Value: &config.Value{},
			Cmd:   newCreate,
			Mock: func(c *test.MockClient) {
				c.MockRdb.EXPECT().CreateCluster(gomock.Any(), &v1alpha.CreateClusterRequest{
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
