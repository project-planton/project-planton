package tofumodule

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/cloudresourcekind"
	"github.com/project-planton/project-planton/internal/cli/version"
	"github.com/project-planton/project-planton/internal/cli/workspace"
	"github.com/project-planton/project-planton/pkg/crkreflect"
	"github.com/project-planton/project-planton/pkg/fileutil"
	"github.com/project-planton/project-planton/pkg/iac/gitrepo"
)

func GetModulePath(moduleDir, kindName string) (string, error) {

	// If the module directory is provided, check if it is a valid terraform module directory
	if moduleDir != "" {
		// If the module directory is not provided, clone the repository and get the terraform module path
		isTerraformModuleDir, err := isTerraformModuleDirectory(moduleDir)
		if err != nil {
			return "", errors.Wrapf(err, "failed to check if %s is a valid terraform module directory", moduleDir)
		}

		// If the module directory is a valid terraform module directory, return the module directory
		if isTerraformModuleDir {
			return moduleDir, nil
		}
	}

	// If the module directory is not a valid terraform module directory,
	//clone the repository and get the terraform module path
	tofuModuleWorkspaceDir, err := getWorkspaceDir()
	if err != nil {
		return "", errors.Wrapf(err, "failed to get tofu module workspace directory")
	}

	// Clone the repository to the workspace directory
	gitRepoName, err := gitrepo.ExtractRepoName(gitrepo.CloneUrl)
	if err != nil {
		return "", errors.Wrapf(err, "failed to extract git repo name from %s", gitrepo.CloneUrl)
	}

	// Check if the cloned repository directory already exists
	terraformModuleRepoPath := filepath.Join(tofuModuleWorkspaceDir, gitRepoName)

	// If the cloned repository directory does not exist, clone the repository
	if _, statErr := os.Stat(terraformModuleRepoPath); os.IsNotExist(statErr) {
		gitCloneCommand := exec.Command("git", "clone", gitrepo.CloneUrl, terraformModuleRepoPath)
		gitCloneCommand.Stdout = os.Stdout
		gitCloneCommand.Stderr = os.Stderr
		if err := gitCloneCommand.Run(); err != nil {
			return "", errors.Wrapf(err, "failed to clone repository from %s to %s", gitrepo.CloneUrl, tofuModuleWorkspaceDir)
		}
	}

	//checkout the project-planton version tag if it is not the default version
	if version.Version != "" && version.Version != version.DefaultVersion {
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

// IsTerraformModuleDirectory checks if the given directory contains any files with .tf extension.
// It returns true if any .tf file exists, false otherwise. If an error occurs during the check, it returns an error.
func isTerraformModuleDirectory(moduleDir string) (bool, error) {
	entries, err := os.ReadDir(moduleDir)
	if err != nil {
		return false, errors.Wrapf(err, "failed to read directory %s", moduleDir)
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".tf") {
			return true, nil
		}
	}
	return false, nil
}

func getTerraformModulePath(moduleRepoDir, kindName string) (string, error) {
	kind := crkreflect.KindFromString(kindName)
	kindProvider := crkreflect.GetProvider(kind)
	if kindProvider == cloudresourcekind.CloudResourceProvider_cloud_resource_provider_unspecified {
		return "", errors.New("failed to get kind provider")
	}

	kindDirPath := filepath.Join(
		moduleRepoDir,
		"apis/org/project_planton/provider",
		strings.ReplaceAll(kindProvider.String(), "_", ""))

	terraformModulePath := filepath.Join(
		kindDirPath,
		strings.ToLower(kindName),
		"v1/iac/tf",
	)

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
	//base directory will always be ${HOME}/.planton-cloud/tofu
	tofuModuleWorkspaceDir := filepath.Join(cliWorkspaceDir, "tofu")
	if !fileutil.IsDirExists(tofuModuleWorkspaceDir) {
		if err := os.MkdirAll(tofuModuleWorkspaceDir, os.ModePerm); err != nil {
			return "", errors.Wrapf(err, "failed to ensure %s dir", tofuModuleWorkspaceDir)
		}
	}
	return tofuModuleWorkspaceDir, nil
}
