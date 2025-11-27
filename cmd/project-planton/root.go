package project_planton

import (
	"fmt"
	"os"

	"github.com/project-planton/project-planton/cmd/project-planton/root"
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
		root.ConfigCmd,
		root.Destroy,
		root.Init,
		root.ListDeploymentComponent,
		root.LoadManifest,
		root.Plan,
		root.Pulumi,
		root.Refresh,
		root.Tofu,
		root.ValidateManifest,
		root.Version,
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
