package stackinput

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/plantoncloud/project-planton/internal/stackinput/credentials"
	"gopkg.in/yaml.v3"
	"os"
)

// BuildStackInputYaml reads two YAML files, combines their contents,
// and returns a new YAML string with "target" and all the credential keys.
func BuildStackInputYaml(targetManifestPath string, stackInputOptions credentials.StackInputCredentialOptions) (string, error) {
	targetContent, err := os.ReadFile(targetManifestPath)
	if err != nil {
		return "", fmt.Errorf("failed to read target manifest file: %w", err)
	}

	stackInputContentMap := map[string]string{
		"target": string(targetContent),
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

// ExtractKindFromTargetManifest reads a YAML file from the given path and returns the value of the 'kind' key.
func ExtractKindFromTargetManifest(targetManifest string) (string, error) {
	// Check if the file exists
	if _, err := os.Stat(targetManifest); os.IsNotExist(err) {
		return "", errors.Wrapf(err, "file not found: %s", targetManifest)
	}

	// Read the YAML file
	fileContent, err := os.ReadFile(targetManifest)
	if err != nil {
		return "", errors.Wrapf(err, "failed to read file: %s", targetManifest)
	}

	// Parse the YAML content
	var yamlData map[string]interface{}
	if err := yaml.Unmarshal(fileContent, &yamlData); err != nil {
		return "", errors.Wrapf(err, "failed to unmarshal YAML content from file: %s", targetManifest)
	}

	// Extract the 'kind' key
	kind, ok := yamlData["kind"]
	if !ok {
		return "", errors.Errorf("key 'kind' not found in YAML file: %s", targetManifest)
	}

	// Ensure the 'kind' key is a string
	kindStr, ok := kind.(string)
	if !ok {
		return "", errors.Errorf("value of 'kind' key is not a string in YAML file: %s", targetManifest)
	}

	return kindStr, nil
}
