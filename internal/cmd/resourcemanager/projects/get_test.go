package projects

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nokamoto/demo20-apis/cloud/resourcemanager/v1alpha"
	"github.com/nokamoto/demo20-cli/internal/config"
	"github.com/nokamoto/demo20-cli/internal/test"
)

func Test_newGet(t *testing.T) {
	xs := test.Cases{
		test.Case{
			Name:  "OK",
			Args:  test.Args("foo"),
			Value: &config.Value{},
			Cmd:   newGet,
			Mock: func(c *test.MockClient) {
				c.MockResourceManager.EXPECT().GetProject(gomock.Any(), &v1alpha.GetProjectRequest{
					Name: "projects/foo",
				}).Return(&v1alpha.Project{
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
