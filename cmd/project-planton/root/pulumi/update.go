package pulumi

import (
	"github.com/project-planton/project-planton/apis/project/planton/shared/iac/pulumi"
	"github.com/project-planton/project-planton/internal/cli/flag"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumistack"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputcredentials"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Update = &cobra.Command{
	Use:     "update",
	Short:   "run pulumi update",
	Aliases: []string{"up"},
	Run:     updateHandler,
}

func updateHandler(cmd *cobra.Command, args []string) {
	inputDir, err := cmd.Flags().GetString(string(flag.InputDir))
	flag.HandleFlagErr(err, flag.InputDir)

	moduleDir, err := cmd.Flags().GetString(string(flag.ModuleDir))
	flag.HandleFlagErrAndValue(err, flag.ModuleDir, moduleDir)

	stackFqdn, err := cmd.Flags().GetString(string(flag.Stack))
	flag.HandleFlagErrAndValue(err, flag.Stack, stackFqdn)

	valueOverrides, err := cmd.Flags().GetStringToString(string(flag.Set))
	flag.HandleFlagErr(err, flag.Set)

	credentialOptions := make([]stackinputcredentials.StackInputCredentialOption, 0)
	targetManifestPath := inputDir + "/target.yaml"

	if inputDir == "" {
		targetManifestPath, err = cmd.Flags().GetString(string(flag.Manifest))
		flag.HandleFlagErrAndValue(err, flag.Manifest, targetManifestPath)
	}

	credentialOptions, err = stackinputcredentials.BuildWithFlags(cmd.Flags())
	if err != nil {
		log.Fatalf("failed to build credentiaal options: %v", err)
	}

	err = pulumistack.Run(moduleDir, stackFqdn, targetManifestPath,
		pulumi.PulumiOperationType_update, false, true, valueOverrides, credentialOptions...)
	if err != nil {
		log.Fatalf("failed to run pulumi: %v", err)
	}
}
