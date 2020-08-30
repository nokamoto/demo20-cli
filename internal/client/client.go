package client

import (
	compute "github.com/nokamoto/demo20-apis/cloud/compute/v1alpha"
	admin "github.com/nokamoto/demo20-apis/cloud/iam/admin/v1alpha"
	rdb "github.com/nokamoto/demo20-apis/cloud/rdb/v1alpha"
	resourcemanager "github.com/nokamoto/demo20-apis/cloud/resourcemanager/v1alpha"
	"github.com/nokamoto/demo20-cli/internal/config"
	"google.golang.org/grpc"
)

// Client represents gRPC clients.
type Client interface {
	Compute() compute.ComputeClient
	AdminIam() admin.IamClient
	ResourceManager() resourcemanager.ResourceManagerClient
	Rdb() rdb.RdbClient
}

type client struct {
	value *config.Value
}

// NewClient returns Client accesses to the gRPC address.
func NewClient(value *config.Value) Client {
	return &client{value: value}
}

func (c *client) con() *grpc.ClientConn {
	con, err := grpc.Dial(c.value.GrpcAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return con
}

// Compute returns ComputeClient.
func (c *client) Compute() compute.ComputeClient {
	return compute.NewComputeClient(c.con())
}

// AdminIam returns IamClient.
func (c *client) AdminIam() admin.IamClient {
	return admin.NewIamClient(c.con())
}

// ResourceManager returns ResourceManagerClient.
func (c *client) ResourceManager() resourcemanager.ResourceManagerClient {
	return resourcemanager.NewResourceManagerClient(c.con())
}

// Rdb returns RdbClient.
func (c *client) Rdb() rdb.RdbClient {
	return rdb.NewRdbClient(c.con())
}
