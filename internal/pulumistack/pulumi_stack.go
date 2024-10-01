package pulumistack

import (
	"buf.build/gen/go/plantoncloud/project-planton/protocolbuffers/go/project/planton/shared/pulumi"
	"github.com/pkg/errors"
	"github.com/plantoncloud/project-planton/internal/pulumimodule"
	"github.com/plantoncloud/project-planton/internal/stackinput"
	cliworkspace "github.com/plantoncloud/project-planton/internal/workspace"
	"os"
	"os/exec"
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

func Run(stackFqdn, targetManifestPath, kubernetesClusterManifestPath string,
	pulumiOperation pulumi.PulumiOperationType, isUpdatePreview bool) error {
	kindName, err := stackinput.ExtractKindFromTargetManifest(targetManifestPath)
	if err != nil {
		return errors.Wrapf(err, "failed to extract kind from %s stack input yaml", targetManifestPath)
	}

	cloneUrl, err := pulumimodule.GetCloneUrl(kindName)
	if err != nil {
		return errors.Wrapf(err, "failed to get clone url for %s kind", kindName)
	}

	stackWorkspaceDir, err := cliworkspace.GetWorkspaceDir(stackFqdn)
	if err != nil {
		return errors.Wrapf(err, "failed to get %s stack worspace directory", stackFqdn)
	}

	pulumiProjectName, err := ExtractProjectName(stackFqdn)
	if err != nil {
		return errors.Wrapf(err, "failed to extract project name from %s stack fqdn", stackFqdn)
	}

	stackInputYamlContent, err := stackinput.BuildStackInputYaml(targetManifestPath, kubernetesClusterManifestPath)
	if err != nil {
		return errors.Wrap(err, "failed to build stack input yaml")
	}

	gitRepoName, err := extractGitRepoName(cloneUrl)
	if err != nil {
		return errors.Wrapf(err, "failed to extract git repo name from %s", cloneUrl)
	}

	// Check if the cloned repository directory already exists
	pulumiModuleRepoPath := stackWorkspaceDir + "/" + gitRepoName
	if _, err := os.Stat(pulumiModuleRepoPath); os.IsNotExist(err) {
		gitCloneCommand := exec.Command("git", "clone", cloneUrl, stackWorkspaceDir)
		if err := gitCloneCommand.Run(); err != nil {
			return errors.Wrapf(err, "failed to clone repository from %s to %s", cloneUrl, stackWorkspaceDir)
		}
	}
	if err := updateProjectNameInPulumiYaml(pulumiModuleRepoPath, pulumiProjectName); err != nil {
		return errors.Wrapf(err, "failed to update project name in %s/Pulumi.yaml", pulumiModuleRepoPath)
	}

	op := pulumiOperation.String()
	if isUpdatePreview {
		op = "preview"
	}

	pulumiCmd := exec.Command("pulumi", op, "--stack", stackFqdn, "--yes")

	// Set the STACK_INPUT_YAML environment variable
	pulumiCmd.Env = append(os.Environ(), "STACK_INPUT_YAML="+stackInputYamlContent)

	// Set the working directory to the repository path
	pulumiCmd.Dir = pulumiModuleRepoPath

	pulumiCmd.Stdout = os.Stdout
	pulumiCmd.Stderr = os.Stderr

	if err := pulumiCmd.Run(); err != nil {
		return errors.Wrapf(err, "failed to execute pulumi command %s", op)
	}

	return nil
}

// extractGitRepoName takes a repository URL and returns the repository name.
func extractGitRepoName(repoUrl string) (string, error) {
	parts := strings.Split(repoUrl, "/")
	if len(parts) < 1 {
		return "", errors.New("invalid repository URL format, expected format <domain>/<user>/<repo>.git")
	}
	repoNameWithGit := parts[len(parts)-1]
	repoName := strings.TrimSuffix(repoNameWithGit, ".git")
	return repoName, nil
}

func updateProjectNameInPulumiYaml(pulumiModuleRepoPath, pulumiProjectName string) error {
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
