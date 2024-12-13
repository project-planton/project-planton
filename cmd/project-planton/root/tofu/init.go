package tofu

import (
	"github.com/project-planton/project-planton/internal/cli/flag"
	"github.com/project-planton/project-planton/internal/iac/pulumi/stackinput/credentials"
	"github.com/project-planton/project-planton/internal/iac/tofu/tofumodule"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Init = &cobra.Command{
	Use:   "init",
	Short: "run tofu init",
	Run:   initHandler,
}

func init() {
	Init.PersistentFlags().StringArray(string(flag.BackendConfig), []string{}, "Configuration to be merged with what is in the\n                          configuration file's 'backend' block. This can be\n                          either a path to an HCL file with key/value\n                          assignments (same format as terraform.tfvars) or a\n                          'key=value' format, and can be specified multiple\n                          times. The backend type must be in the configuration\n                          itself.")
}

func initHandler(cmd *cobra.Command, args []string) {
	inputDir, err := cmd.Flags().GetString(string(flag.InputDir))
	flag.HandleFlagErr(err, flag.InputDir)

	moduleDir, err := cmd.Flags().GetString(string(flag.ModuleDir))
	flag.HandleFlagErrAndValue(err, flag.ModuleDir, moduleDir)

	valueOverrides, err := cmd.Flags().GetStringToString(string(flag.Set))
	flag.HandleFlagErr(err, flag.Set)

	backendConfigList, err := cmd.Flags().GetStringArray(string(flag.BackendConfig))
	flag.HandleFlagErr(err, flag.BackendConfig)

	credentialOptions := make([]credentials.StackInputCredentialOption, 0)
	targetManifestPath := inputDir + "/target.yaml"

	if inputDir == "" {
		targetManifestPath, err = cmd.Flags().GetString(string(flag.Manifest))
		flag.HandleFlagErrAndValue(err, flag.Manifest, targetManifestPath)

		credentialOptions, err = credentials.BuildWithFlags(cmd.Flags())
		if err != nil {
			log.Fatalf("failed to build credentiaal options: %v", err)
		}
	} else {
		credentialOptions, err = credentials.BuildWithFlags(cmd.Flags())
		if err != nil {
			log.Fatalf("failed to build credentiaal options: %v", err)
		}
	}

	err = tofumodule.TofuInit(moduleDir, targetManifestPath, valueOverrides, backendConfigList,
		credentialOptions...)
	if err != nil {
		log.Fatalf("failed to run pulumi: %v", err)
	}
}
