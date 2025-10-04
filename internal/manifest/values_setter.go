package manifest

import (
	"os"

	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/internal/manifest/manifestprotobuf"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"sigs.k8s.io/yaml"
)

func LoadWithOverrides(manifestPath string, valueOverrides map[string]string) (proto.Message, error) {
	manifest, err := LoadManifest(manifestPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load manifest")
	}
	for key, value := range valueOverrides {
		manifest, err = manifestprotobuf.SetProtoField(manifest, key, value)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to set %s=%s", key, value)
		}
	}
	return manifest, nil
}

// ApplyOverridesToFile loads a manifest, applies overrides, and writes to a new temp file.
// Returns the path to the new temp file and whether it's a temp file that needs cleanup.
// If no overrides are provided, returns the original path unchanged.
func ApplyOverridesToFile(manifestPath string, valueOverrides map[string]string) (string, bool, error) {
	// If no overrides, return original path
	if len(valueOverrides) == 0 {
		return manifestPath, false, nil
	}

	// Load manifest with overrides applied
	manifest, err := LoadWithOverrides(manifestPath, valueOverrides)
	if err != nil {
		return "", false, errors.Wrap(err, "failed to load manifest with overrides")
	}

	// Convert to YAML
	jsonBytes, err := protojson.Marshal(manifest)
	if err != nil {
		return "", false, errors.Wrap(err, "failed to marshal manifest to json")
	}

	yamlBytes, err := yaml.JSONToYAML(jsonBytes)
	if err != nil {
		return "", false, errors.Wrap(err, "failed to convert json to yaml")
	}

	// Write to temp file
	tempFile, err := os.CreateTemp("", "manifest-with-overrides-*.yaml")
	if err != nil {
		return "", false, errors.Wrap(err, "failed to create temp file for overrides")
	}
	defer tempFile.Close()

	if _, err := tempFile.Write(yamlBytes); err != nil {
		os.Remove(tempFile.Name())
		return "", false, errors.Wrap(err, "failed to write manifest with overrides")
	}

	return tempFile.Name(), true, nil
}
