package tofu

import (
	"fmt"
	"os"

	"github.com/project-planton/project-planton/apis/project/planton/shared/iac/terraform"
	"github.com/project-planton/project-planton/internal/cli/cliprint"
	"github.com/project-planton/project-planton/internal/cli/flag"
	climanifest "github.com/project-planton/project-planton/internal/cli/manifest"
	"github.com/project-planton/project-planton/internal/manifest"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputproviderconfig"
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

	// Check which manifest source is being used for informative messages
	kustomizeDir, _ := cmd.Flags().GetString(string(flag.KustomizeDir))
	overlay, _ := cmd.Flags().GetString(string(flag.Overlay))

	if kustomizeDir != "" && overlay != "" {
		cliprint.PrintStep(fmt.Sprintf("Building manifest from kustomize overlay: %s", overlay))
	} else {
		cliprint.PrintStep("Loading manifest...")
	}

	// Resolve manifest path with priority: --manifest > --input-dir > --kustomize-dir + --overlay
	targetManifestPath, isTemp, err := climanifest.ResolveManifestPath(cmd)
	if err != nil {
		log.Fatalf("failed to resolve manifest: %v", err)
	}
	if isTemp {
		defer os.Remove(targetManifestPath)
	}

	cliprint.PrintSuccess("Manifest loaded")

	// Apply value overrides if any (creates new temp file if overrides exist)
	if len(valueOverrides) > 0 {
		cliprint.PrintStep(fmt.Sprintf("Applying %d field override(s)...", len(valueOverrides)))
	}

	finalManifestPath, isTempOverrides, err := manifest.ApplyOverridesToFile(targetManifestPath, valueOverrides)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if isTempOverrides {
		defer os.Remove(finalManifestPath)
		targetManifestPath = finalManifestPath
		cliprint.PrintSuccess("Overrides applied")
	}

	// Validate manifest before proceeding (after overrides are applied)
	cliprint.PrintStep("Validating manifest...")
	if err := manifest.Validate(targetManifestPath); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cliprint.PrintSuccess("Manifest validated")

	cliprint.PrintStep("Preparing OpenTofu execution...")
	providerConfigOptions, err := stackinputproviderconfig.BuildWithFlags(cmd.Flags())
	if err != nil {
		log.Fatalf("failed to build credentiaal options: %v", err)
	}
	cliprint.PrintSuccess("Execution prepared")

	cliprint.PrintHandoff("OpenTofu")

	err = tofumodule.RunCommand(
		moduleDir,
		targetManifestPath,
		terraform.TerraformOperationType_plan,
		valueOverrides,
		true,
		isDestroyPlan,
		providerConfigOptions...)
	if err != nil {
		log.Fatalf("failed to run tofu operation: %v", err)
	}
}
