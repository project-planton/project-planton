package builder

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"sigs.k8s.io/kustomize/api/krusty"
	"sigs.k8s.io/kustomize/kyaml/filesys"
)

// BuildManifest builds a kustomize manifest from the specified directory and overlay.
// It writes the result to a temporary file and returns the path to that file.
// The caller is responsible for cleaning up the temp file using Cleanup() or os.Remove().
//
// Parameters:
//   - kustomizeDir: Base directory containing kustomize structure (e.g., ./backend/services/secrets-manager/_kustomize)
//   - overlay: The overlay environment to build (e.g., "prod", "dev", "staging")
//
// Returns:
//   - tempFilePath: Path to the temporary manifest file
//   - error: Any error encountered during the build process
func BuildManifest(kustomizeDir, overlay string) (string, error) {
	// Validate inputs
	if kustomizeDir == "" {
		return "", errors.New("kustomize-dir cannot be empty")
	}
	if overlay == "" {
		return "", errors.New("overlay cannot be empty")
	}

	// Construct the full kustomization path: kustomizeDir/overlays/overlay
	kustomizationPath := filepath.Join(kustomizeDir, "overlays", overlay)

	// Verify the kustomization directory exists
	if _, err := os.Stat(kustomizationPath); os.IsNotExist(err) {
		return "", errors.Errorf("kustomization path does not exist: %s", kustomizationPath)
	}

	// Create kustomize options
	opts := krusty.MakeDefaultOptions()

	// Create kustomizer
	k := krusty.MakeKustomizer(opts)

	// Create filesystem
	fSys := filesys.MakeFsOnDisk()

	// Run kustomize build
	resMap, err := k.Run(fSys, kustomizationPath)
	if err != nil {
		return "", errors.Wrapf(err, "failed to run kustomize build for path: %s", kustomizationPath)
	}

	// Convert resource map to YAML
	yaml, err := resMap.AsYaml()
	if err != nil {
		return "", errors.Wrap(err, "failed to convert kustomize result to YAML")
	}

	// Create temp file
	tempFile, err := os.CreateTemp("", "kustomize-manifest-*.yaml")
	if err != nil {
		return "", errors.Wrap(err, "failed to create temporary file")
	}
	defer tempFile.Close()

	// Write YAML to temp file
	if _, err := tempFile.Write(yaml); err != nil {
		os.Remove(tempFile.Name()) // Clean up on write failure
		return "", errors.Wrap(err, "failed to write manifest to temporary file")
	}

	return tempFile.Name(), nil
}

// Cleanup removes the temporary manifest file created by BuildManifest.
// This is a convenience function; callers can also use os.Remove directly.
func Cleanup(tempFilePath string) error {
	if tempFilePath == "" {
		return nil
	}
	if err := os.Remove(tempFilePath); err != nil && !os.IsNotExist(err) {
		return errors.Wrapf(err, "failed to remove temporary file: %s", tempFilePath)
	}
	return nil
}
