package pulumibinary

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/plantonhq/project-planton/internal/cli/cliprint"
	"github.com/plantonhq/project-planton/internal/cli/version"
	"github.com/plantonhq/project-planton/internal/cli/workspace"
	"github.com/plantonhq/project-planton/pkg/fileutil"
)

const (
	// PulumiDirName is the base directory name for all Pulumi-related files
	// All Pulumi files are stored under ~/.project-planton/pulumi/
	PulumiDirName = "pulumi"

	// BinariesSubDir is the subdirectory for cached binaries
	// Full path: ~/.project-planton/pulumi/binaries/{version}/
	BinariesSubDir = "binaries"

	// WorkspacesSubDir is the subdirectory for Pulumi workspaces
	// Full path: ~/.project-planton/pulumi/workspaces/{stack-fqdn}/
	WorkspacesSubDir = "workspaces"

	// GitHubReleaseBaseURL is the base URL for GitHub releases
	GitHubReleaseBaseURL = "https://github.com/plantonhq/project-planton/releases/download"

	// BinaryPrefix is the prefix for Pulumi component binaries
	BinaryPrefix = "pulumi-"
)

// GetPulumiBaseDir returns the base directory for all Pulumi-related files
// (~/.project-planton/pulumi/)
func GetPulumiBaseDir() (string, error) {
	workspaceDir, err := workspace.GetWorkspaceDir()
	if err != nil {
		return "", errors.Wrap(err, "failed to get workspace directory")
	}
	return filepath.Join(workspaceDir, PulumiDirName), nil
}

// GetBinaryCacheDir returns the path to the binary cache directory
// (~/.project-planton/pulumi/binaries/{version}/)
func GetBinaryCacheDir(releaseVersion string) (string, error) {
	pulumiBaseDir, err := GetPulumiBaseDir()
	if err != nil {
		return "", err
	}

	// Normalize version for directory name
	versionDir := releaseVersion
	if versionDir == "" || versionDir == version.DefaultVersion {
		versionDir = "dev"
	}

	return filepath.Join(pulumiBaseDir, BinariesSubDir, versionDir), nil
}

// GetBinaryPath returns the expected path for a cached binary
func GetBinaryPath(componentName, releaseVersion string) (string, error) {
	cacheDir, err := GetBinaryCacheDir(releaseVersion)
	if err != nil {
		return "", err
	}

	binaryName := BuildBinaryName(componentName)
	return filepath.Join(cacheDir, binaryName), nil
}

// BuildBinaryName constructs the binary filename for a component
// e.g., "AwsEcsService" -> "pulumi-awsecsservice"
func BuildBinaryName(componentName string) string {
	return BinaryPrefix + strings.ToLower(componentName)
}

// BuildDownloadURL constructs the download URL for a component binary.
// The release version can be:
// - A semantic version like "v0.3.2" (downloads from main project-planton release)
// - An auto-release version like "v0.3.1-pulumi-awsecsservice-20260107.01" (downloads from component-specific release)
//
// Examples:
//   - BuildDownloadURL("AwsEcsService", "v0.3.2")
//     -> https://github.com/plantonhq/project-planton/releases/download/v0.3.2/pulumi-awsecsservice.gz
//   - BuildDownloadURL("AwsEcsService", "v0.3.1-pulumi-awsecsservice-20260107.01")
//     -> https://github.com/plantonhq/project-planton/releases/download/v0.3.1-pulumi-awsecsservice-20260107.01/pulumi-awsecsservice.gz
func BuildDownloadURL(componentName, releaseVersion string) string {
	binaryName := BuildBinaryName(componentName)
	artifactName := binaryName + ".gz"
	return fmt.Sprintf("%s/%s/%s", GitHubReleaseBaseURL, releaseVersion, artifactName)
}

// IsBinaryCached checks if a binary is already cached
func IsBinaryCached(componentName, releaseVersion string) (bool, error) {
	binaryPath, err := GetBinaryPath(componentName, releaseVersion)
	if err != nil {
		return false, err
	}

	exists, err := fileutil.IsExists(binaryPath)
	if err != nil {
		return false, errors.Wrapf(err, "failed to check if binary exists at %s", binaryPath)
	}

	if !exists {
		return false, nil
	}

	// Verify it's executable
	info, err := os.Stat(binaryPath)
	if err != nil {
		return false, errors.Wrapf(err, "failed to stat binary at %s", binaryPath)
	}

	// Check if file has execute permission
	return info.Mode()&0111 != 0, nil
}

// EnsureBinary ensures the binary for a component is downloaded and cached.
// The releaseVersion can be:
// - CLI version like "v0.3.2" (uses main project-planton release)
// - Module version like "v0.3.1-pulumi-awsecsservice-20260107.01" (uses component-specific release)
// Returns the path to the binary.
func EnsureBinary(componentName, releaseVersion string) (string, error) {
	// Check if already cached
	cached, err := IsBinaryCached(componentName, releaseVersion)
	if err != nil {
		return "", errors.Wrap(err, "failed to check binary cache")
	}

	binaryPath, err := GetBinaryPath(componentName, releaseVersion)
	if err != nil {
		return "", err
	}

	if cached {
		cliprint.PrintSuccess(fmt.Sprintf("Using cached binary: %s", filepath.Base(binaryPath)))
		return binaryPath, nil
	}

	// Download the binary
	cliprint.PrintStep(fmt.Sprintf("Downloading Pulumi binary for %s...", componentName))

	if err := DownloadBinary(componentName, releaseVersion); err != nil {
		return "", errors.Wrapf(err, "failed to download binary for %s", componentName)
	}

	cliprint.PrintSuccess(fmt.Sprintf("Binary downloaded: %s", filepath.Base(binaryPath)))
	return binaryPath, nil
}

// DownloadBinary downloads and extracts a component binary from GitHub releases.
// The releaseVersion determines which GitHub release to download from:
// - "v0.3.2" -> downloads from https://github.com/plantonhq/project-planton/releases/download/v0.3.2/pulumi-{component}.gz
// - "v0.3.1-pulumi-awsecsservice-20260107.01" -> downloads from that specific auto-release
func DownloadBinary(componentName, releaseVersion string) error {
	// Ensure cache directory exists
	cacheDir, err := GetBinaryCacheDir(releaseVersion)
	if err != nil {
		return err
	}

	if !fileutil.IsDirExists(cacheDir) {
		if err := os.MkdirAll(cacheDir, 0755); err != nil {
			return errors.Wrapf(err, "failed to create cache directory %s", cacheDir)
		}
	}

	// Build download URL - the release version IS the tag
	downloadURL := BuildDownloadURL(componentName, releaseVersion)

	cliprint.PrintInfo(fmt.Sprintf("Downloading from: %s", downloadURL))

	// Download the gzipped binary
	resp, err := http.Get(downloadURL)
	if err != nil {
		return errors.Wrapf(err, "failed to download from %s", downloadURL)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("download failed with status %d: %s", resp.StatusCode, resp.Status)
	}

	// Create gzip reader
	gzReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to create gzip reader")
	}
	defer gzReader.Close()

	// Write to destination
	binaryPath, err := GetBinaryPath(componentName, releaseVersion)
	if err != nil {
		return err
	}

	outFile, err := os.OpenFile(binaryPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return errors.Wrapf(err, "failed to create binary file at %s", binaryPath)
	}
	defer outFile.Close()

	written, err := io.Copy(outFile, gzReader)
	if err != nil {
		// Clean up partial file
		os.Remove(binaryPath)
		return errors.Wrap(err, "failed to extract binary")
	}

	cliprint.PrintInfo(fmt.Sprintf("Extracted %d bytes", written))

	return nil
}

// GetCurrentCLIVersion returns the current CLI version, falling back to "dev" if not set
func GetCurrentCLIVersion() string {
	if version.Version == "" || version.Version == version.DefaultVersion {
		return "dev"
	}
	return version.Version
}

// IsDevVersion checks if the current CLI is a development version
func IsDevVersion() bool {
	return version.Version == "" || version.Version == version.DefaultVersion
}
