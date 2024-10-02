package pulumimodule

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/project-planton/internal/fileutil"
	"github.com/plantoncloud/project-planton/internal/stackinput"
	"github.com/plantoncloud/project-planton/internal/workspace"
	"os"
	"os/exec"
)

func GetPath(moduleDir string, stackFqdn, targetManifestPath string) (string, error) {
	isPulumiModuleDir, err := IsPulumiModuleDirectory(moduleDir)
	if err != nil {
		return "", errors.Wrapf(err, "failed to check if %s is a valid pulumi module directory", moduleDir)
	}
	if isPulumiModuleDir {
		return moduleDir, nil
	}

	kindName, err := stackinput.ExtractKindFromTargetManifest(targetManifestPath)
	if err != nil {
		return "", errors.Wrapf(err, "failed to extract kind from %s stack input yaml", targetManifestPath)
	}

	stackWorkspaceDir, err := workspace.GetWorkspaceDir(stackFqdn)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get %s stack worspace directory", stackFqdn)
	}

	cloneUrl, err := GetCloneUrl(kindName)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get clone url for %s kind", kindName)
	}

	gitRepoName, err := extractGitRepoName(cloneUrl)
	if err != nil {
		return "", errors.Wrapf(err, "failed to extract git repo name from %s", cloneUrl)
	}

	// Check if the cloned repository directory already exists
	pulumiModuleRepoPath := stackWorkspaceDir + "/" + gitRepoName

	if _, err := os.Stat(pulumiModuleRepoPath); os.IsNotExist(err) {
		gitCloneCommand := exec.Command("git", "clone", cloneUrl, pulumiModuleRepoPath)
		gitCloneCommand.Stdout = os.Stdout
		gitCloneCommand.Stderr = os.Stderr
		if err := gitCloneCommand.Run(); err != nil {
			return "", errors.Wrapf(err, "failed to clone repository from %s to %s", cloneUrl, stackWorkspaceDir)
		}
	}
	return pulumiModuleRepoPath, nil
}

// IsPulumiModuleDirectory checks if the given directory contains a Pulumi.yaml file.
// It returns true if the file exists, false otherwise. If an error occurs during the check, it returns an error.
func IsPulumiModuleDirectory(moduleDir string) (bool, error) {
	pulumiYamlPath := moduleDir + "/Pulumi.yaml"
	isExists, err := fileutil.IsExists(pulumiYamlPath)
	if err != nil {
		return false, errors.Wrapf(err, "failed to check if %s exists", pulumiYamlPath)
	}
	return isExists, nil
}
