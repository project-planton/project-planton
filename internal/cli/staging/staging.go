package staging

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/internal/cli/workspace"
	"github.com/project-planton/project-planton/pkg/fileutil"
	"github.com/project-planton/project-planton/pkg/iac/gitrepo"
)

const (
	stagingDirName  = "staging"
	versionFileName = ".version"
)

// GetStagingDir returns the path to the staging directory (~/.project-planton/staging)
func GetStagingDir() (string, error) {
	workspaceDir, err := workspace.GetWorkspaceDir()
	if err != nil {
		return "", errors.Wrap(err, "failed to get workspace directory")
	}
	return filepath.Join(workspaceDir, stagingDirName), nil
}

// GetStagingRepoPath returns the path to the cloned repository in staging
// (~/.project-planton/staging/project-planton)
func GetStagingRepoPath() (string, error) {
	stagingDir, err := GetStagingDir()
	if err != nil {
		return "", err
	}

	repoName, err := gitrepo.ExtractRepoName(gitrepo.CloneUrl)
	if err != nil {
		return "", errors.Wrapf(err, "failed to extract repo name from %s", gitrepo.CloneUrl)
	}

	return filepath.Join(stagingDir, repoName), nil
}

// GetVersionFilePath returns the path to the version tracking file
func GetVersionFilePath() (string, error) {
	stagingDir, err := GetStagingDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(stagingDir, versionFileName), nil
}

// GetCurrentStagingVersion reads the currently checked out version from the version file
func GetCurrentStagingVersion() (string, error) {
	versionFile, err := GetVersionFilePath()
	if err != nil {
		return "", err
	}

	if !fileutil.IsDirExists(filepath.Dir(versionFile)) {
		return "", nil // Staging doesn't exist yet
	}

	content, err := os.ReadFile(versionFile)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", errors.Wrap(err, "failed to read version file")
	}

	return strings.TrimSpace(string(content)), nil
}

// WriteCurrentStagingVersion writes the version to the version file
func WriteCurrentStagingVersion(version string) error {
	versionFile, err := GetVersionFilePath()
	if err != nil {
		return err
	}

	// Ensure staging directory exists
	stagingDir := filepath.Dir(versionFile)
	if !fileutil.IsDirExists(stagingDir) {
		if err := os.MkdirAll(stagingDir, os.ModePerm); err != nil {
			return errors.Wrapf(err, "failed to create staging directory %s", stagingDir)
		}
	}

	if err := os.WriteFile(versionFile, []byte(version+"\n"), 0644); err != nil {
		return errors.Wrap(err, "failed to write version file")
	}

	return nil
}

// StagingExists checks if the staging repository already exists
func StagingExists() (bool, error) {
	repoPath, err := GetStagingRepoPath()
	if err != nil {
		return false, err
	}

	// Check if it's a valid git repository
	gitDir := filepath.Join(repoPath, ".git")
	return fileutil.IsDirExists(gitDir), nil
}

// EnsureStaging ensures the staging repository exists and is at the correct version.
// If staging doesn't exist, it clones the repository.
// If the version doesn't match, it checks out the correct version.
func EnsureStaging(targetVersion string) error {
	exists, err := StagingExists()
	if err != nil {
		return errors.Wrap(err, "failed to check if staging exists")
	}

	repoPath, err := GetStagingRepoPath()
	if err != nil {
		return err
	}

	if !exists {
		// Clone the repository
		if err := cloneToStaging(); err != nil {
			return errors.Wrap(err, "failed to clone repository to staging")
		}
	}

	// Check current version
	currentVersion, err := GetCurrentStagingVersion()
	if err != nil {
		return errors.Wrap(err, "failed to get current staging version")
	}

	// If version matches, we're done
	if currentVersion == targetVersion && targetVersion != "" {
		return nil
	}

	// Checkout the target version if specified
	if targetVersion != "" {
		if err := checkoutVersion(repoPath, targetVersion); err != nil {
			return errors.Wrapf(err, "failed to checkout version %s", targetVersion)
		}

		if err := WriteCurrentStagingVersion(targetVersion); err != nil {
			return errors.Wrap(err, "failed to write version file")
		}
	}

	return nil
}

// cloneToStaging clones the repository to the staging directory
func cloneToStaging() error {
	stagingDir, err := GetStagingDir()
	if err != nil {
		return err
	}

	// Ensure staging directory exists
	if !fileutil.IsDirExists(stagingDir) {
		if err := os.MkdirAll(stagingDir, os.ModePerm); err != nil {
			return errors.Wrapf(err, "failed to create staging directory %s", stagingDir)
		}
	}

	repoPath, err := GetStagingRepoPath()
	if err != nil {
		return err
	}

	fmt.Printf("Cloning ProjectPlanton repository to staging area...\n")
	fmt.Printf("This is a one-time operation. Future executions will use local copy.\n\n")

	cmd := exec.Command("git", "clone", gitrepo.CloneUrl, repoPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "failed to clone repository from %s", gitrepo.CloneUrl)
	}

	fmt.Printf("\nRepository cloned to: %s\n", repoPath)
	return nil
}

// checkoutVersion checks out a specific version/tag/branch in the repository
func checkoutVersion(repoPath, version string) error {
	// First fetch to ensure we have all tags and branches
	fetchCmd := exec.Command("git", "-C", repoPath, "fetch", "--all", "--tags")
	fetchCmd.Stdout = os.Stdout
	fetchCmd.Stderr = os.Stderr
	if err := fetchCmd.Run(); err != nil {
		return errors.Wrap(err, "failed to fetch tags")
	}

	// Checkout the version
	checkoutCmd := exec.Command("git", "-C", repoPath, "checkout", version)
	checkoutCmd.Stdout = os.Stdout
	checkoutCmd.Stderr = os.Stderr

	if err := checkoutCmd.Run(); err != nil {
		return errors.Wrapf(err, "failed to checkout %s", version)
	}

	return nil
}

// Pull fetches and pulls the latest changes from upstream
func Pull() error {
	exists, err := StagingExists()
	if err != nil {
		return errors.Wrap(err, "failed to check if staging exists")
	}

	if !exists {
		// Clone first if staging doesn't exist
		if err := cloneToStaging(); err != nil {
			return err
		}
		return nil
	}

	repoPath, err := GetStagingRepoPath()
	if err != nil {
		return err
	}

	fmt.Printf("Fetching latest changes from upstream...\n\n")

	// Fetch all remotes
	fetchCmd := exec.Command("git", "-C", repoPath, "fetch", "--all", "--tags")
	fetchCmd.Stdout = os.Stdout
	fetchCmd.Stderr = os.Stderr
	if err := fetchCmd.Run(); err != nil {
		return errors.Wrap(err, "failed to fetch from upstream")
	}

	// Pull changes
	pullCmd := exec.Command("git", "-C", repoPath, "pull")
	pullCmd.Stdout = os.Stdout
	pullCmd.Stderr = os.Stderr
	if err := pullCmd.Run(); err != nil {
		return errors.Wrap(err, "failed to pull from upstream")
	}

	// Update version file with current HEAD
	headVersion, err := getCurrentHead(repoPath)
	if err != nil {
		return errors.Wrap(err, "failed to get current HEAD")
	}

	if err := WriteCurrentStagingVersion(headVersion); err != nil {
		return errors.Wrap(err, "failed to update version file")
	}

	return nil
}

// Checkout checks out a specific version in the staging area
func Checkout(version string) error {
	exists, err := StagingExists()
	if err != nil {
		return errors.Wrap(err, "failed to check if staging exists")
	}

	if !exists {
		// Clone first if staging doesn't exist
		if err := cloneToStaging(); err != nil {
			return err
		}
	}

	repoPath, err := GetStagingRepoPath()
	if err != nil {
		return err
	}

	fmt.Printf("Checking out version: %s\n\n", version)

	if err := checkoutVersion(repoPath, version); err != nil {
		return err
	}

	if err := WriteCurrentStagingVersion(version); err != nil {
		return errors.Wrap(err, "failed to update version file")
	}

	fmt.Printf("\nSuccessfully checked out: %s\n", version)
	return nil
}

// getCurrentHead returns the current HEAD reference (tag if on tag, otherwise commit SHA)
func getCurrentHead(repoPath string) (string, error) {
	// Try to get current tag
	tagCmd := exec.Command("git", "-C", repoPath, "describe", "--tags", "--exact-match")
	tagOutput, err := tagCmd.Output()
	if err == nil {
		return strings.TrimSpace(string(tagOutput)), nil
	}

	// Fall back to commit SHA
	shaCmd := exec.Command("git", "-C", repoPath, "rev-parse", "--short", "HEAD")
	shaOutput, err := shaCmd.Output()
	if err != nil {
		return "", errors.Wrap(err, "failed to get HEAD SHA")
	}

	return strings.TrimSpace(string(shaOutput)), nil
}

// CopyToWorkspace copies the staging repository to the specified destination directory.
// Returns the path to the copied repository.
func CopyToWorkspace(destDir string) (string, error) {
	repoPath, err := GetStagingRepoPath()
	if err != nil {
		return "", err
	}

	repoName, err := gitrepo.ExtractRepoName(gitrepo.CloneUrl)
	if err != nil {
		return "", errors.Wrapf(err, "failed to extract repo name")
	}

	destPath := filepath.Join(destDir, repoName)

	// Remove destination if it exists
	if fileutil.IsDirExists(destPath) {
		if err := os.RemoveAll(destPath); err != nil {
			return "", errors.Wrapf(err, "failed to remove existing directory %s", destPath)
		}
	}

	// Ensure destination parent directory exists
	if !fileutil.IsDirExists(destDir) {
		if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
			return "", errors.Wrapf(err, "failed to create destination directory %s", destDir)
		}
	}

	// Copy using cp -r for efficiency
	// Use cp -a to preserve permissions and symlinks
	cmd := exec.Command("cp", "-a", repoPath, destPath)
	if err := cmd.Run(); err != nil {
		return "", errors.Wrapf(err, "failed to copy staging to %s", destPath)
	}

	return destPath, nil
}

// CleanupWorkspaceCopy removes the copied repository from the workspace
func CleanupWorkspaceCopy(repoPath string) error {
	if repoPath == "" {
		return nil
	}

	// Safety check: don't delete the staging directory
	stagingPath, err := GetStagingRepoPath()
	if err != nil {
		return err
	}

	if repoPath == stagingPath {
		return errors.New("refusing to delete staging directory")
	}

	if !fileutil.IsDirExists(repoPath) {
		return nil // Already cleaned up
	}

	if err := os.RemoveAll(repoPath); err != nil {
		return errors.Wrapf(err, "failed to cleanup workspace copy at %s", repoPath)
	}

	return nil
}

// GetStagingInfo returns information about the current staging state
func GetStagingInfo() (exists bool, version string, repoPath string, err error) {
	exists, err = StagingExists()
	if err != nil {
		return false, "", "", err
	}

	if !exists {
		return false, "", "", nil
	}

	version, err = GetCurrentStagingVersion()
	if err != nil {
		return exists, "", "", err
	}

	repoPath, err = GetStagingRepoPath()
	if err != nil {
		return exists, version, "", err
	}

	return exists, version, repoPath, nil
}

// copyDir recursively copies a directory tree
// Note: This is kept as a fallback but we prefer cp -a for efficiency
func copyDir(src string, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Calculate the destination path
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		// Copy file
		return copyFile(path, dstPath)
	})
}

// copyFile copies a single file
func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	srcInfo, err := srcFile.Stat()
	if err != nil {
		return err
	}

	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, srcInfo.Mode())
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

