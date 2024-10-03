package stackinput

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/project-planton/internal/manifestyaml"
	"github.com/plantoncloud/project-planton/internal/stackinput/credentials"
	"google.golang.org/protobuf/encoding/protojson"
	"gopkg.in/yaml.v3"
)

// BuildStackInputYaml reads two YAML files, combines their contents,
// and returns a new YAML string with "target" and all the credential keys.
func BuildStackInputYaml(targetManifestPath string, valueOverrides map[string]string,
	stackInputOptions credentials.StackInputCredentialOptions) (string, error) {
	manifestObject, err := manifestyaml.OverrideValues(targetManifestPath, valueOverrides)
	if err != nil {
		return "", errors.Wrapf(err, "failed to override values in target manifest file")
	}

	var targetContentMap map[string]interface{}
	targetContent, err := protojson.Marshal(manifestObject)
	if err != nil {
		return "", errors.Wrap(err, "failed to marshal manifest object to JSON")
	}

	err = yaml.Unmarshal(targetContent, &targetContentMap)
	if err != nil {
		return "", errors.Wrapf(err, "failed to unmarshal target manifest file")
	}

	stackInputContentMap := map[string]interface{}{
		"target": targetContentMap,
	}

	stackInputContentMap, err = addCredentials(stackInputContentMap, stackInputOptions)
	if err != nil {
		return "", errors.Wrapf(err, "failed to add credentials to stack-input yaml")
	}

	finalStackInputYaml, err := yaml.Marshal(stackInputContentMap)
	if err != nil {
		return "", errors.Wrap(err, "failed to marshal final stack-input yaml")
	}
	return string(finalStackInputYaml), nil
}
