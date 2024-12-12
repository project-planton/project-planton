package pulumimodule

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/internal/cli/version"
	"github.com/project-planton/project-planton/internal/cli/workspace"
	"github.com/project-planton/project-planton/internal/deploymentcomponent"
	"github.com/project-planton/project-planton/internal/iac/gitrepo"
	"github.com/project-planton/project-planton/pkg/fileutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetPath(moduleDir string, stackFqdn, kindName string) (string, error) {
	isPulumiModuleDir, err := IsPulumiModuleDirectory(moduleDir)
	if err != nil {
		return "", errors.Wrapf(err, "failed to check if %s is a valid pulumi module directory", moduleDir)
	}
	if isPulumiModuleDir {
		return moduleDir, nil
	}

	stackWorkspaceDir, err := workspace.GetWorkspaceDir(stackFqdn)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get %s stack workspace directory", stackFqdn)
	}

	gitRepoName, err := extractGitRepoName(gitrepo.CloneUrl)
	if err != nil {
		return "", errors.Wrapf(err, "failed to extract git repo name from %s", gitrepo.CloneUrl)
	}

	// Check if the cloned repository directory already exists
	pulumiModuleRepoPath := filepath.Join(stackWorkspaceDir, gitRepoName)

	if _, statErr := os.Stat(pulumiModuleRepoPath); os.IsNotExist(statErr) {
		gitCloneCommand := exec.Command("git", "clone", gitrepo.CloneUrl, pulumiModuleRepoPath)
		gitCloneCommand.Stdout = os.Stdout
		gitCloneCommand.Stderr = os.Stderr
		if err := gitCloneCommand.Run(); err != nil {
			return "", errors.Wrapf(err, "failed to clone repository from %s to %s", gitrepo.CloneUrl, stackWorkspaceDir)
		}
	}

	//checkout the project-planton version tag if it is not the default version
	if version.Version != version.DefaultVersion {
		gitCheckoutCommand := exec.Command("git", "-C", pulumiModuleRepoPath, "checkout", version.Version)
		gitCheckoutCommand.Stdout = os.Stdout
		gitCheckoutCommand.Stderr = os.Stderr
		if err := gitCheckoutCommand.Run(); err != nil {
			return "", errors.Wrapf(err, "failed to checkout tag %s in %s", version.Version, pulumiModuleRepoPath)
		}
	}

	pulumiModulePath, err := getPulumiModulePath(pulumiModuleRepoPath, kindName)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get pulumi module path for %s", kindName)
	}

	return pulumiModulePath, nil
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

func getPulumiModulePath(moduleRepoDir, kindName string) (string, error) {
	kindProvider := deploymentcomponent.GetProvider(deploymentcomponent.KindName(kindName))
	if kindProvider == shared.KindProvider_kind_provider_unspecified {
		return "", errors.New("failed to get kind provider")
	}

	pulumiModulePath := filepath.Join(moduleRepoDir, "apis/project/planton/provider",
		kindProvider.String(), strings.ToLower(kindName), "v1/iac/pulumi")

	if _, err := os.Stat(pulumiModulePath); os.IsNotExist(err) {
		return "", errors.Wrapf(err, "failed to get %s module directory", kindName)
	}

	return pulumiModulePath, nil
}
