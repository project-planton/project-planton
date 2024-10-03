package root

import (
	"github.com/plantoncloud/project-planton/internal/manifestvalidator"
	log "github.com/sirupsen/logrus"
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
	Args: cobra.ExactArgs(1), //path of the manifest to validate
	Run:  validateHandler,
}

func validateHandler(cmd *cobra.Command, args []string) {
	if err := manifestvalidator.Validate(args[0]); err != nil {
		log.Fatal(err)
	}
}
