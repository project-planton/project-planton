package root

import (
	"fmt"
	"os"

	"github.com/plantonhq/project-planton/internal/cli/cliprint"
	"github.com/plantonhq/project-planton/internal/cli/flag"
	"github.com/plantonhq/project-planton/internal/manifest"
	"github.com/plantonhq/project-planton/pkg/kustomize/builder"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ValidateManifest = &cobra.Command{
	Use:   "validate-manifest [manifest-path]",
	Short: "validate a project-planton manifest",
	Aliases: []string{
		"validate",
	},
	Example: `
	# Validate from file
	project-planton validate manifest.yaml
	
	# Validate from kustomize
	project-planton validate --kustomize-dir _kustomize --overlay prod
	`,
	Args: cobra.MaximumNArgs(1), // Optional manifest path
	Run:  validateHandler,
}

func init() {
	ValidateManifest.PersistentFlags().String(string(flag.KustomizeDir), "", "directory containing kustomize configuration")
	ValidateManifest.PersistentFlags().String(string(flag.Overlay), "", "kustomize overlay to use (e.g., prod, dev, staging)")
}

func validateHandler(cmd *cobra.Command, args []string) {
	var manifestPath string
	var err error

	// If a positional arg is provided, use it as manifest path
	if len(args) > 0 {
		manifestPath = args[0]
	} else {
		// Otherwise, try to build from kustomize flags
		kustomizeDir, _ := cmd.Flags().GetString(string(flag.KustomizeDir))
		overlay, _ := cmd.Flags().GetString(string(flag.Overlay))

		if kustomizeDir != "" && overlay != "" {
			// Build manifest from kustomize
			manifestPath, err = builder.BuildManifest(kustomizeDir, overlay)
			if err != nil {
				log.Fatalf("failed to build kustomize manifest: %v", err)
			}
			defer os.Remove(manifestPath)
		} else if kustomizeDir != "" || overlay != "" {
			log.Fatal("both --kustomize-dir and --overlay flags must be provided together")
			return
		} else {
			log.Fatal("must provide either a manifest path or (--kustomize-dir + --overlay)")
			return
		}
	}

	err = manifest.Validate(manifestPath)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	cliprint.PrintSuccessMessage("manifest is valid")
}
