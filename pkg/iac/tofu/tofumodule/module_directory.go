package tofumodule

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/plantonhq/project-planton/apis/org/project_planton/shared/cloudresourcekind"
	"github.com/plantonhq/project-planton/internal/cli/cliprint"
	"github.com/plantonhq/project-planton/internal/cli/staging"
	"github.com/plantonhq/project-planton/internal/cli/version"
	"github.com/plantonhq/project-planton/internal/cli/workspace"
	"github.com/plantonhq/project-planton/pkg/crkreflect"
	"github.com/plantonhq/project-planton/pkg/fileutil"
)

// GetModulePathResult contains the module path and a cleanup function
type GetModulePathResult struct {
	ModulePath    string
	RepoPath      string
	CleanupFunc   func() error
	ShouldCleanup bool
}

// GetModulePath returns the path to the Terraform/OpenTofu module directory.
// If moduleDir is provided and is a valid Terraform module directory, it returns that.
// Otherwise, it ensures the staging area is set up and copies it to the tofu workspace.
// If moduleVersion is provided, it checks out that version (tag, branch, or commit SHA) in the workspace copy.
// The returned GetModulePathResult includes a cleanup function that should be called after execution
// unless noCleanup is true.
func GetModulePath(moduleDir, kindName, moduleVersion string, noCleanup bool) (*GetModulePathResult, error) {

	// If the module directory is provided, check if it is a valid terraform module directory
	if moduleDir != "" {
		isTerraformModuleDir, err := isTerraformModuleDirectory(moduleDir)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to check if %s is a valid terraform module directory", moduleDir)
		}

		// If the module directory is a valid terraform module directory, return the module directory
		if isTerraformModuleDir {
			return &GetModulePathResult{
				ModulePath:    moduleDir,
				RepoPath:      moduleDir,
				CleanupFunc:   func() error { return nil },
				ShouldCleanup: false,
			}, nil
		}
	}

	// Get the tofu workspace directory
	tofuModuleWorkspaceDir, err := getWorkspaceDir()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get tofu module workspace directory")
	}

	// Determine target version - use CLI version if set and not default
	targetVersion := ""
	if version.Version != "" && version.Version != version.DefaultVersion {
		targetVersion = version.Version
	}

	// Ensure staging is set up with the correct version
	cliprint.PrintStep("Ensuring staging area is ready...")
	if err := staging.EnsureStaging(targetVersion); err != nil {
		return nil, errors.Wrap(err, "failed to ensure staging area")
	}
	// Get and display current staging version
	stagingVersion, _ := staging.GetCurrentStagingVersion()
	if stagingVersion != "" {
		cliprint.PrintSuccess(fmt.Sprintf("Staging area ready (modules version: %s)", stagingVersion))
	} else {
		cliprint.PrintSuccess("Staging area ready")
	}

	// Copy from staging to tofu workspace
	cliprint.PrintStep("Copying modules to workspace...")
	terraformModuleRepoPath, err := staging.CopyToWorkspace(tofuModuleWorkspaceDir)
	if err != nil {
		return nil, errors.Wrap(err, "failed to copy staging to workspace")
	}
	cliprint.PrintSuccess("Modules copied to workspace")

	// If moduleVersion is specified, checkout that version in the workspace copy
	if moduleVersion != "" {
		if err := staging.CheckoutVersionInWorkspace(terraformModuleRepoPath, moduleVersion); err != nil {
			// Clean up on error
			_ = staging.CleanupWorkspaceCopy(terraformModuleRepoPath)
			return nil, errors.Wrapf(err, "failed to checkout module version %s", moduleVersion)
		}
	}

	terraformModulePath, err := getTerraformModulePath(terraformModuleRepoPath, kindName)
	if err != nil {
		// Clean up on error
		_ = staging.CleanupWorkspaceCopy(terraformModuleRepoPath)
		return nil, errors.Wrapf(err, "failed to get terraform module path for %s", kindName)
	}

	// Create cleanup function
	cleanupFunc := func() error {
		return staging.CleanupWorkspaceCopy(terraformModuleRepoPath)
	}

	return &GetModulePathResult{
		ModulePath:    terraformModulePath,
		RepoPath:      terraformModuleRepoPath,
		CleanupFunc:   cleanupFunc,
		ShouldCleanup: !noCleanup,
	}, nil
}

// GetModulePathLegacy is the legacy function signature for backward compatibility.
// It calls GetModulePath with noCleanup=false and no moduleVersion, returns just the module path.
// Note: This does not perform cleanup - callers should migrate to GetModulePath for proper cleanup handling.
func GetModulePathLegacy(moduleDir, kindName string) (string, error) {
	result, err := GetModulePath(moduleDir, kindName, "", false)
	if err != nil {
		return "", err
	}
	return result.ModulePath, nil
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
	//base directory will always be ${HOME}/.project-planton/tofu
	tofuModuleWorkspaceDir := filepath.Join(cliWorkspaceDir, "tofu")
	if !fileutil.IsDirExists(tofuModuleWorkspaceDir) {
		if err := os.MkdirAll(tofuModuleWorkspaceDir, os.ModePerm); err != nil {
			return "", errors.Wrapf(err, "failed to ensure %s dir", tofuModuleWorkspaceDir)
		}
	}
	return tofuModuleWorkspaceDir, nil
}
