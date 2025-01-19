package protobufyaml

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/fileutil"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"os"
	"sigs.k8s.io/yaml"
)

// Load reads the contents of the input file and load it into an object
// the bytes from the input files are first converted to json
// the converted json is then converted into the protobuf message using protojson package.
func Load(inputFile string, obj proto.Message) error {
	isExists, err := fileutil.IsExists(inputFile)
	if err != nil {
		return errors.Wrapf(err, "failed to check file: %s", inputFile)
	}
	if !isExists {
		return errors.New("file does not exist")
	}
	inputFileBytes, err := os.ReadFile(inputFile)
	if err != nil {
		return errors.Wrap(err, "failed to read input file")
	}
	jsonBytes, err := yaml.YAMLToJSON(inputFileBytes)
	if err != nil {
		return errors.Wrap(err, "failed to load yaml to json")
	}
	if err := protojson.Unmarshal(jsonBytes, obj); err != nil {
		return errors.Wrap(err, "failed to load json into proto message")
	}
	return nil
}

func LoadYamlBytes(yamlBytes []byte, obj proto.Message) error {
	jsonBytes, err := yaml.YAMLToJSON(yamlBytes)
	if err != nil {
		return errors.Wrap(err, "failed to load yaml to json")
	}
	if err := protojson.Unmarshal(jsonBytes, obj); err != nil {
		return errors.Wrap(err, "failed to load json into proto message")
	}
	return nil
}
