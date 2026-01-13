package root

import (
	"github.com/plantonhq/project-planton/internal/cli/upgrade"
	"github.com/spf13/cobra"
)

var Upgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "upgrade the project-planton CLI to the latest version",
	Long: `Upgrade the project-planton CLI to the latest available version.

On macOS, if project-planton was installed via Homebrew, this command uses 'brew upgrade --cask'.
On all other platforms (or if Homebrew is not available), it downloads the latest
binary directly from GitHub releases.

Examples:
  # Upgrade to the latest version
  project-planton upgrade

  # Check for updates without installing
  project-planton upgrade --check

  # Force upgrade even if already on latest version
  project-planton upgrade --force`,
	Run: upgradeHandler,
}

func init() {
	Upgrade.Flags().BoolP("check", "c", false, "check for updates without installing")
	Upgrade.Flags().BoolP("force", "f", false, "force upgrade even if already on latest version")
}

func upgradeHandler(cmd *cobra.Command, args []string) {
	checkOnly, _ := cmd.Flags().GetBool("check")
	force, _ := cmd.Flags().GetBool("force")

	upgrade.Run(checkOnly, force)
}
