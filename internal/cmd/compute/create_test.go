package compute

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nokamoto/demo20-apis/cloud/compute/v1alpha"

	"github.com/nokamoto/demo20-cli/internal/config"
	"github.com/nokamoto/demo20-cli/internal/test"
)

func Test_newCreate(t *testing.T) {
	xs := test.Cases{
		{
			Name:  "OK",
			Args:  test.Args("--labels", "foo", "--labels", "bar"),
			Value: &config.Value{},
			Cmd:   newCreate,
			Mock: func(mock *test.MockClient) {
				mock.MockCompute.EXPECT().CreateInstance(gomock.Any(), &v1alpha.CreateInstanceRequest{
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
