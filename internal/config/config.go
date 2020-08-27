package config

import (
	"io/ioutil"

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
	GrpcAddress string `yaml:"grpcAddress"`
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
