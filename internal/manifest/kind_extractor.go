package manifest

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"os"
	"regexp"
	"strings"
)

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

func ConvertKindName(kindName string) string {
	// This uses a Regex to find places where there is an uppercase letter
	// that is followed by a lowercase letter and separates the words using a hyphen
	re := regexp.MustCompile("([a-z])([A-Z])")
	// Replace the matches found by the regex with a hyphen and the matched uppercase letter in lowercase
	formattedName := re.ReplaceAllStringFunc(kindName, func(match string) string {
		return match[:1] + "-" + strings.ToLower(match[1:])
	})
	// Convert the final string to lowercase and return it
	return strings.ToLower(formattedName)
}
