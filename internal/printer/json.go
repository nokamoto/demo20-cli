package printer

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/spf13/cobra"
)

// Proto prints the proto message using json mapping.
func Proto(cmd *cobra.Command, value proto.Message) error {
	m := jsonpb.Marshaler{}
	s, err := m.MarshalToString(value)
	if err != nil {
		return err
	}

	cmd.Println(s)

	return nil
}
