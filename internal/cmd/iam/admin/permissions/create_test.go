package permissions

import (
	"testing"

	"github.com/golang/mock/gomock"
	admin "github.com/nokamoto/demo20-apis/cloud/iam/admin/v1alpha"
	"github.com/nokamoto/demo20-apis/cloud/iam/v1alpha"
	"github.com/nokamoto/demo20-cli/internal/config"
	"github.com/nokamoto/demo20-cli/internal/test"
)

func Test_newCreate(t *testing.T) {
	xs := test.Cases{
		{
			Name:  "OK",
			Args:  test.Args("foo"),
			Value: &config.Value{},
			Cmd:   newCreate,
			Mock: func(c *test.MockClient) {
				c.MockAdmin.EXPECT().CreatePermission(gomock.Any(), &admin.CreatePermissionRequest{
					PermissionId: "foo",
				}).Return(&v1alpha.Permission{
					Name: "permissions/foo",
				}, nil)
			},
			Check: test.Succeeded(&v1alpha.Permission{
				Name: "permissions/foo",
			}),
		},
	}

	xs.Run(t)
}
