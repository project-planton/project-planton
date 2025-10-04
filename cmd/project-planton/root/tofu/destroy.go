package tofu

import (
	"os"

	"github.com/project-planton/project-planton/apis/project/planton/shared/iac/terraform"
	"github.com/project-planton/project-planton/internal/cli/flag"
	"github.com/project-planton/project-planton/internal/cli/manifest"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputcredentials"
	"github.com/project-planton/project-planton/pkg/iac/tofu/tofumodule"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Destroy = &cobra.Command{
	Use:   "destroy",
	Short: "run tofu destroy",
	Run:   destroyHandler,
}

func init() {
	Destroy.PersistentFlags().Bool(string(flag.AutoApprove), false, "Skip interactive approval of plan before applying")
}

func destroyHandler(cmd *cobra.Command, args []string) {
	isAutoApprove, err := cmd.Flags().GetBool(string(flag.AutoApprove))
	flag.HandleFlagErr(err, flag.AutoApprove)

	moduleDir, err := cmd.Flags().GetString(string(flag.ModuleDir))
	flag.HandleFlagErrAndValue(err, flag.ModuleDir, moduleDir)

	valueOverrides, err := cmd.Flags().GetStringToString(string(flag.Set))
	flag.HandleFlagErr(err, flag.Set)

	// Resolve manifest path with priority: --manifest > --input-dir > --kustomize-dir + --overlay
	targetManifestPath, isTemp, err := manifest.ResolveManifestPath(cmd)
	if err != nil {
		log.Fatalf("failed to resolve manifest: %v", err)
	}
	if isTemp {
		defer os.Remove(targetManifestPath)
	}

	credentialOptions, err := stackinputcredentials.BuildWithFlags(cmd.Flags())
	if err != nil {
		log.Fatalf("failed to build credentiaal options: %v", err)
	}

	err = tofumodule.RunCommand(moduleDir, targetManifestPath, terraform.TerraformOperationType_destroy, valueOverrides,
		isAutoApprove,
		false,
		credentialOptions...)
	if err != nil {
		log.Fatalf("failed to run tofu operation: %v", err)
	}
}
