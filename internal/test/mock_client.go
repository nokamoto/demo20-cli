package test

import (
	compute "github.com/nokamoto/demo20-apis/cloud/compute/v1alpha"
	admin "github.com/nokamoto/demo20-apis/cloud/iam/admin/v1alpha"
	rdb "github.com/nokamoto/demo20-apis/cloud/rdb/v1alpha"
	resourcemanager "github.com/nokamoto/demo20-apis/cloud/resourcemanager/v1alpha"
	"github.com/nokamoto/demo20-cli/internal/test/mock/iam/mockadmin"
	"github.com/nokamoto/demo20-cli/internal/test/mock/mockcompute"
	"github.com/nokamoto/demo20-cli/internal/test/mock/mockrdb"
	"github.com/nokamoto/demo20-cli/internal/test/mock/mockresourcemanager"
)

// MockClient represents mocked gRPC clients.
type MockClient struct {
	MockCompute         *mockcompute.MockComputeClient
	MockAdmin           *mockadmin.MockIamClient
	MockResourceManager *mockresourcemanager.MockResourceManagerClient
	MockRdb             *mockrdb.MockRdbClient
}

// Compute returns a mocked compute client.
func (c *MockClient) Compute() compute.ComputeClient {
	return c.MockCompute
}

// AdminIam returns a mocked iam admin client.
func (c *MockClient) AdminIam() admin.IamClient {
	return c.MockAdmin
}

// ResourceManager returns a mocked resourcemanager client.
func (c *MockClient) ResourceManager() resourcemanager.ResourceManagerClient {
	return c.MockResourceManager
}

// Rdb returns a mocked rdb client.
func (c *MockClient) Rdb() rdb.RdbClient {
	return c.MockRdb
}
