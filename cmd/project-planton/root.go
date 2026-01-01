package project_planton

import (
	"fmt"
	"os"

	"github.com/project-planton/project-planton/cmd/project-planton/root"
	"github.com/project-planton/project-planton/cmd/project-planton/root/webapp"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "project-planton",
	Short: "Unified Interface for Multi-Cloud Infrastructure",
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.DisableSuggestions = true
	rootCmd.AddCommand(
		root.Apply,
		root.Checkout,
		root.CloudResourceApplyCmd,
		root.CloudResourceCreateCmd,
		root.CloudResourceDeleteCmd,
		root.CloudResourceGetCmd,
		root.CloudResourceListCmd,
		root.CloudResourceUpdateCmd,
		root.ConfigCmd,
		root.CredentialCreateCmd,
		root.CredentialDeleteCmd,
		root.CredentialGetCmd,
		root.CredentialListCmd,
		root.CredentialUpdateCmd,
		root.Destroy,
		root.Init,
		root.LoadManifest,
		root.ModulesVersion,
		root.Plan,
		root.Pull,
		root.Pulumi,
		root.Refresh,
		root.StackUpdateStreamOutputCmd,
		root.Tofu,
		root.ValidateManifest,
		root.Version,
		webapp.WebAppCmd,
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
