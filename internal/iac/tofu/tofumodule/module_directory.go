package tofumodule

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/internal/cli/version"
	"github.com/project-planton/project-planton/internal/cli/workspace"
	"github.com/project-planton/project-planton/internal/iac/gitrepo"
	"github.com/project-planton/project-planton/internal/provider"
	"github.com/project-planton/project-planton/pkg/fileutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func getModulePath(moduleDir, kindName string) (string, error) {
	isTerraformModuleDir, err := isTerraformModuleDirectory(moduleDir)
	if err != nil {
		return "", errors.Wrapf(err, "failed to check if %s is a valid terraform module directory", moduleDir)
	}
	if isTerraformModuleDir {
		return moduleDir, nil
	}

	tofuModuleWorkspaceDir, err := getWorkspaceDir()
	if err != nil {
		return "", errors.Wrapf(err, "failed to get %s stack workspace directory")
	}

	gitRepoName, err := gitrepo.ExtractRepoName(gitrepo.CloneUrl)
	if err != nil {
		return "", errors.Wrapf(err, "failed to extract git repo name from %s", gitrepo.CloneUrl)
	}

	// Check if the cloned repository directory already exists
	terraformModuleRepoPath := filepath.Join(tofuModuleWorkspaceDir, gitRepoName)

	if _, statErr := os.Stat(terraformModuleRepoPath); os.IsNotExist(statErr) {
		gitCloneCommand := exec.Command("git", "clone", gitrepo.CloneUrl, terraformModuleRepoPath)
		gitCloneCommand.Stdout = os.Stdout
		gitCloneCommand.Stderr = os.Stderr
		if err := gitCloneCommand.Run(); err != nil {
			return "", errors.Wrapf(err, "failed to clone repository from %s to %s", gitrepo.CloneUrl, tofuModuleWorkspaceDir)
		}
	}

	//checkout the project-planton version tag if it is not the default version
	if version.Version != version.DefaultVersion {
		gitCheckoutCommand := exec.Command("git", "-C", terraformModuleRepoPath, "checkout", version.Version)
		gitCheckoutCommand.Stdout = os.Stdout
		gitCheckoutCommand.Stderr = os.Stderr
		if err := gitCheckoutCommand.Run(); err != nil {
			return "", errors.Wrapf(err, "failed to checkout tag %s in %s", version.Version, terraformModuleRepoPath)
		}
	}

	terraformModulePath, err := getTerraformModulePath(terraformModuleRepoPath, kindName)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get terraform module path for %s", kindName)
	}

	return terraformModulePath, nil
}

// isTerraformModuleDirectory checks if the given directory contains any files with .tf extension.
// It returns true if any .tf files exists, false otherwise. If an error occurs during the check, it returns an error.
func isTerraformModuleDirectory(moduleDir string) (bool, error) {
	pulumiYamlPath := moduleDir + "/Pulumi.yaml"
	isExists, err := fileutil.IsExists(pulumiYamlPath)
	if err != nil {
		return false, errors.Wrapf(err, "failed to check if %s exists", pulumiYamlPath)
	}
	return isExists, nil
}

func getTerraformModulePath(moduleRepoDir, kindName string) (string, error) {
	kindProvider := provider.GetProvider(provider.KindName(kindName))
	if kindProvider == shared.KindProvider_kind_provider_unspecified {
		return "", errors.New("failed to get kind provider")
	}

	terraformModulePath := filepath.Join(moduleRepoDir, "apis/project/planton/provider",
		kindProvider.String(), strings.ToLower(kindName), "v1/iac/tofu")

	if _, err := os.Stat(terraformModulePath); os.IsNotExist(err) {
		return "", errors.Wrapf(err, "failed to get %s module directory", kindName)
	}

	return terraformModulePath, nil
}

// getWorkspaceDir returns the path of the workspace directory to which terraform module repo can be cloned.
func getWorkspaceDir() (string, error) {
	cliWorkspaceDir, err := workspace.GetWorkspaceDir()
	if err != nil {
		return "", errors.Wrapf(err, "failed to get cli workspace directory")
	}
	//base directory will always be ${HOME}/.planton-cloud/pulumi
	tofuModuleWorkspaceDir := filepath.Join(cliWorkspaceDir, "tofu")
	if !fileutil.IsDirExists(tofuModuleWorkspaceDir) {
		if err := os.MkdirAll(tofuModuleWorkspaceDir, os.ModePerm); err != nil {
			return "", errors.Wrapf(err, "failed to ensure %s dir", tofuModuleWorkspaceDir)
		}
	}
	return tofuModuleWorkspaceDir, nil
}
