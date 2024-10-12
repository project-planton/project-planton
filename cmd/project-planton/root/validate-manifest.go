package root

import (
	"fmt"
	"github.com/project-planton/project-planton/internal/manifest"
	"github.com/spf13/cobra"
)

var ValidateManifest = &cobra.Command{
	Use:   "validate-manifest",
	Short: "validate a project-planton manifest",
	Aliases: []string{
		"validate",
	},
	Example: `
	project-planton validate manifest.yaml
	`,
	Args: cobra.ExactArgs(1), //path or url of the manifest to validate
	Run:  validateHandler,
}

func validateHandler(cmd *cobra.Command, args []string) {
	err := manifest.Validate(args[0])
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	println("manifest is valid")
}
