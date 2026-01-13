package upgrade

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/plantonhq/project-planton/internal/cli/cliprint"
)

// UpgradeViaHomebrew upgrades the CLI using Homebrew cask
func UpgradeViaHomebrew() error {
	// Step 1: Update Homebrew
	fmt.Println()
	cliprint.PrintStep("Updating Homebrew...")

	updateCmd := exec.Command("brew", "update")
	updateCmd.Stdout = os.Stdout
	updateCmd.Stderr = os.Stderr
	if err := updateCmd.Run(); err != nil {
		return fmt.Errorf("failed to update Homebrew: %w", err)
	}

	// Step 2: Upgrade project-planton cask
	fmt.Println()
	cliprint.PrintStep("Upgrading project-planton...")

	upgradeCmd := exec.Command("brew", "upgrade", "--cask", "project-planton")
	upgradeCmd.Stdout = os.Stdout
	upgradeCmd.Stderr = os.Stderr
	if err := upgradeCmd.Run(); err != nil {
		// brew upgrade returns non-zero if already up to date, which is fine
		// Check if project-planton is actually installed and at latest
		return fmt.Errorf("Homebrew upgrade failed: %w", err)
	}

	return nil
}
