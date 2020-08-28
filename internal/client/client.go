package client

import (
	compute "github.com/nokamoto/demo20-apis/cloud/compute/v1alpha"
	"github.com/nokamoto/demo20-cli/internal/config"
	"google.golang.org/grpc"
)

// Client represents gRPC clients.
type Client struct {
	value *config.Value
}

// NewClient returns Client accesses to the gRPC address.
func NewClient(value *config.Value) *Client {
	return &Client{value: value}
}

func (d *Client) con() *grpc.ClientConn {
	c, err := grpc.Dial(d.value.GrpcAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return c
}

// Compute returns ComputeClient.
func (d *Client) Compute() compute.ComputeClient {
	return compute.NewComputeClient(d.con())
}
