package upgrade

import (
	"errors"
	"fmt"
	"os"
	"runtime"

	"github.com/fatih/color"
	"github.com/plantonhq/project-planton/internal/cli/cliprint"
	"github.com/plantonhq/project-planton/internal/cli/version"
)

// Run executes the upgrade command
func Run(checkOnly bool, force bool) {
	currentVersion := version.Version
	if currentVersion == "" {
		currentVersion = "dev"
	}

	// Step 1: Check for latest version
	cliprint.PrintStep("Checking for updates...")

	latestVersion, err := GetLatestVersion()
	if err != nil {
		cliprint.PrintError(fmt.Sprintf("Failed to check for updates: %v", err))
		fmt.Println()
		fmt.Println("You can manually download the latest version from:")
		fmt.Println("  https://github.com/plantonhq/project-planton/releases")
		os.Exit(1)
	}

	// Step 2: Compare versions
	needsUpgrade := CompareVersions(currentVersion, latestVersion)

	fmt.Println()
	if needsUpgrade {
		// Show versions with color distinction when update is available
		yellow := color.New(color.FgYellow).SprintFunc()
		green := color.New(color.FgGreen, color.Bold).SprintFunc()
		fmt.Printf("Current version: %s\n", yellow(currentVersion))
		fmt.Printf("Latest version:  %s\n", green(latestVersion))
	} else {
		// Show versions in matching green when up-to-date
		green := color.New(color.FgGreen, color.Bold).SprintFunc()
		fmt.Printf("Current version: %s\n", green(currentVersion))
		fmt.Printf("Latest version:  %s\n", green(latestVersion))
	}

	if !needsUpgrade && !force {
		fmt.Println()
		cliprint.PrintSuccess(fmt.Sprintf("project-planton is already up to date (%s)", currentVersion))
		return
	}

	if checkOnly {
		if needsUpgrade {
			fmt.Println()
			orange := color.New(color.FgYellow, color.Bold).SprintFunc()
			fmt.Printf("%s A new version is available!\n", orange("⚡"))
			fmt.Println()
			blue := color.New(color.FgCyan, color.Bold).SprintFunc()
			fmt.Printf("Run %s to update.\n", blue("project-planton upgrade"))
		}
		return
	}

	if !needsUpgrade && force {
		fmt.Println()
		cliprint.PrintStep("Forcing upgrade...")
	}

	// Step 3: Detect upgrade method
	method := DetectUpgradeMethod()

	fmt.Println()
	cliprint.PrintStep(fmt.Sprintf("Upgrade method: %s", method.String()))

	// Step 4: Perform upgrade
	var upgradeErr error
	switch method {
	case MethodHomebrew:
		upgradeErr = UpgradeViaHomebrew()
	case MethodDirectDownload:
		upgradeErr = UpgradeViaDirect(latestVersion)
	}

	if upgradeErr != nil {
		handleUpgradeError(upgradeErr, latestVersion)
		os.Exit(1)
	}

	// Step 5: Success message
	fmt.Println()
	cliprint.PrintSuccess(fmt.Sprintf("Successfully upgraded to %s", latestVersion))

	// Show platform-specific notes
	if method == MethodDirectDownload {
		fmt.Println()
		cliprint.PrintStep("Note: You may need to restart your terminal for changes to take effect.")
	}
}

// handleUpgradeError handles and displays upgrade errors with helpful suggestions
func handleUpgradeError(err error, latestVersion string) {
	fmt.Println()

	// Check for permission errors
	var permErr *PermissionError
	if errors.As(err, &permErr) {
		cliprint.PrintError(permErr.Error())
		fmt.Println()
		fmt.Println("Try running with sudo:")
		fmt.Println("  sudo project-planton upgrade")
		fmt.Println()
		fmt.Println("Or download manually to a user directory:")
		goos, goarch := GetPlatformInfo()
		downloadURL := BuildDownloadURL(latestVersion, goos, goarch)
		fmt.Printf("  curl -LO %s\n", downloadURL)
		if goos == "windows" {
			fmt.Println("  # Extract the zip file and move project-planton.exe to your PATH")
		} else {
			fmt.Println("  tar -xzf cli_*.tar.gz")
			fmt.Println("  chmod +x project-planton")
			fmt.Println("  mv project-planton ~/.local/bin/")
		}
		return
	}

	// Generic error
	cliprint.PrintError(fmt.Sprintf("Upgrade failed: %v", err))
	fmt.Println()
	fmt.Println("You can manually download the latest version from:")
	fmt.Println("  https://github.com/plantonhq/project-planton/releases")

	// Show platform-specific instructions
	goos, goarch := GetPlatformInfo()
	downloadURL := BuildDownloadURL(latestVersion, goos, goarch)

	fmt.Println()
	fmt.Println("Or download directly:")
	if runtime.GOOS == "windows" {
		fmt.Printf("  Invoke-WebRequest -Uri \"%s\" -OutFile \"cli.zip\"\n", downloadURL)
		fmt.Println("  Expand-Archive -Path \"cli.zip\" -DestinationPath \".\"")
		fmt.Println("  Move-Item -Path \"project-planton.exe\" -Destination \"C:\\Windows\\System32\\\"")
	} else {
		fmt.Printf("  curl -LO %s\n", downloadURL)
		fmt.Println("  tar -xzf cli_*.tar.gz")
		fmt.Println("  chmod +x project-planton")
		if runtime.GOOS == "darwin" {
			fmt.Println("  xattr -dr com.apple.quarantine project-planton  # Remove macOS quarantine")
		}
		fmt.Println("  sudo mv project-planton /usr/local/bin/")
	}
}
