package pulumistack

import (
	"os"
	"strings"

	"github.com/pkg/errors"
)

// ExtractProjectName extracts the project name from the stack FQDN.
func ExtractProjectName(stackFqdn string) (string, error) {
	parts := strings.Split(stackFqdn, "/")
	if len(parts) != 3 {
		return "", errors.New("invalid stack fqdn format, expected format <org>/<project>/<stack>")
	}
	return parts[1], nil
}

// UpdateProjectNameInPulumiYaml updates the project name in Pulumi.yaml file
// to match the project name from the stack FQDN.
func UpdateProjectNameInPulumiYaml(pulumiModuleRepoPath, pulumiProjectName string) error {
	// Check if the cloned repository contains Pulumi.yaml file
	pulumiYamlPath := pulumiModuleRepoPath + "/Pulumi.yaml"
	if _, err := os.Stat(pulumiYamlPath); os.IsNotExist(err) {
		return errors.Errorf("Pulumi.yaml file is missing in the repository at %s", pulumiModuleRepoPath)
	}

	// Update the Pulumi.yaml file with the new project name
	pulumiYamlContent, err := os.ReadFile(pulumiYamlPath)
	if err != nil {
		return errors.Wrapf(err, "failed to read Pulumi.yaml from %s", pulumiYamlPath)
	}

	lines := strings.Split(string(pulumiYamlContent), "\n")
	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "name:") {
			lines[i] = "name: " + pulumiProjectName
			break
		}
	}

	updatedYamlContent := strings.Join(lines, "\n")
	if err := os.WriteFile(pulumiYamlPath, []byte(updatedYamlContent), 0644); err != nil {
		return errors.Wrapf(err, "failed to write updated Pulumi.yaml to %s", pulumiYamlPath)
	}
	return nil
}
