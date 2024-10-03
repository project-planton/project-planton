package project_planton

import (
	"fmt"
	"github.com/plantoncloud/project-planton/cmd/project-planton/root"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "project-planton",
	Short: "Unified Interface for Multi-Cloud Infrastructure",
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.DisableSuggestions = true
	rootCmd.AddCommand(
		root.LoadManifest,
		root.Pulumi,
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
