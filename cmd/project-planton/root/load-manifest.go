package root

import (
	"github.com/plantoncloud/project-planton/internal/cli/flag"
	"github.com/plantoncloud/project-planton/internal/manifest"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var LoadManifest = &cobra.Command{
	Use:   "load-manifest",
	Short: "load a project-planton manifest from provided path",
	Example: `
	project-planton load manifest.yaml
	`,
	Args: cobra.ExactArgs(1), //path of the manifest to load
	Run:  loadManifestHandler,
}

func init() {
	LoadManifest.PersistentFlags().StringToString(string(flag.Set), map[string]string{}, "override resource manifest values using key=value pairs")
}

func loadManifestHandler(cmd *cobra.Command, args []string) {
	valueOverrides, err := cmd.Flags().GetStringToString(string(flag.Set))
	flag.HandleFlagErr(err, flag.Set)

	manifestPath := args[0]
	updatedManifest, err := manifest.LoadWithOverrides(manifestPath, valueOverrides)
	if err != nil {
		log.Fatal(err)
	}
	if err := manifest.Print(updatedManifest); err != nil {
		log.Fatal(err)
	}
}
