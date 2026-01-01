package pulumi

import (
	"fmt"
	"os"

	"github.com/project-planton/project-planton/internal/cli/cliprint"
	"github.com/project-planton/project-planton/internal/cli/flag"
	climanifest "github.com/project-planton/project-planton/internal/cli/manifest"
	"github.com/project-planton/project-planton/internal/manifest"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumistack"
	"github.com/spf13/cobra"
)

var Delete = &cobra.Command{
	Use:     "delete",
	Short:   "delete/remove a pulumi stack",
	Aliases: []string{"rm"},
	Long: `Delete a Pulumi stack and all its configuration/state from the backend.

⚠️  WARNING: This is a destructive operation that permanently removes the stack metadata.

IMPORTANT: This command does NOT destroy cloud resources. If your stack still has 
resources deployed, you should run 'project-planton pulumi destroy' first to tear 
down the infrastructure, then use this command to remove the stack metadata.

Use --force to remove the stack even if it still contains resources (not recommended 
unless you're sure the resources have been cleaned up manually or through other means).`,
	Run: deleteHandler,
}

func deleteHandler(cmd *cobra.Command, args []string) {
	moduleDir, err := cmd.Flags().GetString(string(flag.ModuleDir))
	flag.HandleFlagErrAndValue(err, flag.ModuleDir, moduleDir)

	// Stack can be provided via flag or extracted from manifest
	stackFqdn, err := cmd.Flags().GetString(string(flag.Stack))
	flag.HandleFlagErr(err, flag.Stack)

	valueOverrides, err := cmd.Flags().GetStringToString(string(flag.Set))
	flag.HandleFlagErr(err, flag.Set)

	force, err := cmd.Flags().GetBool(string(flag.Force))
	flag.HandleFlagErr(err, flag.Force)

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
		cliprint.PrintError(fmt.Sprintf("Failed to resolve manifest: %v", err))
		os.Exit(1)
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

	cliprint.PrintStep("Deleting Pulumi stack...")

	noCleanup, _ := cmd.Flags().GetBool(string(flag.NoCleanup))
	moduleVersion, _ := cmd.Flags().GetString(string(flag.ModuleVersion))

	cliprint.PrintHandoff("Pulumi")

	err = pulumistack.Remove(moduleDir, stackFqdn, targetManifestPath, valueOverrides, force, moduleVersion, noCleanup)
	if err != nil {
		cliprint.PrintPulumiFailure()
		os.Exit(1)
	}
	cliprint.PrintPulumiSuccess()
}
