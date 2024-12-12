package tofu

import (
	"github.com/project-planton/project-planton/internal/deploymentcomponent/manifest"
	"github.com/project-planton/project-planton/internal/iac/tofu/tfvars"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var LoadTfVars = &cobra.Command{
	Use:   "load-tfvars",
	Short: "load a project-planton manifest into tfvars format",
	Example: `
	project-planton tofu load-tfvars --manifest manifest.yaml
	`,
	Args: cobra.ExactArgs(1), //path of the manifest to load
	Run:  loadTfVarsHandler,
}

func loadTfVarsHandler(cmd *cobra.Command, args []string) {
	manifestPath := args[0]
	updatedManifest, err := manifest.LoadWithOverrides(manifestPath, map[string]string{})
	if err != nil {
		log.Fatal(err)
	}
	tfvarsString, err := tfvars.ProtoToTFVars(updatedManifest)
	if err != nil {
		log.Fatal(err)
	}
	println(tfvarsString)
}
