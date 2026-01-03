package tofu

import (
	"fmt"
	"os"

	"github.com/plantonhq/project-planton/apis/org/project_planton/shared/iac/terraform"
	"github.com/plantonhq/project-planton/internal/cli/cliprint"
	"github.com/plantonhq/project-planton/internal/cli/flag"
	climanifest "github.com/plantonhq/project-planton/internal/cli/manifest"
	"github.com/plantonhq/project-planton/internal/manifest"
	"github.com/plantonhq/project-planton/pkg/iac/stackinput/stackinputproviderconfig"
	"github.com/plantonhq/project-planton/pkg/iac/tofu/tofumodule"
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
	Destroy.PersistentFlags().String(string(flag.ModuleVersion), "",
		"Checkout a specific version (tag, branch, or commit SHA) of the IaC modules in the workspace copy.\n"+
			"This allows using a different module version than what's in the staging area without affecting it.")
	Destroy.PersistentFlags().Bool(string(flag.NoCleanup), false, "Do not cleanup the workspace copy after execution (keeps cloned modules)")
}

func destroyHandler(cmd *cobra.Command, args []string) {
	isAutoApprove, err := cmd.Flags().GetBool(string(flag.AutoApprove))
	flag.HandleFlagErr(err, flag.AutoApprove)

	moduleDir, err := cmd.Flags().GetString(string(flag.ModuleDir))
	flag.HandleFlagErrAndValue(err, flag.ModuleDir, moduleDir)

	valueOverrides, err := cmd.Flags().GetStringToString(string(flag.Set))
	flag.HandleFlagErr(err, flag.Set)

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

	moduleVersion, _ := cmd.Flags().GetString(string(flag.ModuleVersion))
	noCleanup, _ := cmd.Flags().GetBool(string(flag.NoCleanup))

	err = tofumodule.RunCommand(moduleDir, targetManifestPath, terraform.TerraformOperationType_destroy, valueOverrides,
		isAutoApprove,
		false,
		moduleVersion, noCleanup,
		providerConfigOptions...)
	if err != nil {
		log.Fatalf("failed to run tofu operation: %v", err)
	}
}
