package test

import (
	compute "github.com/nokamoto/demo20-apis/cloud/compute/v1alpha"
	admin "github.com/nokamoto/demo20-apis/cloud/iam/admin/v1alpha"
	"github.com/nokamoto/demo20-cli/internal/test/mock/iam/mockadmin"
	"github.com/nokamoto/demo20-cli/internal/test/mock/mockcompute"
)

// MockClient represents mocked gRPC clients.
type MockClient struct {
	MockCompute *mockcompute.MockComputeClient
	MockAdmin   *mockadmin.MockIamClient
}

// Compute returns a mocked compute client.
func (c *MockClient) Compute() compute.ComputeClient {
	return c.MockCompute
}

// AdminIam returns a mocked iam admin client.
func (c *MockClient) AdminIam() admin.IamClient {
	return c.MockAdmin
}
