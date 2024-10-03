package manifestyaml

import (
	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"os"
	"sigs.k8s.io/yaml"
)

func LoadManifest(manifestPath string) (proto.Message, error) {
	manifestYamlBytes, err := os.ReadFile(manifestPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read manifest file")
	}
	jsonBytes, err := yaml.YAMLToJSON(manifestYamlBytes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load yaml to json")
	}

	kindName, err := ExtractKindFromTargetManifest(manifestPath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to extract kind from %s stack input yaml", manifestPath)
	}

	manifest := DeploymentComponentMap[DeploymentComponent(ConvertKindName(kindName))]

	if manifest == nil {
		return nil, errors.Errorf("deployment-component does not contain %s", ConvertKindName(kindName))
	}

	if err := protojson.Unmarshal(jsonBytes, manifest); err != nil {
		return nil, errors.Wrap(err, "failed to load json into proto message")
	}
	return manifest, nil
}
