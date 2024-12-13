package tofu

import (
	"github.com/project-planton/project-planton/apis/project/planton/shared/tofu"
	"github.com/project-planton/project-planton/internal/cli/flag"
	"github.com/project-planton/project-planton/internal/iac/pulumi/stackinput/credentials"
	"github.com/project-planton/project-planton/internal/iac/tofu/tofumodule"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Refresh = &cobra.Command{
	Use:   "refresh",
	Short: "run tofu refresh",
	Run:   refreshHandler,
}

func refreshHandler(cmd *cobra.Command, args []string) {
	inputDir, err := cmd.Flags().GetString(string(flag.InputDir))
	flag.HandleFlagErr(err, flag.InputDir)

	moduleDir, err := cmd.Flags().GetString(string(flag.ModuleDir))
	flag.HandleFlagErrAndValue(err, flag.ModuleDir, moduleDir)

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

	err = tofumodule.RunOperation(moduleDir, targetManifestPath, tofu.TofuOperationType_refresh, valueOverrides,
		credentialOptions...)
	if err != nil {
		log.Fatalf("failed to run pulumi: %v", err)
	}
}
