package tofu

import (
	"fmt"
	"os"

	"github.com/project-planton/project-planton/apis/project/planton/shared/iac/terraform"
	"github.com/project-planton/project-planton/internal/cli/flag"
	climanifest "github.com/project-planton/project-planton/internal/cli/manifest"
	"github.com/project-planton/project-planton/internal/manifest"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputcredentials"
	"github.com/project-planton/project-planton/pkg/iac/tofu/tofumodule"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Plan = &cobra.Command{
	Use:   "plan",
	Short: "run tofu plan",
	Run:   planHandler,
}

func init() {
	Plan.PersistentFlags().Bool(string(flag.Destroy), false, "Select the \"destroy\" planning mode, which "+
		"creates a plan\n  to destroy all objects currently managed by this\n  OpenTofu configuration instead "+
		"of the usual behavior.")
}

func planHandler(cmd *cobra.Command, args []string) {
	moduleDir, err := cmd.Flags().GetString(string(flag.ModuleDir))
	flag.HandleFlagErrAndValue(err, flag.ModuleDir, moduleDir)

	valueOverrides, err := cmd.Flags().GetStringToString(string(flag.Set))
	flag.HandleFlagErr(err, flag.Set)

	isDestroyPlan, err := cmd.Flags().GetBool(string(flag.Destroy))
	flag.HandleFlagErr(err, flag.Destroy)

	// Resolve manifest path with priority: --manifest > --input-dir > --kustomize-dir + --overlay
	targetManifestPath, isTemp, err := climanifest.ResolveManifestPath(cmd)
	if err != nil {
		log.Fatalf("failed to resolve manifest: %v", err)
	}
	if isTemp {
		defer os.Remove(targetManifestPath)
	}

	// Validate manifest before proceeding
	if err := manifest.Validate(targetManifestPath); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	credentialOptions, err := stackinputcredentials.BuildWithFlags(cmd.Flags())
	if err != nil {
		log.Fatalf("failed to build credentiaal options: %v", err)
	}

	err = tofumodule.RunCommand(
		moduleDir,
		targetManifestPath,
		terraform.TerraformOperationType_plan,
		valueOverrides,
		true,
		isDestroyPlan,
		credentialOptions...)
	if err != nil {
		log.Fatalf("failed to run tofu operation: %v", err)
	}
}
