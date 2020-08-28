package printer

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/spf13/cobra"
)

// ProtoString returns a string of the proto message using json mapping.
func ProtoString(value proto.Message) (string, error) {
	m := jsonpb.Marshaler{}
	return m.MarshalToString(value)
}

// Proto prints the proto message using json mapping.
func Proto(cmd *cobra.Command, value proto.Message) error {
	s, err := ProtoString(value)
	if err != nil {
		return err
	}

	cmd.Println(s)

	return nil
}
