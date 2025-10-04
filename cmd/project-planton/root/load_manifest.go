package root

import (
	"os"

	"github.com/project-planton/project-planton/internal/cli/flag"
	"github.com/project-planton/project-planton/internal/manifest"
	"github.com/project-planton/project-planton/pkg/kustomize/builder"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var LoadManifest = &cobra.Command{
	Use:   "load-manifest [manifest-path]",
	Short: "load a project-planton manifest from provided path or kustomize",
	Example: `
	# Load from file
	project-planton load-manifest manifest.yaml
	
	# Load from kustomize
	project-planton load-manifest --kustomize-dir _kustomize --overlay prod
	
	# Load with overrides
	project-planton load-manifest --kustomize-dir _kustomize --overlay prod --set spec.version=v1.2.3
	`,
	Args: cobra.MaximumNArgs(1), // Optional manifest path
	Run:  loadManifestHandler,
}

func init() {
	LoadManifest.PersistentFlags().String(string(flag.KustomizeDir), "", "directory containing kustomize configuration")
	LoadManifest.PersistentFlags().String(string(flag.Overlay), "", "kustomize overlay to use (e.g., prod, dev, staging)")
	LoadManifest.PersistentFlags().StringToString(string(flag.Set), map[string]string{}, "override resource manifest values using key=value pairs")
}

func loadManifestHandler(cmd *cobra.Command, args []string) {
	valueOverrides, err := cmd.Flags().GetStringToString(string(flag.Set))
	flag.HandleFlagErr(err, flag.Set)

	var manifestPath string

	// If a positional arg is provided, use it as manifest path
	if len(args) > 0 {
		manifestPath = args[0]
	} else {
		// Otherwise, try to resolve from kustomize flags
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

	updatedManifest, err := manifest.LoadWithOverrides(manifestPath, valueOverrides)
	if err != nil {
		log.Fatal(err)
	}
	if err := manifest.Print(updatedManifest); err != nil {
		log.Fatal(err)
	}
}
