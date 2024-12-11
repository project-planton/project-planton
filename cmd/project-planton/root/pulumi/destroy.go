package pulumi

import (
	"github.com/project-planton/project-planton/apis/project/planton/shared/pulumi"
	"github.com/project-planton/project-planton/internal/cli/flag"
	"github.com/project-planton/project-planton/internal/iac/pulumi/pulumistack"
	"github.com/project-planton/project-planton/internal/iac/pulumi/stackinput/credentials"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Destroy = &cobra.Command{
	Use:   "destroy",
	Short: "run pulumi destroy",
	Run:   destroyHandler,
}

func destroyHandler(cmd *cobra.Command, args []string) {
	inputDir, err := cmd.Flags().GetString(string(flag.InputDir))
	flag.HandleFlagErr(err, flag.InputDir)

	moduleDir, err := cmd.Flags().GetString(string(flag.ModuleDir))
	flag.HandleFlagErrAndValue(err, flag.ModuleDir, moduleDir)

	stackFqdn, err := cmd.Flags().GetString(string(flag.Stack))
	flag.HandleFlagErrAndValue(err, flag.Stack, stackFqdn)

	valueOverrides, err := cmd.Flags().GetStringToString(string(flag.Set))
	flag.HandleFlagErr(err, flag.Set)

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

	err = pulumistack.Run(moduleDir, stackFqdn, targetManifestPath,
		pulumi.PulumiOperationType_destroy, false, valueOverrides, credentialOptions...)
	if err != nil {
		log.Fatalf("failed to run pulumi: %v", err)
	}
}
