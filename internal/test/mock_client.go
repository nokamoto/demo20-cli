package test

import (
	compute "github.com/nokamoto/demo20-apis/cloud/compute/v1alpha"
	mockcompute "github.com/nokamoto/demo20-cli/internal/test/mock/compute"
)

// MockClient represents mocked gRPC clients.
type MockClient struct {
	MockCompute *mockcompute.MockComputeClient
}

// Compute creates a new mock compute client.
func (c *MockClient) Compute() compute.ComputeClient {
	return c.MockCompute
}
