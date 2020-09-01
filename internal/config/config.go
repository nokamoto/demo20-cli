package config

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/nokamoto/demo20-apis/cloud/api"
	"github.com/nokamoto/demo20-apps/pkg/sdk/metadata"
	"gopkg.in/yaml.v2"
)

const (
	// Filename is a configration file name.
	Filename = ".cloud"
	// Extension is a configration file extension.
	Extension = "yaml"
)

// Value represetns a commandline configuration.
type Value struct {
	GrpcAddress       string `yaml:"grpcAddress"`
	ProjectID         string `yaml:"projectID"`
	MachineUserAPIKey string `yaml:"machineUserAPIKey"`
}

var (
	// Default represents a default configration.
	Default Value
)

// View returns a YAML string of Value.
func (v Value) View() string {
	bytes, _ := yaml.Marshal(&v)
	return string(bytes)
}

// Write writes a YAML string to the file.
func (v Value) Write(filename string) error {
	bytes, err := yaml.Marshal(&v)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, bytes, 0644)
}

// OutgoingContext returns a new outgoing context.
func (v Value) OutgoingContext() context.Context {
	ctx, err := metadata.AppendToOutgoingContext(context.Background(), &api.Metadata{
		Credential: &api.Metadata_MachineUserApiKey{
			MachineUserApiKey: v.MachineUserAPIKey,
		},
		Parent: fmt.Sprintf("projects/%s", v.ProjectID),
	})
	if err != nil {
		panic(err)
	}
	return ctx
}
