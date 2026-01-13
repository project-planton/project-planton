package upgrade

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	// githubReleasesURL is the GitHub API endpoint for releases
	githubReleasesURL = "https://api.github.com/repos/plantonhq/project-planton/releases"
	// httpTimeout is the timeout for HTTP requests
	httpTimeout = 10 * time.Second
)

// GitHubRelease represents a GitHub release from the API
type GitHubRelease struct {
	TagName    string `json:"tag_name"`
	Draft      bool   `json:"draft"`
	Prerelease bool   `json:"prerelease"`
}

// cliVersionInfo holds parsed CLI version information for comparison
type cliVersionInfo struct {
	Tag      string
	Major    int
	Minor    int
	Patch    int
	Date     int // YYYYMMDD as integer for comparison
	Revision int // The .N suffix
}

// cliVersionRegex matches CLI release tags like v0.3.15-cli.20260113.0
var cliVersionRegex = regexp.MustCompile(`^v?(\d+)\.(\d+)\.(\d+)-cli\.(\d{8})\.(\d+)$`)

// parseCliVersion parses a CLI version tag into its components
func parseCliVersion(tag string) (*cliVersionInfo, error) {
	matches := cliVersionRegex.FindStringSubmatch(tag)
	if matches == nil {
		return nil, fmt.Errorf("not a valid CLI version tag")
	}

	major, _ := strconv.Atoi(matches[1])
	minor, _ := strconv.Atoi(matches[2])
	patch, _ := strconv.Atoi(matches[3])
	date, _ := strconv.Atoi(matches[4])
	revision, _ := strconv.Atoi(matches[5])

	return &cliVersionInfo{
		Tag:      tag,
		Major:    major,
		Minor:    minor,
		Patch:    patch,
		Date:     date,
		Revision: revision,
	}, nil
}

// isGreaterThan returns true if v is greater than other
func (v *cliVersionInfo) isGreaterThan(other *cliVersionInfo) bool {
	if v.Major != other.Major {
		return v.Major > other.Major
	}
	if v.Minor != other.Minor {
		return v.Minor > other.Minor
	}
	if v.Patch != other.Patch {
		return v.Patch > other.Patch
	}
	if v.Date != other.Date {
		return v.Date > other.Date
	}
	return v.Revision > other.Revision
}

// GetLatestVersion fetches the latest CLI version from GitHub releases
// It filters for CLI-specific releases (tagged with -cli.) and returns
// the highest version based on semver + date + revision
func GetLatestVersion() (string, error) {
	client := &http.Client{
		Timeout: httpTimeout,
	}

	req, err := http.NewRequest("GET", githubReleasesURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Set Accept header for GitHub API
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch releases: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch releases: HTTP %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	var releases []GitHubRelease
	if err := json.Unmarshal(body, &releases); err != nil {
		return "", fmt.Errorf("failed to parse releases: %w", err)
	}

	if len(releases) == 0 {
		return "", fmt.Errorf("no releases found")
	}

	// Find the highest CLI version from published releases
	var highest *cliVersionInfo

	for _, release := range releases {
		// Skip drafts and pre-releases
		if release.Draft || release.Prerelease {
			continue
		}

		// Only consider CLI releases (tags containing "-cli.")
		if !strings.Contains(release.TagName, "-cli.") {
			continue
		}

		// Parse the CLI version tag
		version, err := parseCliVersion(release.TagName)
		if err != nil {
			continue
		}

		// Check if this is the highest version so far
		if highest == nil || version.isGreaterThan(highest) {
			highest = version
		}
	}

	if highest == nil {
		return "", fmt.Errorf("no valid CLI releases found")
	}

	return highest.Tag, nil
}

// CompareVersions returns true if latestVersion is newer than currentVersion
func CompareVersions(currentVersion, latestVersion string) bool {
	// If versions are the same, no upgrade needed
	if currentVersion == latestVersion {
		return false
	}

	// If current is "dev" or empty, always consider latest as newer
	if currentVersion == "dev" || currentVersion == "" {
		return true
	}

	// Try to parse both as CLI versions
	currentCli, errCurrent := parseCliVersion(currentVersion)
	latestCli, errLatest := parseCliVersion(latestVersion)

	// If both parse as CLI versions, compare them properly
	if errCurrent == nil && errLatest == nil {
		return latestCli.isGreaterThan(currentCli)
	}

	// Fall back to string comparison if parsing fails
	return currentVersion != latestVersion
}
