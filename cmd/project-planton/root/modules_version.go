package root

import (
	"fmt"
	"os"

	"github.com/plantonhq/project-planton/internal/cli/cliprint"
	"github.com/plantonhq/project-planton/internal/cli/staging"
	"github.com/spf13/cobra"
)

var ModulesVersion = &cobra.Command{
	Use:   "modules-version",
	Short: "Show the current version of IaC modules in the staging area",
	Long: `Display the currently checked out version of the ProjectPlanton IaC modules
in the local staging area.

The staging area (~/.project-planton/staging/project-planton) maintains a cached copy
of the ProjectPlanton repository containing all IaC modules (Pulumi and Terraform/OpenTofu).

This command reads the version from the staging area's .version file and displays it.
If the staging area doesn't exist, it will indicate that no modules are cached yet.

Use 'project-planton checkout <version>' to switch to a different version.
Use 'project-planton pull' to update to the latest version from upstream.`,
	Example: `  # Check current modules version
  project-planton modules-version

  # Typical workflow
  project-planton modules-version     # Check current version
  project-planton checkout v0.2.273   # Switch to specific version
  project-planton modules-version     # Verify the switch`,
	Run: modulesVersionHandler,
}

func modulesVersionHandler(cmd *cobra.Command, args []string) {
	exists, version, repoPath, err := staging.GetStagingInfo()
	if err != nil {
		cliprint.PrintError(fmt.Sprintf("Failed to get staging info: %v", err))
		os.Exit(1)
	}

	if !exists {
		fmt.Println("No IaC modules cached yet.")
		fmt.Println("")
		fmt.Println("Run 'project-planton pull' to clone the modules to the staging area,")
		fmt.Println("or run any apply/preview/destroy command to automatically set up staging.")
		return
	}

	fmt.Println("IaC Modules Staging Area")
	fmt.Println("========================")
	fmt.Printf("Location: %s\n", repoPath)
	if version != "" {
		fmt.Printf("Version:  %s\n", version)
	} else {
		fmt.Println("Version:  (unknown - .version file not found)")
	}
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("  project-planton pull                  Update to latest from upstream")
	fmt.Println("  project-planton checkout <version>    Switch to a specific version")
}
