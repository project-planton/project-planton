package manifest

import (
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"sigs.k8s.io/yaml"
)

func Print(input proto.Message) error {
	marshalJsonBytes, err := protojson.Marshal(input)
	if err != nil {
		return errors.Wrap(err, "failed to yaml marshalJsonBytes")
	}
	marshalYamlBytes, err := yaml.JSONToYAML(marshalJsonBytes)
	if err != nil {
		return errors.Wrap(err, "failed to marshal json to yaml")
	}

	fmt.Printf("%v\n", string(marshalYamlBytes))
	return nil
}
