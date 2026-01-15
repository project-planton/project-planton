package iacflags

import (
	"github.com/plantonhq/project-planton/internal/cli/flag"
	"github.com/spf13/cobra"
)

// AddManifestSourceFlags adds flags for specifying the manifest source.
// Priority order: --clipboard > --stack-input > --manifest > --input-dir > --kustomize-dir+--overlay
func AddManifestSourceFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolP(string(flag.Clipboard), "c", false,
		"read manifest content from system clipboard")

	cmd.PersistentFlags().StringP(string(flag.Manifest), "f", "",
		"path of the deployment-component manifest file")

	cmd.PersistentFlags().StringP(string(flag.StackInput), "i", "",
		"path to a YAML file containing the stack input (extracts manifest from target field)")

	cmd.PersistentFlags().String(string(flag.InputDir), "",
		"directory containing target.yaml and credential yaml files")

	cmd.PersistentFlags().String(string(flag.KustomizeDir), "",
		"directory containing kustomize configuration")

	cmd.PersistentFlags().String(string(flag.Overlay), "",
		"kustomize overlay to use (e.g., prod, dev, staging)")
}
