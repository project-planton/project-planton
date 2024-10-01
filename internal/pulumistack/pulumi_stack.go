package pulumistack

import (
	"buf.build/gen/go/plantoncloud/project-planton/protocolbuffers/go/project/planton/shared/pulumi"
	"context"
	"github.com/pkg/errors"
	"github.com/plantoncloud/project-planton/internal/pulumimodule"
	cliworkspace "github.com/plantoncloud/project-planton/internal/workspace"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/workspace"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strings"
)

// ExtractProjectName extracts the project name from the stack FQDN.
func ExtractProjectName(stackFqdn string) (string, error) {
	parts := strings.Split(stackFqdn, "/")
	if len(parts) != 3 {
		return "", errors.New("invalid stack fqdn format, expected format <org>/<project>/<stack>")
	}
	return parts[1], nil
}

func Run(stackFqdn, targetManifestPath string, pulumiOperation pulumi.PulumiOperationType, isUpdatePreview bool) error {
	kindName, err := extractKindFromYaml(targetManifestPath)
	if err != nil {
		return errors.Wrapf(err, "failed to extract kind from %s stack input yaml", targetManifestPath)
	}

	cloneUrl, err := pulumimodule.GetCloneUrl(kindName)
	if err != nil {
		return errors.Wrapf(err, "failed to get clone url for %s kind", kindName)
	}

	gitRepo := auto.GitRepo{
		URL:     cloneUrl,
		Shallow: false,
	}

	stackWorkspaceDir, err := cliworkspace.GetWorkspaceDir(stackFqdn)
	if err != nil {
		return errors.Wrapf(err, "failed to get %s stack worspace directory", stackFqdn)
	}

	pulumiProjectName, err := ExtractProjectName(stackFqdn)
	if err != nil {
		return errors.Wrapf(err, "failed to extract project name from %s stack fqdn", stackFqdn)
	}

	//setup pulumi automation-api stack
	pulumiStack, err := auto.UpsertStackLocalSource(
		context.Background(),
		stackFqdn,
		stackWorkspaceDir,
		auto.Repo(gitRepo),
		auto.Project(workspace.Project{
			Name: tokens.PackageName(pulumiProjectName),
			Runtime: workspace.NewProjectRuntimeInfo(
				pulumi.PulumiProjectRuntime_go.String(),
				map[string]interface{}{}),
		}),
		auto.WorkDir(stackWorkspaceDir),
		auto.EnvVars(map[string]string{
			"STACK_INPUT": targetManifestPath,
		}))
	if err != nil {
		return errors.Wrapf(err, "failed to setup pulumi automation-api stack %s", stackFqdn)
	}

	switch pulumiOperation {
	case pulumi.PulumiOperationType_refresh:
		if _, err := pulumiStack.Refresh(context.Background()); err != nil {
			return errors.Wrapf(err, "failed to refresh %s stack", stackFqdn)
		}
	case pulumi.PulumiOperationType_update:
		if isUpdatePreview {
			if _, err := pulumiStack.Preview(context.Background()); err != nil {
				return errors.Wrapf(err, "failed to preview %s stack", stackFqdn)
			}
		} else {
			if _, err := pulumiStack.Up(context.Background()); err != nil {
				return errors.Wrapf(err, "failed to update %s stack", stackFqdn)
			}
		}
	case pulumi.PulumiOperationType_destroy:
		if _, err := pulumiStack.Destroy(context.Background()); err != nil {
			return errors.Wrapf(err, "failed to destroy %s stack", stackFqdn)
		}
	}
	return nil
}

// extractKindFromYaml reads a YAML file from the given path and returns the value of the 'kind' key.
func extractKindFromYaml(yamlPath string) (string, error) {
	// Check if the file exists
	if _, err := os.Stat(yamlPath); os.IsNotExist(err) {
		return "", errors.Wrapf(err, "file not found: %s", yamlPath)
	}

	// Read the YAML file
	fileContent, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		return "", errors.Wrapf(err, "failed to read file: %s", yamlPath)
	}

	// Parse the YAML content
	var yamlData map[string]interface{}
	if err := yaml.Unmarshal(fileContent, &yamlData); err != nil {
		return "", errors.Wrapf(err, "failed to unmarshal YAML content from file: %s", yamlPath)
	}

	// Extract the 'kind' key
	kind, ok := yamlData["kind"]
	if !ok {
		return "", errors.Errorf("key 'kind' not found in YAML file: %s", yamlPath)
	}

	// Ensure the 'kind' key is a string
	kindStr, ok := kind.(string)
	if !ok {
		return "", errors.Errorf("value of 'kind' key is not a string in YAML file: %s", yamlPath)
	}

	return kindStr, nil
}
