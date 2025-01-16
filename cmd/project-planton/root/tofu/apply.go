package tofu

import (
	"github.com/project-planton/project-planton/apis/project/planton/shared/tofu"
	"github.com/project-planton/project-planton/internal/cli/flag"
	"github.com/project-planton/project-planton/internal/iac/stackinput/credentials"
	"github.com/project-planton/project-planton/internal/iac/tofu/tofumodule"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Apply = &cobra.Command{
	Use:   "apply",
	Short: "run tofu apply",
	Run:   applyHandler,
}

func init() {
	Apply.PersistentFlags().Bool(string(flag.AutoApprove), false, "Skip interactive approval of plan before applying")
}

func applyHandler(cmd *cobra.Command, args []string) {
	isAutoApprove, err := cmd.Flags().GetBool(string(flag.AutoApprove))
	flag.HandleFlagErr(err, flag.AutoApprove)

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

	err = tofumodule.RunCommand(moduleDir, targetManifestPath, tofu.TofuOperationType_apply,
		valueOverrides,
		isAutoApprove,
		credentialOptions...)
	if err != nil {
		log.Fatalf("failed to run tofu operation: %v", err)
	}
}
